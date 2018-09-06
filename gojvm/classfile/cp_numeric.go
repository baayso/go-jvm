package classfile

import "math"

// CONSTANT_Integer_info使用4字节存储整数常量，其结构定义如下：
/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
// CONSTANT_Integer_info正好可以容纳一个Java的int型常量，
// 但实际上比int更小的boolean、byte、short和char类型的常量也放在CONSTANT_Integer_info中。
type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = int32(bytes)
}

// CONSTANT_Float_info使用4字节存储IEEE754单精度浮点数常量，结构如下：
/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = math.Float32frombits(bytes)
}

// CONSTANT_Long_info使用8字节存储整数常量，结构如下：
/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = int64(bytes)
}

// CONSTANT_Double_info使用8字节存储IEEE754双精度浮点数，结构如下：
/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = math.Float64frombits(bytes)
}
