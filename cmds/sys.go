package cmds

import (
	"sys_tools/utils"
	"time"
)

//-memo 命令处理海曙
func memoCmd() {
	utils.Clear()
	for {
		utils.MoveCursor(1, 1)
		utils.Memo()
		utils.Println()
		utils.Swap()
		time.Sleep(time.Duration(3 * time.Second))
	}
}

//-cpu 命令处理函数
func cpuCmd() {
	utils.Clear()
	for {
		utils.MoveCursor(1, 1)
		utils.CPUInfo()
		utils.CPULoad()
		utils.Println()
		utils.CPUTimes(false)
		time.Sleep(time.Duration(3 * time.Second))
	}
}

//-net 命令处理函数
func netCmd() {
	cmd2 := utils.GetSecondCmdLineArgs()
	switch cmd2 {
	case "-card":
		utils.NetInterfaces()
		break
	case "-conn":
		utils.NetConnections()
		break
	default:
		helpCmd()
		break
	}

}

//-sys 命令处理函数
func sysCmd() {
	utils.SystemInfo()
}

//-proc 命令处理函数
func processInfoCmd() {
	params := utils.GetSecondCmdLineArgs()
	if utils.IsBlankStr(params) {
		helpCmd()
		return
	}

	switch params {
	case "-pid":
		param3 := utils.GetCmdLineArgs(3)
		if utils.IsBlankStr(params) {
			helpCmd()
			return
		}

		pid, err1 := utils.Str2Int(param3)
		if err1 != nil {
			utils.Println("请使用正确的pid")
			return
		}
		utils.ProcessInfo(int32(pid))
		break
	case "-list":
		utils.ProcessListByKeyword("")
		break
	default:
		helpCmd()
		break
	}

}

//-disk 命令处理函数
func diskCmd() {
	utils.Disk()
}
