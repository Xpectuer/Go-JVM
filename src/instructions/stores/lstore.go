package stores

import (
	"github/Xpectuer/jvmgo/instructions/base"
	"github/Xpectuer/jvmgo/rtda"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

func _lstore(frame *rtda.Frame, index uint) {
	os := frame.OperandStack()
	val := os.PopLong()
	frame.LocalVars().SetLong(index, val)
}

type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }
