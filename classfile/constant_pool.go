package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())

	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		// 从常量池索引1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		// Long类型Info，与Double类型Info多一位
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp

}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid Constant Pool Index!")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptionIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

