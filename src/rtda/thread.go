package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	//  TODO: parameterized default stack size
	return &Thread{
		stack: newStack(1024),
	}
}

// return the Program Counter of the Thread
func (thread *Thread) PC() int { return thread.pc }

// set PC for branching or instruction execution
func (thread *Thread) SetPC(pc int) { thread.pc = pc }

// push the method frame
func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

// pop the method frame
func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

// get current frame
func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top()
}
