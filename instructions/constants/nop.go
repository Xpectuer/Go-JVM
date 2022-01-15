package constants

import (
	"github/Xpectuer/jvmgo/rtda"
	"github/Xpecuter/jvmgo/base"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

type NOP struct{ base.NoOperandsInstruction }

func (nop *NOP) Execute(frame *rtda.Frame) {
	// do
}
