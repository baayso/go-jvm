package stack

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// 栈指令直接对操作数栈进行操作，共9条：
// pop和pop2指令将栈顶变量弹出，
// dup系列指令复制栈顶变量，
// swap指令交换栈顶的两个变量。
// 和其他类型的指令不同，栈指令并不关心变量类型。

// Duplicate the top operand stack value
type DUP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
             \_
               |
               V
[...][c][b][a][a]
*/
func (d *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	slot := stack.PopSlot()

	stack.PushSlot(slot)

	stack.PushSlot(slot)
}

// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
func (d *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot1)

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
func (d *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot1)

	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values
type DUP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
func (d *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values and insert two or three values down
type DUP2_X1 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
func (d *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type DUP2_X2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/
func (d *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
