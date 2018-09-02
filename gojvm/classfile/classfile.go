package classfile

// 构成class文件的基本数据单位是字节，可以把整个class文件当成一个字节流来处理。
// 稍大一些的数据由连续多个字节构成，这些数据在class文件中以大端（big-endian）方式存储。
// 为了描述class文件格式，Java虚拟机规范定义了u1、u2和u4三种数据类型来表示1、2和4字节无符号整数，
// 分别对应Go语言的uint8、uint16和uint32类型。
// 相同类型的多条数据一般按表（table）的形式存储在class文件中。
// 表由表头和表项（item）构成，表头是u2或u4整数。假设表头是n，后面就紧跟着n个表项数据。

// Java虚拟机规范使用一种类似C语言的结构体语法来描述class文件格式。
// 整个class文件被描述为一个ClassFile结构，代码如下：
/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// Java虚拟机规范定义的class文件格式
type ClassFile struct {
	// magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}
