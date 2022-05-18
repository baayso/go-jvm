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
	"github.com/baayso/go-jvm/instructions/base"
	rtda "github.com/baayso/go-jvm/runtimedataarea"
)

// aload系列指令操作引用类型变量

type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)

	frame.OperandStack().PushRef(ref)
}

func (a *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, a.Index)
}

func (a *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (a *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (a *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (a *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
