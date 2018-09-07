package classfile

// Deprecated和Synthetic是最简单的两种属性，仅起标记作用，不包含任何数据。
// 这两种属性都是JDK1.1引入的，可以出现在ClassFile、field_info和method_info结构中：
// 它们的结构定义如下：
/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}

Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
// 由于不包含任何数据，所以attribute_length的值必须是0
type MarkerAttribute struct{}

// “标记”属性没有数据，所以readInfo()方法是空的
func (m *MarkerAttribute) readInfo(reader *ClassReader) {}

// Deprecated属性用于指出类、接口、字段或方法已经不建议使用，编译器等工具可以根据Deprecated属性输出警告信息。
// J2SE 5.0之前可以使用Javadoc提供的@deprecated标签指示编译器给类、接口、字段或方法添加Deprecated属性。
// 从J2SE 5.0开始，也可以使用@Deprecated注解。
type DeprecatedAttribute struct {
	MarkerAttribute
}

// Synthetic属性用来标记源文件中不存在、由编译器生成的类成员，
// 引入Synthetic属性主要是为了支持嵌套类和嵌套接口。
// 具体细节请参考Java虚拟机规范。
type SyntheticAttribute struct {
	MarkerAttribute
}
