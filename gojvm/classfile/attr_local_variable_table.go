package classfile

// LocalVariableTable属性表中存放方法的局部变量信息。
// 和SourceFile属性一样都属于调试信息，都不是运行时必需的。
// 在使用javac编译器编译Java程序时，默认会在class文件中生成这些信息。
// 可以使用javac提供的-g：none选项来关闭这些信息的生成，具体请参考javac用法。
// 结构定义如下：
/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (this *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()

	this.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)

	for i := range this.localVariableTable {
		this.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
