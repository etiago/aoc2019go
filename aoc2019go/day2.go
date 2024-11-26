package aoc2019go

import "fmt"

func Day2() string {
	input := ReadPuzzleFile("inputs/day2.txt")
	original_instructions := ParsePuzzleToIntArray(input)

	instructions := make([]int, len(original_instructions))
	copy(instructions, original_instructions)

	// Part 1
	instructions[1] = 12
	instructions[2] = 2

	halted := false
	programState := ProgramState{instructions, 0}

	for !halted {
		programState.executeInstruction()
		halted = programState.isHalted()
	}
	fmt.Print("Part 1: ")
	fmt.Println(programState.memory[0])

	// Part 2
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(instructions, original_instructions)
			instructions[1] = noun
			instructions[2] = verb

			halted = false

			programState = ProgramState{instructions, 0}

			for !halted {
				programState.executeInstruction()
				halted = programState.isHalted()
			}

			if programState.memory[0] == 19690720 {
				fmt.Print("Part 2: ")
				fmt.Println(100*noun + verb)
				break
			}
		}
	}

	return input
}
