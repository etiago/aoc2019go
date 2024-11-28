package aoc2019go

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func (programState *ProgramState) extractArgs(numArgs int, modes int) []int {
	args := make([]int, numArgs)

	for i := 0; i < numArgs; i++ {
		arg := programState.memory[programState.instructionPointer+i+1]
		if modes%10 == 0 {
			args[i] = programState.memory[arg]
		} else {
			args[i] = arg
		}
		modes /= 10
	}

	return args
}

func (programState *ProgramState) extractInOutArgs() InOutArgs {
	instruction := programState.memory[programState.instructionPointer]
	opcode := instruction % 100
	modes := instruction / 100

	var in []int
	var out int
	var hasOut bool

	switch opcode {
	case 1:
		in = programState.extractArgs(2, modes)
		out = programState.memory[programState.instructionPointer+3]
		hasOut = true
	case 2:
		in = programState.extractArgs(2, modes)
		out = programState.memory[programState.instructionPointer+3]
		hasOut = true
	case 3:
		out = programState.memory[programState.instructionPointer+1]
		hasOut = true
	case 4:
		in = programState.extractArgs(1, modes)
	case 5:
		in = programState.extractArgs(2, modes)
	case 6:
		in = programState.extractArgs(2, modes)
	case 7:
		in = programState.extractArgs(2, modes)
		out = programState.memory[programState.instructionPointer+3]
		hasOut = true
	case 8:
		in = programState.extractArgs(2, modes)
		out = programState.memory[programState.instructionPointer+3]
		hasOut = true
	case 99:
		return InOutArgs{}
	default:
		panic("Invalid opcode")
	}

	return InOutArgs{in, out, hasOut}
}

func (programState *ProgramState) ExecuteToHalt() {
	for !programState.isHalted() {
		programState.executeInstruction()
	}
}

func (programState *ProgramState) ExecuteToHaltAsync(computer_id int, in <-chan int, out chan<- int, wg *sync.WaitGroup, out_final *chan int) {
	defer wg.Done()

	for !programState.isHalted() {
		programState.executeInstructionWithChannels(computer_id, in, out, out_final)
	}

	if out_final != nil {
		*out_final <- programState.last_output
	}
}

func (programState *ProgramState) executeInstructionWithChannels(computer_id int, in <-chan int, out chan<- int, out_final *chan int) {
	instruction := programState.memory[programState.instructionPointer]
	opcode := instruction % 100

	inOutArgs := programState.extractInOutArgs()

	switch opcode {
	case 1:
		programState.add(&inOutArgs)
	case 2:
		programState.multiply(&inOutArgs)
	case 3:
		programState.inputFromChannel(&inOutArgs, in)
	case 4:
		programState.outputToChannel(&inOutArgs, out)
	case 5:
		programState.jumpIfTrue(&inOutArgs)
	case 6:
		programState.jumpIfFalse(&inOutArgs)
	case 7:
		programState.lessThan(&inOutArgs)
	case 8:
		programState.equals(&inOutArgs)
	case 99:
		return
	default:
		panic("Invalid opcode")
	}
}

func (programState *ProgramState) executeInstruction() {
	instruction := programState.memory[programState.instructionPointer]
	opcode := instruction % 100

	inOutArgs := programState.extractInOutArgs()

	switch opcode {
	case 1:
		programState.add(&inOutArgs)
	case 2:
		programState.multiply(&inOutArgs)
	case 3:
		programState.input(&inOutArgs)
	case 4:
		programState.output(&inOutArgs)
	case 5:
		programState.jumpIfTrue(&inOutArgs)
	case 6:
		programState.jumpIfFalse(&inOutArgs)
	case 7:
		programState.lessThan(&inOutArgs)
	case 8:
		programState.equals(&inOutArgs)
	case 99:
		return
	default:
		panic("Invalid opcode")
	}
}

func (programState *ProgramState) jumpIfTrue(inOutArgs *InOutArgs) {
	if inOutArgs.in[0] != 0 {
		programState.instructionPointer = inOutArgs.in[1]
	} else {
		programState.instructionPointer += 3
	}
}

func (programState *ProgramState) jumpIfFalse(inOutArgs *InOutArgs) {
	if inOutArgs.in[0] == 0 {
		programState.instructionPointer = inOutArgs.in[1]
	} else {
		programState.instructionPointer += 3
	}
}

func (programState *ProgramState) lessThan(inOutArgs *InOutArgs) {
	if inOutArgs.in[0] < inOutArgs.in[1] {
		programState.memory[inOutArgs.out] = 1
	} else {
		programState.memory[inOutArgs.out] = 0
	}
	programState.instructionPointer += 4
}

func (programState *ProgramState) equals(inOutArgs *InOutArgs) {
	if inOutArgs.in[0] == inOutArgs.in[1] {
		programState.memory[inOutArgs.out] = 1
	} else {
		programState.memory[inOutArgs.out] = 0
	}
	programState.instructionPointer += 4
}

func (programState *ProgramState) add(inOutArgs *InOutArgs) {
	operand1 := inOutArgs.in[0]
	operand2 := inOutArgs.in[1]
	result := operand1 + operand2
	programState.memory[inOutArgs.out] = result
	programState.instructionPointer += 4
}

func (programState *ProgramState) multiply(inOutArgs *InOutArgs) {
	operand1 := inOutArgs.in[0]
	operand2 := inOutArgs.in[1]
	result := operand1 * operand2
	programState.memory[inOutArgs.out] = result
	programState.instructionPointer += 4
}

func (programState *ProgramState) input(inOutArgs *InOutArgs) {
	fmt.Print("Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // reads until newline
	line := scanner.Text()
	input, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}

	programState.memory[inOutArgs.out] = input
	programState.instructionPointer += 2
}

func (programState *ProgramState) inputFromChannel(inOutArgs *InOutArgs, in <-chan int) {
	input, ok := <-in
	if !ok {
		programState.halted = true
		return
	}
	programState.memory[inOutArgs.out] = input
	programState.instructionPointer += 2
}

func (programState *ProgramState) outputToChannel(inOutArgs *InOutArgs, out chan<- int) {
	output := inOutArgs.in[0]
	programState.last_output = output
	out <- output
	programState.instructionPointer += 2
}

func (programState *ProgramState) output(inOutArgs *InOutArgs) {
	output := inOutArgs.in[0]
	println(output)
	programState.instructionPointer += 2
}

func (programState *ProgramState) isHalted() bool {
	return programState.halted || programState.memory[programState.instructionPointer] == 99
}
