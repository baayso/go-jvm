package classfile

// ExceptionTableEntry : 异常处理
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// CodeAttribute : Code是变长属性，只存在于method_info结构中。Code属性中存放字节码等方法相关信息。
// 其结构定义如下：
/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
// max_stack给出操作数栈的最大深度，max_locals给出局部变量表大小。
// 接着是字节码，存在u1表中。最后是异常处理表和属性表。
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {
	c.maxStack = reader.readUint16()
	c.maxLocals = reader.readUint16()

	codeLength := reader.readUint32()
	c.code = reader.readByte(codeLength)

	c.exceptionTable = readExceptionTable(reader)
	c.attributes = readAttributes(reader, c.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()

	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}

	return exceptionTable
}
