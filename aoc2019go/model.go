package aoc2019go

type ProgramState struct {
	computer_id        int
	memory             []int
	instructionPointer int
	halted             bool
	last_output        int
}

type InOutArgs struct {
	in     []int
	out    int
	hasOut bool
}
