package constants

import (
	"github/Xpectuer/jvmgo/instructions/base"
	"github/Xpectuer/jvmgo/rtda"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

type BIPUSH struct{ val int8 } // Push byte

func (bip *BIPUSH) FetchOperand(reader *base.BytecodeReader) {
	bip.val = reader.ReadInt8()
}

func (bip *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(bip.val)
	os := frame.OperandStack()
	os.PushInt(i)
}

type SIPUSH struct{ val int16 } // Push short

func (sip *SIPUSH) FetchOperand(reader *base.BytecodeReader) {
	sip.val = reader.ReadInt16()
}

func (sip *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(sip.val)
	os := frame.OperandStack()
	os.PushInt(i)
}
