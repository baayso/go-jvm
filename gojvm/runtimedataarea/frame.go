package runtimedataarea

type Frame struct {
	lower        *Frame        // 指向下一个Frame（用于实现链表数据结构）
	localVars    LocalVars     // 局部变量表指针
	operandStack *OperandStack // 操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	// 执行方法所需的局部变量表大小和操作数栈深度是由编译器预先计算好的，存储在class文件method_info结构的Code属性中。
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}
