package loads

import (
	"github/Xpectuer/jvmgo/instructions/base"
	"github/Xpectuer/jvmgo/rtda"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

// method to reuse
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	os := frame.OperandStack()
	os.PushInt(val)
}

type ILOAD struct{ base.Index8Instruction }

func (ild *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(ild.Index))
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (ild *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (ild *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (ild *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
