package command

import (
	"flag"
	"fmt"
	"os"
)

// 命令行选项和参数
type Command struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

// getters
func (c *Command) HelpFlag() bool {
	return c.helpFlag
}

func (c *Command) VersionFlag() bool {
	return c.versionFlag
}

func (c *Command) CpOption() string {
	return c.cpOption
}

func (c *Command) Class() string {
	return c.class
}

func (c *Command) Args() []string {
	return c.args
}

// 解析命令
func ParseCommand() *Command {
	cmd := &Command{}

	flag.Usage = PrintUsage

	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd

}

// 如果 ParseCmd() 函数解析命令失败，它就调用 printUsage() 函数把命令的用法打印到控制台
func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
