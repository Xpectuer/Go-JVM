package base

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

type BytecodeReader struct {
	code []byte
	pc   int
}

// avoid creating a new BOR instance each time ins is decoded
func (bor *BytecodeReader) Reset(code []byte, pc int) {
	bor.code = code
	bor.pc = pc
}

func (bor *BytecodeReader) ReadUint8() uint8 {
	// fetch byte code
	i := bor.code[bor.pc]
	bor.pc++
	return i
}

func (bor *BytecodeReader) ReadUint16() uint16 {
	b1 := uint16(bor.ReadUint8())
	b2 := uint16(bor.ReadUint8())
	return (b1 << 8) | b2
}

func (bor *BytecodeReader) ReadInt16() int16 {
	return int16(bor.ReadUint16())
}

func (bor *BytecodeReader) ReadInt32() int32 {
	b1 := int32(bor.ReadUint8())
	b2 := int32(bor.ReadUint8())
	b3 := int32(bor.ReadUint8())
	b4 := int32(bor.ReadUint8())

	return (b1 << 24) | (b2 << 16) | (b3 << 8) | b4
}

func (bor *BytecodeReader) ReadInt32s() []int {
	return []int{}
}

func (bor *BytecodeReader) SkipPadding() {
	return
}
