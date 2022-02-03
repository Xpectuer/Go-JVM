package base

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

type BytecodeReader struct {
	code []byte
	pc   int
}

// avoid creating a new bcr instance each time ins is decoded
func (bcr *BytecodeReader) Reset(code []byte, pc int) {
	bcr.code = code
	bcr.pc = pc
}

func (bcr *BytecodeReader) ReadUint8() uint8 {
	// fetch byte code
	i := bcr.code[bcr.pc]
	bcr.pc++
	return i
}

func (bcr *BytecodeReader) ReadInt8() int8 {
	return int8(bcr.ReadUint8())
}


func (bcr *BytecodeReader) ReadUint16() uint16 {
	b1 := uint16(bcr.ReadUint8())
	b2 := uint16(bcr.ReadUint8())
	return (b1 << 8) | b2
}

func (bcr *BytecodeReader) ReadInt16() int16 {
	return int16(bcr.ReadUint16())
}

func (bcr *BytecodeReader) ReadInt32() int32 {
	b1 := int32(bcr.ReadUint8())
	b2 := int32(bcr.ReadUint8())
	b3 := int32(bcr.ReadUint8())
	b4 := int32(bcr.ReadUint8())

	return (b1 << 24) | (b2 << 16) | (b3 << 8) | b4
}

func (bcr *BytecodeReader) ReadInt32s() []int {
	return []int{}
}

func (bcr *BytecodeReader) SkipPadding() {
	return
}
