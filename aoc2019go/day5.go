package aoc2019go

func Day5() string {
	input := ReadPuzzleFile("inputs/day5.txt")
	original_instructions := ParsePuzzleToIntArray(input)

	instructions := make([]int, len(original_instructions))
	copy(instructions, original_instructions)

	// Part 1
	programState := ProgramState{instructions, 0}

	programState.ExecuteToHalt()

	// Part 2
	copy(instructions, original_instructions)
	programState = ProgramState{instructions, 0}

	programState.ExecuteToHalt()

	return input
}
