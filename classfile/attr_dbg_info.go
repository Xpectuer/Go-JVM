package classfile


// 用于提示 Debug信息中的行号与字节码行数对应关系
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (lnta *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	lnta.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lnta.lineNumberTable {
		lnta.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}


// 方法的局部变量描述
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}
 
func (lvta *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	lvta.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)

	for i := range lvta.localVariableTable {
		lvta.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}