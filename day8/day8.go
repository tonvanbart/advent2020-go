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

}

func contains(location int, seen map[int]bool) bool {
	_, present := seen[location]
	return present
}
