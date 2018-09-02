package classfile

// 常量池实际上也是一个表，但是有三点需要特别注意。
// 第一，表头给出的常量池大小比实际大1。假设表头给出的值是n，那么常量池的实际大小是n–1。
// 第二，有效的常量池索引是1~n–1。0是无效索引，表示不指向任何常量。
// 第三，CONSTANT_Long_info和CONSTANT_Double_info各占两个位置。
// 也就是说，如果常量池中存在这两种常量，实际的常量数量比n–1还要少，而且1~n–1的某些数也会变成无效索引。
type ConstantPool []ConstantInfo