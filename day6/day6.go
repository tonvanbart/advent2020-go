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
	sharedAnswerCount := 0
	// commons := make([]int, 500)
	for i, group := range groups {
		uniques := make(map[rune]bool)
		letters := strings.ReplaceAll(group, "\n", "")
		for _, answer := range []rune(letters) {
			uniques[answer] = true
		}
		counts[i] = len(uniques)

		// part 2
		persons := strings.Split(group, "\n") // an array of answers per person
		sharedAnswers := makemap([]rune(persons[0]))
		for _, person := range persons[1:] {
			personAnswers := makemap([]rune(person))
			sharedAnswers = intersect(sharedAnswers, personAnswers)
		}
		fmt.Printf("shared answers for group: %v\n", sharedAnswers)
		sharedAnswerCount += len(sharedAnswers)
	}
	total := 0
	for _, count := range counts {
		total += count
	}
	fmt.Printf("array starts %v and total is %d\n", counts[0:10], total)
	fmt.Printf("the total count of all shared answers=%d\n", sharedAnswerCount)

}

// return a map that is the intersection of the two argument maps.
func intersect(a map[rune]bool, b map[rune]bool) map[rune]bool {
	result := make(map[rune]bool)
	for key, _ := range a {
		_, exist := b[key]
		if exist {
			result[key] = true
		}
	}
	return result
}

func makemap(letters []rune) map[rune]bool {
	result := make(map[rune]bool)
	for _, letter := range letters {
		result[letter] = true
	}
	return result
}
