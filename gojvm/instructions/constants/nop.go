/*
	常量指令把常量推入操作数栈顶。
	常量可以来自三个地方：隐含在操作码里、操作数和运行时常量池。
	常量指令共有21条。
*/
package constants

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// 什么都不做的指令
type NOP struct {
	base.NoOperandsInstruction
}

func (n *NOP) Execute(frame *rtda.Frame) {
	//  什么也不用做
}
