package cmds

import (
	"fmt"
	"strings"
	"sys_tools/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

//获取内存信息
func memoCmd() {
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

	stats1, err1 := mem.SwapMemory()
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println("\n交换区信息:")
	fmt.Println("\t总共:", stats1.Total/(1024*1024), "MB")
	fmt.Println("\t空闲:", stats1.Free/(1024*1024), "MB")
	fmt.Println("\t已用:", stats1.Used/(1024*1024), "MB")
	fmt.Printf("\t使用率:%.2f%s\n", stats1.UsedPercent, "%")
}

//获取cpu信息
func cpuCmd() {
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

	fmt.Println()
	stats1, err1 := cpu.Times(true)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println("名称    \t用户态    \t系统态    \t空闲    \tio等待    \tnice")
	for i := range stats1 {
		stat1 := stats1[i]
		fmt.Printf("%s    \t%.2f    \t%.2f    \t%.2f    \t%.2f    \t%.2f\n",
			stat1.CPU, stat1.User/stat1.Total(),
			stat1.System/stat1.Total(), stat1.Idle/stat1.Total(),
			stat1.Iowait/stat1.Total(), stat1.Nice)
	}

}

// 网络信息
func netCmd() {
	cmd2 := utils.GetSecondCmdLineArgs()
	switch cmd2 {
	case "-card":
		netInterfaces()
		break
	case "-conn":
		netConnections()
		break
	default:
		helpCmd()
		break
	}

}

//网卡信息
func netInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-10s\t%-20s\t%s\t%-35s\t%s\n", "名称", "mac", "mtu", "状态", "ip地址")
	for i := range interfaces {
		// fmt.Println(interfaces[i])
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

//网络连接信息
func netConnections() {
	conns, err := net.Connections("")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, conn := range conns {
		fmt.Println(conn.Fd, conn.Family)
	}
}

//获取操作系统信息
func sysCmd() {
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
	fmt.Println("\t当前进程数:", info.Procs)
	fmt.Println("\t系统运行时长:", info.Uptime, "秒")
}

//获取进程列表
func processCmd(keywords string) {
	processes, err := process.Processes()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s \t %-40s \t %s\n", "pid", "name", "cmdline")
	for _, p := range processes {

		name, _ := p.Name()
		cmdline, _ := p.Cmdline()

		if !utils.IsBlankStr(keywords) {
			if strings.Contains(cmdline, keywords) {
				fmt.Printf("%d \t %-40s \t %s\n", p.Pid, name, cmdline)
			}
		} else {
			fmt.Printf("%d \t %-40s \t %s\n", p.Pid, name, cmdline)
		}
	}
}

//获取指定进程的详细信息
func processInfoCmd() {
	params := utils.GetSecondCmdLineArgs()
	if utils.IsBlankStr(params) {
		helpCmd()
		return
	}

	pid, err1 := utils.Str2Int(params)
	if err1 != nil {
		helpCmd()
		return
	}

	processes, err := process.Processes()
	if err != nil {
		fmt.Println(err)
		return
	}

	var proc *process.Process = nil

	for _, e := range processes {
		if int(e.Pid) == pid {
			proc = e
			break
		}
	}
	fmt.Println(proc.MemoryInfo())
	fmt.Println(proc.MemoryPercent())
	fmt.Println(proc.NumThreads())
	fmt.Println()
}

func diskCmd() {
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
