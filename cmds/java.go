package cmds

import (
	"fmt"
	"sys_tools/utils"
)

//java程序相关命令
func javaCmd() {
	cmd2 := utils.GetSecondCmdLineArgs()

	if utils.IsBlankStr(cmd2) {
		helpCmd()
		return
	}

	switch cmd2 {
	case "-version":
		utils.ExecShellCmd("java -version")
		break
	case "-pid":
		utils.ProcessListByKeyword("java")
		break
	case "-heap":
		pid := utils.GetCmdLineArgs(3)
		if utils.IsBlankStr(pid) {
			helpCmd()
			return
		}
		utils.ExecShellCmd("jmap -histo " + pid)
		break
	case "-dump":
		pid := utils.GetCmdLineArgs(3)
		path := utils.GetCmdLineArgs(4)
		if utils.IsBlankStr(pid) || utils.IsBlankStr(path) {
			helpCmd()
			return
		}
		res := fmt.Sprintf("jmap -dump:format=b,file=%s %s", path, pid)
		utils.ExecShellCmd(res)
	default:
		helpCmd()
		break
	}
}
