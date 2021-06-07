package classfile

type ConstantUtf8Info struct {
	str string
}

func (cui *ConstantUtf8Info) readInfo(reader *ClassReader) {
	len := uint32(reader.readUint16())
	bytes := reader.readBytes(len)
	cui.str = decodeMUTF8(bytes)
}

// TODO：完善MUTF8
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}


