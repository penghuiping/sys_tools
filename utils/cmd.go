package utils

import (
	"os"
	"os/exec"
)

//GetFirstCmdLineArgs 获取命令行第一个参数
func GetFirstCmdLineArgs() string {
	return GetCmdLineArgs(1)
}

//GetSecondCmdLineArgs 获取命令行第二个参数
func GetSecondCmdLineArgs() string {
	return GetCmdLineArgs(2)
}

//GetCmdLineArgs 获取命令行指定的第几个参数
func GetCmdLineArgs(index int) string {
	if len(os.Args) > index {
		return os.Args[index]
	}
	return ""
}

//ExecShellCmd 执行shell命令
func ExecShellCmd(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	result, err := cmd.CombinedOutput()
	if err != nil {
		Println(err)
	} else {
		Println(string(result))
	}
}
