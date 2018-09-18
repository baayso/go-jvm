/*
	栈指令直接对操作数栈进行操作，共9条：
	pop和pop2指令将栈顶变量弹出，
	dup系列指令复制栈顶变量，
	swap指令交换栈顶的两个变量。
	和其他类型的指令不同，栈指令并不关心变量类型。
*/
package stack

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/
func (s *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot1)

	stack.PushSlot(slot2)
}
