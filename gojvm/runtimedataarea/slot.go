package runtimedataarea

// 局部变量表是按索引访问的，所以很自然，可以把它想象成一个数组。
// 根据Java虚拟机规范，这个数组的每个元素至少可以容纳一个int或引用值，两个连续的元素可以容纳一个long或double值。
type Slot struct {
	num int32   // 存放整数
	ref *Object // 存放引用
}
