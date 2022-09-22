package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	groups := strings.Split(string(input), "\n\n")
	fmt.Printf("There are %d groups in the input\n", len(groups))

	counts := make([]int, 500)
	for i, group := range groups {
		uniques := make(map[rune]bool)
		letters := strings.ReplaceAll(group, "\n", "")
		for _, answer := range []rune(letters) {
			uniques[answer] = true
		}
		counts[i] = len(uniques)
	}
	total := 0
	for _, count := range counts {
		total += count
	}
	fmt.Printf("array starts %v and total is %d\n", counts[0:10], total)

}
