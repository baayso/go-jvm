package classfile

// ConstantStringInfo : CONSTANT_String_info常量表示java.lang.String字面量，结构如下：
/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
// CONSTANT_String_info本身并不存放字符串数据，只存了常量池索引，
// 这个索引指向一个CONSTANT_Utf8_info常量。
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// 读取常量池索引
func (c *ConstantStringInfo) readInfo(reader *ClassReader) {
	c.stringIndex = reader.readUint16()
}

// 按索引从常量池中查找字符串
func (c *ConstantStringInfo) String() string {
	return c.cp.getUtf8(c.stringIndex)
}
