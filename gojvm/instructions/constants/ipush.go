package constants

import (
	"github.com/baayso/jvm/gojvm/instructions/base"
	rtda "github.com/baayso/jvm/gojvm/runtimedataarea"
)

// bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶。
type BIPUSH struct {
	val int8
}

func (b *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	b.val = reader.ReadInt8()
}

func (b *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(b.val)

	frame.OperandStack().PushInt(i)
}

// sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶。
type SIPUSH struct {
	val int16
}

func (s *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	s.val = reader.ReadInt16()
}

func (s *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(s.val)

	frame.OperandStack().PushInt(i)
}
