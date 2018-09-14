package loads

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// 加载指令从局部变量表获取变量，然后推入操作数栈顶。
// iload系列操作int变量

type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)

	frame.OperandStack().PushInt(val)
}

func (i *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, i.Index)
}

func (i *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (i *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (i *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (i *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
