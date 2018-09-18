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

// dload系列操作double类型变量

type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)

	frame.OperandStack().PushDouble(val)
}

func (d *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, d.Index)
}

func (d *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (d *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (d *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (d *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
