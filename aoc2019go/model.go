package aoc2019go

type ProgramState struct {
	memory             []int
	instructionPointer int
}

type InOutArgs struct {
	in     []int
	out    int
	hasOut bool
}
