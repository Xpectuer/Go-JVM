package classfile
// 给出字段名字和描述符
// 如 变量的类型和名字、方法的签名（参数、返回值）与 方法名

// 用 名字+描述符 唯一确定一个 变量或方法 
// 因此方法可以被重载
type ConstantNameAndTypeInfo struct {
	nameIndex uint16 
	descriptionIndex uint16 
}