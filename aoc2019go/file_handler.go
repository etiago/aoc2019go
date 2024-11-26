package aoc2019go

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadPuzzleFile(file_path string) string {
	content, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ParsePuzzleToIntArray(puzzle string) []int {
	instructions_strs := strings.Split(puzzle, ",")
	instructions := make([]int, len(instructions_strs))

	for i, instruction_str := range instructions_strs {
		instruction, err := strconv.Atoi(instruction_str)
		if err != nil {
			log.Fatal(err)
		}

		instructions[i] = instruction
	}
	return instructions
}
