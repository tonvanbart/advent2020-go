package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	program := strings.Split(string(input), "\n")
	fmt.Printf("%d instructions in program\n", len(program))

	acc := 0
	pc := 0
	seen := make(map[int]bool)
	for !contains(pc, seen) {
		seen[pc] = true
		parts := strings.Split(program[pc], " ")
		opcode := parts[0]
		value, _ := strconv.Atoi(parts[1])
		if opcode == "nop" {
			pc = pc + 1
		} else if opcode == "acc" {
			acc = acc + value
			pc = pc + 1
		} else if opcode == "jmp" {
			pc = pc + value
		}
	}
	fmt.Printf("loop at line %d with acc=%d \n", pc+1, acc)

	// part 2: switch jmp for nop, one occurrence at a time
	for i := range program {
		// shallow copy to not modify the original
		var copy []string
		copy = append(copy, program...)
		if copy[i][0:3] == "acc" {
			// not changing acc, nothing to do
			continue
		} else if copy[i][0:3] == "jmp" {
			copy[i] = strings.Replace(copy[i], "jmp", "nop", -1)
		} else {
			copy[i] = strings.Replace(copy[i], "nop", "jmp", -1)
		}
		looped, result := runProgram(copy)
		if !looped {
			fmt.Printf("program ended normally, acc=%d\n", result)
			break
		}
	}
}

func contains(location int, seen map[int]bool) bool {
	_, present := seen[location]
	return present
}

// run the given program and return a bool indicating if it looped, and the final value of the accumulator.
func runProgram(program []string) (bool, int) {
	acc := 0
	pc := 0
	seen := make(map[int]bool)
	for !contains(pc, seen) && pc < len(program) {
		seen[pc] = true
		parts := strings.Split(program[pc], " ")
		opcode := parts[0]
		value, _ := strconv.Atoi(parts[1])
		if opcode == "nop" {
			pc = pc + 1
		} else if opcode == "acc" {
			acc = acc + value
			pc = pc + 1
		} else if opcode == "jmp" {
			pc = pc + value
		}
	}
	if pc == len(program) {
		return false, acc
	} else {
		return true, acc
	}
}
