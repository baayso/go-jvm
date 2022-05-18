package runtimedataarea

import "math"

// 操作数栈的实现方式和局部变量表类似
// 操作数栈的大小是编译器已经确定的，所以可以用[]Slot实现。size字段用于记录栈顶位置。
type OperandStack struct {
	size  uint // 栈顶位置
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}

	return nil
}

// int

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].num
}

// float变量先转成int类型，然后按int变量处理

func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)

	o.PushInt(int32(bits))
}

func (o *OperandStack) PopFloat() float32 {
	bits := uint32(o.PopInt())

	return math.Float32frombits(bits)
}

// 把long变量推入栈顶时，要拆成两个int变量。
// 弹出时，先弹出两个int变量，然后组装成一个long变量

func (o *OperandStack) PushLong(val int64) {
	o.PushInt(int32(val))
	o.PushInt(int32(val >> 32))
}

func (o *OperandStack) PopLong() int64 {
	high := uint32(o.PopInt()) // 取high值要在取low值之前
	low := uint32(o.PopInt())

	return int64(high)<<32 | int64(low)
}

// double变量先转成long类型，然后按long变量处理

func (o *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)

	o.PushLong(int64(bits))
}

func (o *OperandStack) PopDouble() float64 {
	bits := uint64(o.PopLong())

	return math.Float64frombits(bits)
}

// 引用类型

func (o *OperandStack) PushRef(ref *Object) {
	o.slots[o.size].ref = ref
	o.size++
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	ref := o.slots[o.size].ref

	// 弹出引用后，把Slot结构体的ref字段设置成nil，这是为了帮助Go的垃圾收集器回收Object结构体实例
	o.slots[o.size].ref = nil

	return ref
}

func (o *OperandStack) PushSlot(slot Slot) {
	o.slots[o.size] = slot
	o.size++
}

func (o *OperandStack) PopSlot() Slot {
	o.size--
	return o.slots[o.size]
}
