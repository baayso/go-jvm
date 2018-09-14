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
