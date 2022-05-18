package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// 读取 u1(1字节)类型数据
func (c *ClassReader) readUint8() uint8 {
	val := c.data[0]

	// ClassReader并没有使用索引记录数据位置，而是使用reslice语法跳过已经读取的数据
	c.data = c.data[1:]

	return val
}

// 读取 u2(2字节)类型数据
func (c *ClassReader) readUint16() uint16 {
	// BigEndian可以从[]byte中解码多字节数据
	val := binary.BigEndian.Uint16(c.data)

	c.data = c.data[2:]

	return val
}

// 读取 u4(4字节)类型数据
func (c *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(c.data)

	c.data = c.data[4:]

	return val
}

// 读取uint64（Java虚拟机规范并没有定义u8）类型数据
func (c *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(c.data)

	c.data = c.data[8:]

	return val
}

// 读取uint16表，表的大小由开头的uint16数据指出
func (c *ClassReader) readUint16s() []uint16 {
	n := c.readUint16()

	s := make([]uint16, n)

	for i := range s {
		s[i] = c.readUint16()
	}

	return s
}

// 读取指定数量的字节
func (c *ClassReader) readByte(n uint32) []byte {
	bytes := c.data[:n]

	c.data = c.data[n:]

	return bytes
}
