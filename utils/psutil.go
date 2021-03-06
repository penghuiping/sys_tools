package utils

import (
	"sort"
	"strings"
	"syscall"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

//CPUInfo 获取cpu信息
func CPUInfo() {
	stats, err := cpu.Info()
	if err != nil {
		Println(err)
		return
	}
	Println("cpu信息:")
	for i := 0; i < len(stats); i++ {
		Println("\tcpu名称:", stats[i].ModelName)
		Println("\tcpu主频:", stats[i].Mhz, "Mhz")
		Println("\tcpu核数:", stats[i].Cores)
		Println()
	}
}

//用于 非totalcpu
var percpuStats []cpu.TimesStat = nil

//用于 totalcpu
var totalcpuStats []cpu.TimesStat = nil

//CPUTimes 获取cpu指标
func CPUTimes(percpu bool) {

	stats1, err1 := cpu.Times(percpu)
	if err1 != nil {
		Println(err1)
		return
	}

	var lastTimeStats []cpu.TimesStat = nil

	Printf("%-10s|%-15s|%-15s|%-15s|%-15s|%-15s|%-15s|%-15s\n",
		"Name", "User", "System", "Idle", "Iowait", "Nice", "Irq", "Softirq")
	for i := range stats1 {
		stat1 := stats1[i]

		if percpu {
			lastTimeStats = percpuStats
		} else {
			lastTimeStats = totalcpuStats
		}

		if lastTimeStats != nil && len(lastTimeStats) > 0 {
			//可以计算percent
			user := (stat1.User - lastTimeStats[i].User) / (stat1.Total() - lastTimeStats[i].Total()) * 100
			system := (stat1.System - lastTimeStats[i].System) / (stat1.Total() - lastTimeStats[i].Total()) * 100
			idle := (stat1.Idle - lastTimeStats[i].Idle) / (stat1.Total() - lastTimeStats[i].Total()) * 100
			iowait := (stat1.Iowait - lastTimeStats[i].Iowait) / (stat1.Total() - lastTimeStats[i].Total()) * 100
			nice := (stat1.Nice - lastTimeStats[i].Nice) / (stat1.Total() - lastTimeStats[i].Total()) * 100
			irq := (stat1.Irq - lastTimeStats[i].Irq) / (stat1.Total() - lastTimeStats[i].Total()) * 100
			softIrq := (stat1.Softirq - lastTimeStats[i].Softirq) / (stat1.Total() - lastTimeStats[i].Total()) * 100

			Printf("%-10s|%-15.2f|%-15.2f|%-15.2f|%-15.2f|%-15.2f|%-15.2f|%-15.2f\n",
				stat1.CPU, user, system, idle, iowait, nice, irq, softIrq)
		} else {
			Printf("%-10s|%-15.2f|%-15.2f|%-15.2f|%-15.2f|%-15.2f|%-15.2f|%-15.2f\n",
				stat1.CPU, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
		}

	}

	if percpu {
		percpuStats = stats1
	} else {
		totalcpuStats = stats1
	}
}

//CPULoad 获取cpu负载
func CPULoad() {
	avg, err2 := load.Avg()
	if err2 != nil {
		Println(err2)
		return
	}

	Println("cpu负载:")
	Printf("\t1分钟load: %.2f\n", avg.Load1)
	Printf("\t5分钟load: %.2f\n", avg.Load5)
	Printf("\t15分钟load: %.2f\n", avg.Load15)
}

//NetInterfaces 网卡信息
func NetInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		Println(err)
		return
	}

	Printf("%-15s|%-20s|%-10s|%-35s|%s\n", "Name", "Mac", "MTU", "Status", "Ip")
	for i := range interfaces {
		inter := interfaces[i]
		status := ""
		for index, flag := range inter.Flags {
			status = status + flag
			if index != len(inter.Flags)-1 {
				status = status + ","
			}
		}
		Printf("%-15s|%-20s|%-10d|%-35s|%s\n", inter.Name, inter.HardwareAddr, inter.MTU, status, inter.Addrs)
	}
}

//NetConnections 网络连接信息
func NetConnections() {
	conns, err := net.Connections("inet")
	if err != nil {
		Println(err)
		return
	}

	Printf("%-8s|%-40s|%-40s|%-20s|%s\n", "PID", "SRC_IP", "DIST_IP", "STATUS", "TYPE")
	for _, conn := range conns {
		laddr := conn.Laddr.IP + ":" + Int2Str(int(conn.Laddr.Port))
		raddr := conn.Raddr.IP + ":" + Int2Str(int(conn.Raddr.Port))
		Printf("%-8d|%-40s|%-40s|%-20s|%s\n", conn.Pid, laddr, raddr, conn.Status, UDPAndTCPMap(conn.Type))
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
		Println(err)
	}

	Printf(`内存信息: 
	总共:%dMB
	空闲:%dMB
	已用:%dMB
	使用率:%.2f%s`,
		stats.Total/(1024*1024), stats.Free/(1024*1024), stats.Used/(1024*1024), stats.UsedPercent, "%")

}

//Swap 交换区
func Swap() {
	stats1, err1 := mem.SwapMemory()
	if err1 != nil {
		Println(err1)
	}
	Printf(`交换区信息: 
	总共:%dMB
	空闲:%dMB
	已用:%dMB
	使用率:%.2f%s`,
		stats1.Total/(1024*1024), stats1.Free/(1024*1024), stats1.Used/(1024*1024), stats1.UsedPercent, "%")

}

//Disk 磁盘
func Disk() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		Println(err)
		return
	}

	Printf("%-20s|%-10s|%-40s|%-10s|%-10s|%-10s|%-10s|%s\n",
		"Device", "Fstype", "Opts", "Total(MB)", "Free(MB)", "Used(MB)", "Percent(%)", "Mountpoint")
	for _, p := range partitions {
		usage, _ := disk.Usage(p.Mountpoint)
		Printf("%-20s|%-10s|%-40s|%-10d|%-10d|%-10d|%-10.2f|%s\n",
			p.Device, p.Fstype, p.Opts, usage.Total/1000000, usage.Free/1000000, usage.Used/1000000, usage.UsedPercent, p.Mountpoint)
	}
}

type processInfo struct {
	pid         int32
	name        string
	cmdline     string
	CPUPercent  float64
	MemoPercent float64
}

//ProcessListByKeyword 根据关键字获取进程列表
func ProcessListByKeyword(keyword string) {
	processes, err := process.Processes()

	if err != nil {
		Println(err)
		return
	}

	var processeInfos []processInfo = make([]processInfo, 0)
	for _, p := range processes {
		tmp := processInfo{}
		name, _ := p.Name()
		cmdline, _ := p.Cmdline()
		cpu, _ := p.CPUPercent()
		memo, _ := p.MemoryPercent()
		p.Threads()
		tmp.pid = p.Pid
		tmp.name = name
		tmp.cmdline = cmdline
		tmp.CPUPercent = cpu
		tmp.MemoPercent = float64(memo)
		processeInfos = append(processeInfos, tmp)
	}

	less := func(i, j int) bool {
		return processeInfos[i].CPUPercent > processeInfos[j].CPUPercent
	}
	sort.SliceStable(processeInfos, less)

	Clear()
	MoveCursor(1, 1)
	Printf("%-5s|%-10s|%-10s|%s\n", "Pid", "Cpu(%)", "Memo(%)", "Name")
	for i, p := range processeInfos {
		if i > 20 {
			break
		}

		if !IsBlankStr(keyword) {
			if strings.Contains(p.cmdline, keyword) {
				Printf("%-5d|%-10.2f|%-10.2f|%s\n", p.pid, p.CPUPercent, p.MemoPercent, p.name)
			}
		} else {
			Printf("%-5d|%-10.2f|%-10.2f|%s\n", p.pid, p.CPUPercent, p.MemoPercent, p.name)
		}
	}
}

//ProcessInfo 根据pid获取进程详细信息
func ProcessInfo(pid int32) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		Println(err)
		return
	}

	cpuPercent, _ := proc.CPUPercent()
	memo, _ := proc.MemoryInfo()
	percent, _ := proc.MemoryPercent()
	threads, _ := proc.NumThreads()
	cmdline, _ := proc.Cmdline()
	name, _ := proc.Name()
	ctx, _ := proc.NumCtxSwitches()
	conns, _ := proc.Connections()

	Println("进程名:", name)
	Printf("cpu使用率:%.2f%s\n", cpuPercent, "%")
	Printf("内存使用率:%.2f%s\n", percent, "%")
	Println("内存信息:", memo)
	Println("线程数量:", threads)
	Println("网络连接:", conns)
	Println("上下文切换数量:", ctx)
	Println("启动命令:", cmdline)
}

//SystemInfo 获取操作系统基本信息
func SystemInfo() {
	info, err := host.Info()

	if err != nil {
		Println(err)
		return
	}
	Println("操作系统信息:")
	Println("\thostname:", info.Hostname)
	Println("\thostId:", info.HostID)
	Println("\tkernelVersion:", info.KernelVersion)
	Println("\tkernelArch:", info.KernelArch)
	Println("\tos:", info.OS)
	Println("\tplatformVersion:", info.PlatformVersion)
	Println("\tprocess number:", info.Procs)
	Println("\tuptime:", info.Uptime, "秒")
}

//NumberOfProccess ...
// func NumberOfProccess(status ProcessStatus) {
// 	// processes, err := process.Processes()

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// for _, proc := range processes {
// 	// 	proc.Status
// 	// }

// }

//ProcessStatus 进程状态
type ProcessStatus string

//ProcessRunning 进程运行
const ProcessRunning ProcessStatus = "R"

//ProcessSleep 睡眠状态
const ProcessSleep ProcessStatus = "S"

//ProcessStop 运行状态
const ProcessStop ProcessStatus = "T"

//ProcessIdle 空闲状态
const ProcessIdle ProcessStatus = "I"

//ProcessZombie 僵尸进程
const ProcessZombie ProcessStatus = "Z"

//ProcessWait 等待状态
const ProcessWait ProcessStatus = "W"

//ProcessLock  锁状态
const ProcessLock ProcessStatus = "L"
