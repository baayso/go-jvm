package classfile

// ConstantMemberRefInfo : CONSTANT_Fieldref_info表示字段符号引用。
// CONSTANT_Methodref_info表示普通（非接口）方法符号引用。
// CONSTANT_InterfaceMethodref_info表示接口方法符号引用。
// 这三种常量结构一模一样，结构如下：
/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
// class_index和name_and_type_index都是常量池索引，分别指向CONSTANT_Class_info和CONSTANT_NameAndType_info常量。
//
// 定义一个统一的结构体ConstantMemberrefInfo来表示这3种常量：
type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *ConstantMemberRefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

// ConstantFieldrefInfo : CONSTANT_Fieldref_info表示字段符号引用。
type ConstantFieldrefInfo struct {
	ConstantMemberRefInfo
}

// ConstantMethodrefInfo : CONSTANT_Methodref_info表示普通（非接口）方法符号引用。
type ConstantMethodrefInfo struct {
	ConstantMemberRefInfo
}

// ConstantInterfaceMethodrefInfo : CONSTANT_InterfaceMethodref_info表示接口方法符号引用。
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberRefInfo
}
