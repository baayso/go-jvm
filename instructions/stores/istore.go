/*
	存储指令与加载指令刚好相反，存储指令把变量从操作数栈顶弹出，然后存入局部变量表。
	和加载指令一样，存储指令也可以分为6类。
*/
package stores

import (
	"github.com/baayso/go-jvm/instructions/base"
	rtda "github.com/baayso/go-jvm/runtimedataarea"
)

type ISTORE struct{ base.Index8Instruction }
type ISTORE_0 struct{ base.NoOperandsInstruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type ISTORE_3 struct{ base.NoOperandsInstruction }

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()

	frame.LocalVars().SetInt(index, val)
}

func (i *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, i.Index)
}

func (i *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (i *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (i *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (i *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
