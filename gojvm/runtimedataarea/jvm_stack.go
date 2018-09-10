package runtimedataarea

// 用链表（Linked List）数据结构来实现Java虚拟机栈，
// 这样栈就可以按需使用内存空间，而且弹出的帧也可以及时被Go的垃圾收集器回收。
type Stack struct {
	maxSize uint   // 栈的容量（最多可以容纳多少帧）
	size    uint   // 栈的当前大小
	_top    *Frame // 栈顶指针
}

// 创建Stack结构体实例，参数表示要创建的Stack最多可以容纳多少帧
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 把帧推入栈顶
func (s *Stack) push(frame *Frame) {
	// 按照Java虚拟机规范，应该抛出StackOverflowError异常
	if s.size > s.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if s._top != nil {
		frame.lower = s._top
	}

	s._top = frame
	s.size++
}

// 把栈顶帧弹出
func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}

	top := s._top
	s._top = top.lower
	top.lower = nil
	s.size--

	return top
}

// 返回栈顶帧，但并不弹出
func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}

	return s._top
}
