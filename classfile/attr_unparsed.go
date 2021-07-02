package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (upa *UnparsedAttribute) readInfo(reader *ClassReader) {
	upa.info = reader.readBytes(upa.length)
}
