package classfile

// CONSTANT_Class_info常量表示类或者接口的符号引用，结构如下：
/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
// 和CONSTANT_String_info类似，name_index是常量池索引，指向CONSTANT_Utf8_info常量。
// 类和超类索引，以及接口表中的接口索引指向的都是CONSTANT_Class_info常量
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

// 读取常量池索引
func (c *ConstantClassInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
}

// 按索引从常量池中查找字符串
func (c *ConstantClassInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}
