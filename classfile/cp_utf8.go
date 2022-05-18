package classfile

// ConstantUtf8Info : CONSTANT_Utf8_info常量里放的是MUTF-8编码的字符串，结构如下：
/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
// 字符串在class文件中是以MUTF-8（Modified UTF-8）方式编码的。
// MUTF-8编码方式和UTF-8大致相同，但并不兼容。差别有两点：
// 一是null字符（代码点U+0000）会被编码成2字节：0xC0、0x80；
// 二是补充字符（Supplementary Characters，代码点大于U+FFFF的Unicode字符）是按UTF-16拆分为代理对（Surrogate Pair）分别编码的。
type ConstantUtf8Info struct {
	str string
}

func (c *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readByte(length)
	c.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
