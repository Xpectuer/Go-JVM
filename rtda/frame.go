package rtda

type Frame struct {
	// next pointer
	lower *Frame
	// size of local variables is predefined by compiler
	// stored in the method_info structure (see classfile/attr_code.go)
	localVars    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
