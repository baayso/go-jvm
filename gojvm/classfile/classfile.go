package classfile

import "fmt"

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

// 将[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}

	cf = &ClassFile{}
	cf.read(cr)

	return
}

// 依次调用其他方法解析class文件
func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)

	c.constantPool = readConstantPool(reader)

	// 常量池之后是类访问标志，这是一个16位的“bitmask”，指出class文件定义的是类还是接口，访问级别是public还是private，等等。
	c.accessFlags = reader.readUint16()

	// 类访问标志之后是两个u2类型的常量池索引，分别给出类名和超类名。
	// class文件存储的类名类似完全限定名，但是把点换成了斜线，Java语言规范把这种名字叫作二进制名（binary names）。
	// 因为每个类都有名字，所以thisClass必须是有效的常量池索引。
	// 除java.lang.Object之外，其他类都有超类，所以superClass只在Object.class中是0，
	// 在其他class文件中必须是有效的常量池索引。
	c.thisClass = reader.readUint16()
	c.superClass = reader.readUint16()

	// 类和超类索引后面是接口索引表，表中存放的也是常量池索引，给出该类实现的所有接口的名字。
	c.interfaces = reader.readUint16s()

	c.fields = readMembers(reader, c.constantPool)
	c.methods = readMembers(reader, c.constantPool)
	c.attributes = readAttributes(reader, c.constantPool)
}

// Java虚拟机规范规定，如果加载的class文件不符合要求的格式，Java虚拟机实现就抛出java.lang.ClassFormatError异常
func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32() // 读4个字节

	// // class文件的魔数是“0xCAFEBABE”
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 魔数之后是class文件的次版本号和主版本号，都是u2类型。
// 假设某class文件的主版本号是M，次版本号是m，那么完整的版本号可以表示成“M.m”的形式。
// 次版本号只在J2SE 1.2之前用过，从1.2开始基本上就没什么用了（都是0）。
// 主版本号在J2SE 1.2之前是45，从1.2开始，每次有大的Java版本发布，都会加1。
//
// 特定的Java虚拟机实现只能支持版本号在某个范围内的class文件。
// Oracle的实现是完全向后兼容的，比如Java SE 8支持版本号为45.0~52.0的class文件。
// 如果版本号不在支持的范围内，Java虚拟机实现就抛出java.lang.UnsupportedClassVersionError异常。
// 我们参考Java 8，支持版本号为45.0~52.0的class文件。
func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	c.minorVersion = reader.readUint16()
	c.majorVersion = reader.readUint16()

	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

// 从常量池查找类名
func (c *ClassFile) ClassName() string {
	return c.constantPool.getClassName(c.thisClass)
}

// 从常量池查找超类名
func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}

	return "" // 只有 java.lang.Object 没有超类
}

// 从常量池查找接口名
func (c *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(c.interfaces))

	for i, cpIndex := range c.interfaces {
		interfaceNames[i] = c.constantPool.getClassName(cpIndex)
	}

	return interfaceNames
}

// getter

func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}
