package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	min    int
	max    int
	letter string
}

func matches(passw string, rule Rule) bool {
	count := strings.Count(passw, rule.letter)
	return count >= rule.min && count <= rule.max
}

func matches2(passw string, rule Rule) bool {
	runes := []rune(passw)
	cletter := []rune(rule.letter)[0]
	ix1 := rule.min - 1
	ix2 := rule.max - 1
	return xor(runes[ix1] == cletter, runes[ix2] == cletter)
}

func xor(a bool, b bool) bool {
	return (a || b) && !(a && b)
}

func makeRule(definition string) *Rule {
	parts := strings.Split(definition, " ")
	bounds := strings.Split(parts[0], "-")
	r := new(Rule)
	r.min, _ = strconv.Atoi(bounds[0])
	r.max, _ = strconv.Atoi(bounds[1])
	r.letter = parts[1]
	return r
}

func main() {
	input, err := os.ReadFile("../data/day2-input.txt")
	if err != nil {
		log.Fatal("error reading file", err)
	}
	inputLines := strings.Split(string(input), "\n")
	fmt.Println(len(inputLines))

	valid := 0
	valid2 := 0
	for i := 0; i < len(inputLines); i++ {
		parts := strings.Split(inputLines[i], ": ")
		rule := makeRule(parts[0])
		if matches(parts[1], *rule) {
			// fmt.Printf("'%s' has %s between %d and %d times\n", parts[1], rule.letter, rule.min, rule.max)
			valid++
		}
		if matches2(parts[1], *rule) {
			fmt.Printf("'%s' has %s at pos %d or %d\n", parts[1], rule.letter, rule.min, rule.max)
			valid2++
		}

	}

	fmt.Printf("There are %d valid passwords\n", valid)
	fmt.Printf("There are %d valid passwords under rule 2\n", valid2)
}
