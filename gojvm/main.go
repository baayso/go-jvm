package main

import (
	"fmt"
	"strings"

	"github.com/baayso/jvm/gojvm/classfile"
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

	className := strings.Replace(cmd.Class(), ".", "/", -1)

	cf := loadClass(className, cp)

	println(cmd.Class())

	printClassInfo(cf)
}

// 读取并解析class文件
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)

	if err != nil {
		panic(err)
	}

	return cf
}

// 打印class文件的一些重要信息
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())

	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}

	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}
