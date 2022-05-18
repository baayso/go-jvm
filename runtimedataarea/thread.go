package runtimedataarea

// 线程私有的运行时数据区用于辅助执行Java字节码。
// 每个线程都有自己的pc寄存器（Program Counter）和Java虚拟机栈（JVM Stack）。
// Java虚拟机栈又由栈帧（Stack Frame，后面简称帧）构成，
// 帧中保存方法执行的状态，包括局部变量表（Local Variable）和操作数栈（Operand Stack）等。
// 在任一时刻，某一线程肯定是在执行某个方法。
// 这个方法叫作该线程的当前方法；执行该方法的帧叫作线程的当前帧；声明该方法的类叫作当前类。
// 如果当前方法是Java方法，则pc寄存器中存放当前正在执行的Java虚拟机指令的地址，否则，当前方法是本地方法，pc寄存器中的值没有明确定义。

// 实现线程私有的运行时数据区
type Thread struct {
	pc    int    // pc寄存器（Program Counter）
	stack *Stack // Java虚拟机栈（JVM Stack），Java虚拟机栈又由栈帧（Stack Frame，简称帧）
}

// 创建Thread实例
func NewThread() *Thread {
	return &Thread{
		// todo: 修改命令行工具，添加-Xss选项来指定这个参数
		stack: newStack(1024),
	}
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

// 返回当前帧
func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

// getter

func (t *Thread) PC() int {
	return t.pc
}

// setter

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}
