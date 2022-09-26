package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	rulesin := strings.Split(string(input), "\n")
	fmt.Printf("nr of rules %d\n", len(rulesin))

	rules := make(map[string]map[string]int)
	re, err := regexp.Compile(`(\d*)( \w+ \w+ bag)`)
	if err != nil {
		log.Fatal(err)
	}
	keyexp, err := regexp.Compile(`^(\w+ \w+ bag)`)
	for _, line := range rulesin {
		key := keyexp.FindString(line)
		parts := re.FindAllString(line, -1)
		// fmt.Println(key, parts, len(parts))
		rules[key] = makeset(parts)
	}

	// printrules(rules)

	// how many ways to get to a shiny gold bag?
	count := 0
	for bag := range rules {
		if cancontain(bag, "shiny gold bag", rules) {
			count++
		}
	}
	fmt.Printf("%d bags can contain a shiny gold bag\n", count)

	// how many bags inside a shiny gold bag?
	bagcontents := countContents("shiny gold bag", rules, 0)
	fmt.Printf("In total there are %d bags inside a shiny gold bag\n", bagcontents)

}

// return a map of string to bool, containing all input strings as keys and the number as values.
func makeset(bags []string) map[string]int {
	result := make(map[string]int)
	for _, bag := range bags {
		if bag == "no other bag" {
			result[bag] = 0
		} else {
			number, _ := strconv.Atoi(string(bag[0]))
			key := bag[2:]
			result[key] = number
		}
	}
	return result

}

// recursive search if bag x eventually can contain bag y under the given rules
// return true if there is a way
func cancontain(x string, y string, rules map[string]map[string]int) bool {
	rulex := rules[x]
	if len(rulex) == 1 && contains(rulex, "no other bag") {
		return false
	}
	if contains(rulex, y) {
		return true
	}
	// loop over the bags that x contains and recurse
	for bag := range rulex {
		if cancontain(bag, y, rules) {
			return true
		}
	}
	// not found anywhere
	return false
}

// recursive count of number of bags contained in "bag"
func countContents(bag string, rules map[string]map[string]int, acc int) int {
	// log.Printf("countContents(%s, rules, %d)\n", bag, acc)
	bagrule := rules[bag]
	if len(bagrule) == 1 && contains(bagrule, "no other bag") {
		return acc
	}
	for inner, count := range bagrule {
		within := countContents(inner, rules, 0)
		acc = acc + count + (count * within)
	}
	return acc
}

// return true if the given map contains the given key
func contains(rule map[string]int, key string) bool {
	_, exists := rule[key]
	return exists
}

// display the rules
func printrules(rules map[string]map[string]int) {
	for bag, containing := range rules {
		fmt.Printf("bag %s contains: ", bag)
		for description, number := range containing {
			fmt.Printf("%d %s ", number, description)
		}
		fmt.Println()
	}
}
