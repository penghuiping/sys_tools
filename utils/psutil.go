package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"strconv"
	"strings"
	"syscall"
)

//CPUInfo 获取cpu信息
func CPUInfo() {
	stats, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("cpu信息:")
	for i := 0; i < len(stats); i++ {
		fmt.Println("\tcpu名称:", stats[i].ModelName)
		fmt.Println("\tcpu主频:", stats[i].Mhz, "Mhz")
		fmt.Println("\tcpu核数:", stats[i].Cores)
		fmt.Println()
	}
}

//CPUTimes 获取cpu指标
func CPUTimes(percpu bool) {
	stats1, err1 := cpu.Times(percpu)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Printf("%-10s\t%s\t%s\t%s\t%s\t%s\n", "name", "user", "sys", "idle", "iowait", "nice")
	for i := range stats1 {
		stat1 := stats1[i]
		fmt.Printf("%-10s\t%.2f\t%.2f\t%.2f\t%.2f\t%.2f\n",
			stat1.CPU, stat1.User/stat1.Total(),
			stat1.System/stat1.Total(), stat1.Idle/stat1.Total(),
			stat1.Iowait/stat1.Total(), stat1.Nice)
	}
}

//CPULoad 获取cpu负载
func CPULoad() {
	avg, err2 := load.Avg()
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println("1分钟load:\t", avg.Load1)
	fmt.Println("5分钟load:\t", avg.Load5)
	fmt.Println("15分钟load:\t", avg.Load15)
}

//NetInterfaces 网卡信息
func NetInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-10s\t%-20s\t%s\t%-35s\t%s\n", "名称", "mac", "mtu", "状态", "ip地址")
	for i := range interfaces {
		inter := interfaces[i]
		status := ""
		for index, flag := range inter.Flags {
			status = status + flag
			if index != len(inter.Flags)-1 {
				status = status + ","
			}
		}
		fmt.Printf("%-10s\t%-20s\t%d\t%-35s\t%s\n", inter.Name, inter.HardwareAddr, inter.MTU, status, inter.Addrs)
	}
}

//NetConnections 网络连接信息
func NetConnections() {
	conns, err := net.Connections("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\t%-40s\t%-40s\t%-10s\t%s\n", "PID", "SRC", "DIST", "STATUS", "TYPE")
	for _, conn := range conns {
		laddr := conn.Laddr.IP + ":" + strconv.Itoa(int(conn.Laddr.Port))
		raddr := conn.Raddr.IP + ":" + strconv.Itoa(int(conn.Raddr.Port))
		fmt.Printf("%d\t%-40s\t%-40s\t%-10s\t%s\n", conn.Pid, laddr, raddr, conn.Status, UDPAndTCPMap(conn.Type))
	}
}

/***
"unix": syscall.AF_UNIX,
"TCP":  syscall.SOCK_STREAM,
"UDP":  syscall.SOCK_DGRAM,
"IPv4": syscall.AF_INET,
"IPv6": syscall.AF_INET6,
**/
//UDPAndTCPMap ...
func UDPAndTCPMap(value uint32) string {
	var name string = ""
	switch value {
	case syscall.SOCK_STREAM:
		name = "TCP"
		break
	case syscall.SOCK_DGRAM:
		name = "UDP"
		break
	}
	return name
}

//Memo 内存信息
func Memo() {
	stats, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("内存信息:")
	fmt.Println("\t总共:", stats.Total/(1024*1024), "MB")
	fmt.Println("\t空闲:", stats.Free/(1024*1024), "MB")
	fmt.Println("\t已用:", stats.Used/(1024*1024), "MB")
	fmt.Printf("\t使用率:%.2f%s\n", stats.UsedPercent, "%")
}

//Swap 交换区
func Swap() {
	stats1, err1 := mem.SwapMemory()
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println("交换区信息:")
	fmt.Println("\t总共:", stats1.Total/(1024*1024), "MB")
	fmt.Println("\t空闲:", stats1.Free/(1024*1024), "MB")
	fmt.Println("\t已用:", stats1.Used/(1024*1024), "MB")
	fmt.Printf("\t使用率:%.2f%s\n", stats1.UsedPercent, "%")
}

//Disk 磁盘
func Disk() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-20s \t %-30s \t %-10s \t %-30s \t %-10s \t %-10s \t %-10s \t %-10s\n", "Device", "Mountpoint", "Fstype", "Opts", "Total(MB)", "Free(MB)", "Used(MB)", "Percent")
	for _, p := range partitions {
		usage, _ := disk.Usage(p.Mountpoint)
		fmt.Printf("%-20s \t %-30s \t %-10s \t %-30s \t %-10d \t %-10d \t %-10d \t %-10.2f\n", p.Device, p.Mountpoint, p.Fstype, p.Opts, usage.Total/1000000, usage.Free/1000000, usage.Used/1000000, usage.UsedPercent)
	}
}

//ProcessListByKeyword 根据关键字获取进程列表
func ProcessListByKeyword(keyword string) {
	processes, err := process.Processes()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s \t %-40s \t %s\n", "pid", "name", "cmdline")
	for _, p := range processes {
		name, _ := p.Name()
		cmdline, _ := p.Cmdline()

		if !IsBlankStr(keyword) {
			if strings.Contains(cmdline, keyword) {
				fmt.Printf("%d \t %-40s \t %s\n", p.Pid, name, cmdline)
			}
		} else {
			fmt.Printf("%d \t %-40s \t %s\n", p.Pid, name, cmdline)
		}
	}
}

//ProcessInfo 根据pid获取进程详细信息
func ProcessInfo(pid int32) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(proc.MemoryInfo())
	fmt.Println(proc.MemoryPercent())
	fmt.Println(proc.NumThreads())
}

//SystemInfo 获取操作系统基本信息
func SystemInfo() {
	info, err := host.Info()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("操作系统信息:")
	fmt.Println("\thostname:", info.Hostname)
	fmt.Println("\thostId:", info.HostID)
	fmt.Println("\tkernelVersion:", info.KernelVersion)
	fmt.Println("\tkernelArch:", info.KernelArch)
	fmt.Println("\tos:", info.OS)
	fmt.Println("\tplatformVersion:", info.PlatformVersion)
	fmt.Println("\tprocess number:", info.Procs)
	fmt.Println("\tuptime:", info.Uptime, "秒")
}
