package base

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
import "github/Xpectuer/jvmgo/rtda"

type Instruction interface {
	FetchOperand(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (noi *NoOperandsInstruction) FetchOperand(reader *BytecodeReader) {
	// do nothing
}

type BranchInstruction struct {
	Offset int
}

func (bi *BranchInstruction) FetchOperand(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}

// store and load
type Index8Instruction struct {
	Index uint
}

// store and load instructions only have 1 byte index of operand
func (i8i *Index8Instruction) FetchOperand(reader *BytecodeReader) {
	i8i.Index = uint(reader.ReadUint8())
}

// store and load
type Index16Instruction struct {
	Index uint
}

// store and load instructions only have 1 byte index of operand
func (i16i *Index16Instruction) FetchOperand(reader *BytecodeReader) {
	i16i.Index = uint(reader.ReadUint16())
}
