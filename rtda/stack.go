package rtda

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

type Stack struct {
	// capacity of the stack
	maxSize uint
	// current size of the stack
	size uint
	// top pointer to the stack
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (ls *Stack) push(frame *Frame) {
	if ls.size >= ls.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if ls._top != nil {
		// insert the frame node
		frame.lower = ls._top
	}
	ls._top = frame
	ls.size++
}

func (ls *Stack) pop() *Frame {
	if ls._top == nil {
		panic("pop(): jmv stack empty!")
	}

	top := ls._top
	ls._top = top.lower
	// dettach the top frame
	top.lower = nil
	ls.size--
	return top
}

func (ls *Stack) top() *Frame {
	if ls._top == nil {
		panic("top(): jvm stack empty")
	}
	return ls._top
}
