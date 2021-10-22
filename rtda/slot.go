package rtda

// 通过uintptr实现的slot数组会被垃圾回收，因此采用Operand与引用复合的结构体Slot实现局部变量表
type Slot struct {
	num int32
	ref *Object
}
