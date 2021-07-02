package classfile

type ConstantMethodHandleInfo struct {
	referenceKind uint8
	referenceIndex uint16 
}

func (cmhi *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	cmhi.referenceKind = reader.readUint8()
	cmhi.referenceIndex = reader.readUint16()
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16 
}

func (cmti *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	cmti.descriptorIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16 
	nameAndTypeIndex uint16
}

func (cidi *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	cidi.bootstrapMethodAttrIndex = reader.readUint16()
	cidi.nameAndTypeIndex = reader.readUint16()
}



