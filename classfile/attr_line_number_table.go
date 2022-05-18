package classfile

// LineNumberTableAttribute : LineNumberTable属性表存放方法的行号信息，
// 和SourceFile属性一样都属于调试信息，都不是运行时必需的。
// 在使用javac编译器编译Java程序时，默认会在class文件中生成这些信息。
// 可以使用javac提供的-g：none选项来关闭这些信息的生成，具体请参考javac用法。
// 结构定义如下：
/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (a *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()

	a.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)

	for i := range a.lineNumberTable {
		a.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
