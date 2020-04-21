package cmds

import (
	"fmt"
	"sys_tools/utils"
)

//-memo 命令处理海曙
func memoCmd() {
	utils.Memo()
	fmt.Println()
	utils.Swap()
}

//-cpu 命令处理函数
func cpuCmd() {
	utils.CPUInfo()
	fmt.Println()
	utils.CPUTimes(false)
	fmt.Println()
	utils.CPUTimes(true)
	fmt.Println()
	utils.CPULoad()
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

	pid, err1 := utils.Str2Int(params)
	if err1 != nil {
		helpCmd()
		return
	}
	utils.ProcessInfo(int32(pid))
}

//-disk 命令处理函数
func diskCmd() {
	utils.Disk()
}
