package cmds

import (
	"fmt"
	"reflect"
	"sort"
	"sys_tools/utils"
)

func getCmd() map[string]string {

	return map[string]string{
		"-V":    "打印版本信息",
		"-help": "打印帮助信息",
		"-java": `java相关命令:

		  -version 获取java当前版本
		  -heap 获取java堆内存信息 -heap [pid]
		  -dump 输出java堆内存信息到文件 -dump [pid] [path]
		  -pid 获取java进程pid
		  `,
		"-memo": "打印系统内存信息",
		"-cpu":  "打印系统cpu信息",
		"-sys":  "打印操作系统信息",
		"-proc": "获取某个进程信息，使用方法 -proc [pid]",
		"-disk": "获取磁盘信息",
		"-net": `打印系统网络信息
		 
		  -card 获取网卡信息
		  -conn 获取socket连接信息
		  `,
		"-http": `http请求,使用方法:

		  -http [url] -X [get/post] -H [header1] -H [header2] -D [body]
		  -X 指定http方法
		  -H 指定http头信息
		  -D 指定http体
		  `,
	}
}

func cmdMap() map[string]interface{} {
	return map[string]interface{}{
		"-help": helpCmd,
		"-V":    versionCmd,
		"-java": javaCmd,
		"-memo": memoCmd,
		"-cpu":  cpuCmd,
		"-net":  netCmd,
		"-sys":  sysCmd,
		"-http": httpCmd,
		"-proc": processInfoCmd,
		"-disk": diskCmd,
	}
}

//CmdExec ...
func CmdExec() {
	cmd := utils.GetFirstCmdLineArgs()
	method := cmdMap()[cmd]
	if method != nil {
		f := reflect.ValueOf(method)
		f.Call([]reflect.Value{})
	} else {
		helpCmd()
	}
}

//打印帮助
func helpCmd() {
	fmt.Println("帮助信息:")
	cmd := getCmd()
	list := make([]string, 0)
	for k := range cmd {
		list = append(list, k)
	}
	sort.Strings(list)
	for _, v := range list {
		fmt.Println("\t"+v+":", "\t"+cmd[v])
	}
}

//打印版本信息
func versionCmd() {
	fmt.Println("当前版本为:0.0.1-SNAPSHOT")
}
