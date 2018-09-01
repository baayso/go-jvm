package main

import (
	"fmt"
	"strings"

	"github.com/baayso/jvm/gojvm/classpath"
	"github.com/baayso/jvm/gojvm/command"
)

func main() {

	// 先调用 command.ParseCommand()函数解析命令行参数，如果一切正常，
	// 则调 startJVM()函数启动Java虚拟机。如果解析出现错误，
	// 或者用户输入了-help选项，则调用 command.PrintUsage()函数打印出
	// 帮助信息。如果用户输入了 -version选项，则输出版本信息。

	cmd := command.ParseCommand()

	if cmd.VersionFlag() {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag() || cmd.Class() == "" {
		command.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *command.Command) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption())

	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.Class(), cmd.Args())

	className := strings.Replace(cmd.Class(), ".", "/", -1)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.Class())
		return
	}

	fmt.Printf("class data: %v\n", classData)
}
