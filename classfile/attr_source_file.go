package classfile

// SourceFileAttribute : SourceFile是可选定长属性，只会出现在ClassFile结构中，用于指出源文件名。
// 其结构定义如下：
/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
// attribute_length的值必须是2。sourcefile_index是常量池索引，指向CONSTANT_Utf8_info常量。
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (s *SourceFileAttribute) readInfo(reader *ClassReader) {
	s.sourceFileIndex = reader.readUint16()
}

func (s *SourceFileAttribute) FileName() string {
	return s.cp.getUtf8(s.sourceFileIndex)
}
