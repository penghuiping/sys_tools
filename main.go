package main

import (
	"os"
	"os/signal"
	"sys_tools/cmds"
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
