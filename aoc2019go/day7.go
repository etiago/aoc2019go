package aoc2019go

import (
	"fmt"
	"sync"

	"gonum.org/v1/gonum/stat/combin"
)

func initializeComputers() {

}
func Day7() string {
	input := ReadPuzzleFile("inputs/day7.txt")
	original_instructions := ParsePuzzleToIntArray(input)

	// Part 1
	var wg sync.WaitGroup

	const NumComputers = 5
	// Slice to store input and output channels for each computer
	channels := make([]chan int, NumComputers+1)

	for i := 0; i <= NumComputers; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < NumComputers; i++ {
		inputChannelIndex := i
		outputChannelIndex := i + 1

		instructions := make([]int, len(original_instructions))
		copy(instructions, original_instructions)

		// instructions[1] = 12
		// instructions[2] = 2

		programState := ProgramState{instructions, 0}

		wg.Add(1)
		go programState.ExecuteToHaltAsync(
			channels[inputChannelIndex],
			channels[outputChannelIndex],
			&wg)
	}

	phase_settings := combin.Permutations(NumComputers, NumComputers)

	fmt.Println("About to send phase settings...")
	for _, phase_setting := range phase_settings {
		fmt.Println(phase_setting)
		for j := 0; j < NumComputers; j++ {
			channels[j] <- phase_setting[j]
		}

		channels[0] <- 0

		// Read output from last computer
		fmt.Println("About to read output...")
		output := <-channels[5]
		fmt.Println(output)
	}
	wg.Wait()

	return "Done"
	// // Part 2
	// copy(instructions, original_instructions)
	// programState = ProgramState{instructions, 0}

	// return input
}
