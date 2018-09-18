/*
	存储指令与加载指令刚好相反，存储指令把变量从操作数栈顶弹出，然后存入局部变量表。
	和加载指令一样，存储指令也可以分为6类。
*/
package stores

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()

	frame.LocalVars().SetRef(index, ref)
}

func (a *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, a.Index)
}

func (a *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

func (a *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (a *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (a *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
