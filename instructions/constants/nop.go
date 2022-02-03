package constants

import (
	"github/Xpectuer/jvmgo/instructions/base"
	"github/Xpectuer/jvmgo/rtda"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

type NOP struct{ base.NoOperandsInstruction }

func (nop *NOP) Execute(frame *rtda.Frame) {
	// do
}
