package constants

import (
	"github/Xpectuer/jvmgo/rtda"
	"github/Xpecuter/jvmgo/base"
)

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
// mock inheritation

/*
 ???: Invalid Method Expr
 func (an *ACONST_NULL) Execute(frame *rtda.Frame) {
 	frame.OperandStack().PushRef(nil)
 }
*/

// aconst_null pushes a nil reference onto the top of operand stack
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (an *ACONST_NULL) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushRef(nil)
}

// dconst_0 pushes a double 0 onto the top of operand stack
type DCONST_0 struct{ base.NoOperandsInstruction }

func (dc0 *DCONST_0) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushDouble(0.0)
}

// dconst_1 pushes a double 1.0 onto the top of operand stack
type DCONST_1 struct{ base.NoOperandsInstruction }

func (dc1 *DCONST_1) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushDouble(1.0)
}

// fconst_0 pushes a float 0.0 onto the top of operand stack
type FCONST_0 struct{ base.NoOperandsInstruction }

func (fc0 *FCONST_0) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

func (fc1 *FCONST_1) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperandsInstruction }

func (fc2 *FCONST_2) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushFloat(2.0)
}

type ICONST_M1 struct{ base.NoOperandsInstruction }

func (icm1 *ICONST_M1) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (ic0 *ICONST_0) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (ic1 *ICONST_1) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (ic2 *ICONST_2) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (ic3 *ICONST_3) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (ic4 *ICONST_4) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (ic5 *ICONST_5) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushInt(5)
}

type LCONST_0 struct{ base.NoOperandsInstruction }

func (lc0 *LCONST_0) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (lc1 *LCONST_1) Execute(frame *rtda.Frame) {
	os := frame.OperandStack()
	os.PushLong(1)
}
