package stores

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// 存储指令把变量从操作数栈顶弹出，然后存入局部变量表。

type FSTORE struct{ base.Index8Instruction }
type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()

	frame.LocalVars().SetFloat(index, val)
}

func (f *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, f.Index)
}

func (f *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

func (f *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

func (f *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

func (f *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
