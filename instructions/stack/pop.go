/*
	栈指令直接对操作数栈进行操作，共9条：
	pop和pop2指令将栈顶变量弹出，
	dup系列指令复制栈顶变量，
	swap指令交换栈顶的两个变量。
	和其他类型的指令不同，栈指令并不关心变量类型。
*/
package stack

import (
	"github.com/baayso/go-jvm/instructions/base"
	rtda "github.com/baayso/go-jvm/runtimedataarea"
)

// pop指令把栈顶变量弹出
// pop指令用于弹出int、float等占用一个操作数栈位置的变量
type POP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (p *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// pop2指令用于弹出double、long在操作数栈中占据两个位置的变量
type POP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (p *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
