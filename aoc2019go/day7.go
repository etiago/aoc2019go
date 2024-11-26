package aoc2019go

import (
	"fmt"
	"sync"

	"gonum.org/v1/gonum/stat/combin"
)

func part1(original_instructions []int) int {
	// Part 1
	var wg sync.WaitGroup

	const NumComputers = 5
	// Slice to store input and output channels for each computer
	channels := make([]chan int, NumComputers+1)
	for i := 0; i <= NumComputers; i++ {
		channels[i] = make(chan int)
	}

	highest_output := 0
	phase_settings := combin.Permutations(NumComputers, NumComputers)

	for _, phase_setting := range phase_settings {
		for i := 0; i < NumComputers; i++ {
			inputChannelIndex := i
			outputChannelIndex := i + 1

			instructions := make([]int, len(original_instructions))
			copy(instructions, original_instructions)

			programState := ProgramState{instructions, 0}

			wg.Add(1)
			go programState.ExecuteToHaltAsync(
				channels[inputChannelIndex],
				channels[outputChannelIndex],
				&wg)
		}

		for j := 0; j < NumComputers; j++ {
			channels[j] <- phase_setting[j]
		}

		channels[0] <- 0

		// Read output from last computer
		output := <-channels[5]

		if output > highest_output {
			highest_output = output
		}
		wg.Wait()
	}

	fmt.Print("Part 1: ")
	fmt.Println(highest_output)

	// Close all channels
	for i := 0; i <= NumComputers; i++ {
		close(channels[i])
	}

	return highest_output
}

func Day7() string {
	input := ReadPuzzleFile("inputs/day7.txt")
	original_instructions := ParsePuzzleToIntArray(input)

	part1(original_instructions)

	return "Done"
}
