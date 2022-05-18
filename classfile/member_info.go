package classfile

// 接口索引表之后是字段表和方法表，分别存储字段和方法信息。字段和方法的基本结构大致相同，差别仅在于属性表。
// 下面是Java虚拟机规范给出的字段结构定义。
/*
field_info {
    u2                access_flags;
    u2                name_index;
    u2                descriptor_index;
    u2                attributes_count;
    attribute_info    attributes[attributes_count];
}
*/
// 和类一样，字段和方法也有自己的访问标志。
// 访问标志之后是一个常量池索引，给出字段名或方法名，然后又是一个常量池索引，给出字段或方法的描述符，最后是属性表。

// MemberInfo : 为了避免重复代码，用一个结构体统一表示字段和方法
type MemberInfo struct {
	cp              ConstantPool // 保存常量池指针
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// 读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)

	for i := range members {
		members[i] = readMember(reader, cp)
	}

	return members
}

// 读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}

// Name : 从常量池查找字段或方法名
func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}

// Descriptor : 从常量池查找字段或方法描述符
func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}
