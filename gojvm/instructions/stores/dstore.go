package stores

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// 存储指令把变量从操作数栈顶弹出，然后存入局部变量表。

type DSTORE struct{ base.Index8Instruction }
type DSTORE_0 struct{ base.NoOperandsInstruction }
type DSTORE_1 struct{ base.NoOperandsInstruction }
type DSTORE_2 struct{ base.NoOperandsInstruction }
type DSTORE_3 struct{ base.NoOperandsInstruction }

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()

	frame.LocalVars().SetDouble(index, val)
}

func (d *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, d.Index)
}

func (d *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

func (d *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

func (d *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

func (d *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
