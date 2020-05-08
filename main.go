package main

import (
	"github.com/penghuiping/sys_tools/cmds"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	controlCExit()
	cmds.CmdExec()
}

//用于处理ctrl+c程序退出
func controlCExit() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)

	go func() {
		<-channel
		os.Exit(0)
	}()
}
