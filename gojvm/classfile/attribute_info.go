package classfile

// 和常量池类似，各种属性表达的信息也各不相同，因此无法用统一的结构来定义。
// 不同之处在于，常量是由Java虚拟机规范严格定义的，共有14种。
// 但属性是可以扩展的，不同的虚拟机实现可以定义自己的属性类型。
// 由于这个原因，Java虚拟机规范没有使用tag，而是使用属性名来区别不同的属性。
// 属性数据放在属性名之后的u1表中，这样Java虚拟机实现就可以跳过自己无法识别的属性。
// 属性的结构定义如下：
/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
// 注意：属性表中存放的属性名实际上并不是编码后的字符串，而是常量池索引，指向常量池中的CONSTANT_Utf8_info常量。
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}
