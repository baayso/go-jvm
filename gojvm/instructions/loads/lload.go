package loads

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// 加载指令从局部变量表获取变量，然后推入操作数栈顶。
// lload系列操作long变量

type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)

	frame.OperandStack().PushLong(val)
}

func (l *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, l.Index)
}

func (l *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (l *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (l *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (l *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
