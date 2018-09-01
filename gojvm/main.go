package main

import (
	"fmt"

	cmd "github.com/baayso/jvm/gojvm/command"
)

func main() {

	// 先调用 parseCommand()函数解析命令行参数，如果一切正常，
	// 则调 startJVM()函数启动Java虚拟机。如果解析出现错误，
	// 或者用户输入了-help选项，则调用 printUsage()函数打印出
	// 帮助信息。如果用户输入了-version选项，则输出版本信息。

	command := cmd.ParseCommand()

	if command.VersionFlag() {
		fmt.Println("version 0.0.1")
	} else if command.HelpFlag() || command.Class() == "" {
		cmd.PrintUsage()
	} else {
		startJVM(command)
	}
}

func startJVM(command *cmd.Command) {
	fmt.Printf("classpath:%s class:%s args:%v\n", command.CpOption(), command.Class(), command.Args())
}
