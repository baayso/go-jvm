/*
	加载指令从局部变量表获取变量，然后推入操作数栈顶。
	加载指令共33条，按照所操作变量的类型可以分为6类：
	aload系列指令操作引用类型变量、
	dload系列操作double类型变量、
	fload系列操作float变量、
	iload系列操作int变量、
	lload系列操作long变量、
	xaload操作数组。
*/
package loads

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// fload系列操作float变量

type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)

	frame.OperandStack().PushFloat(val)
}

func (f *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, f.Index)
}

func (f *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (f *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (f *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (f *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
