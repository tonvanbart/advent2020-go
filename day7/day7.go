package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	rulesin := strings.Split(string(input), "\n")
	fmt.Printf("nr of rules %d\n", len(rulesin))

	rules := make(map[string]map[string]bool)
	re, err := regexp.Compile(`(\w+ \w+ bag)`)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range rulesin {
		parts := re.FindAllString(line, -1)
		rules[parts[0]] = makeset(parts[1:])
	}

	// how many ways to get to a shiny gold bag?
	count := 0
	for bag, _ := range rules {
		if cancontain(bag, "shiny gold bag", rules) {
			count++
		}
	}
	fmt.Printf("%d bags can contain a shiny gold bag", count)

}

// return a map of string to bool, containing all input strings as keys.
func makeset(bags []string) map[string]bool {
	result := make(map[string]bool)
	for _, bag := range bags {
		result[bag] = true
	}
	return result

}

// recursive search if bag x eventually can contain bag y under the given rules
// return true if there is a way
func cancontain(x string, y string, rules map[string]map[string]bool) bool {
	rulex := rules[x]
	if len(rulex) == 1 && contains(rulex, "no other bag") {
		return false
	}
	if contains(rulex, y) {
		return true
	}
	// loop over the bags that x contains and recurse
	for bag, _ := range rulex {
		if cancontain(bag, y, rules) {
			return true
		}
	}
	// not found anywhere
	return false
}

// return true if the given map contains the given key
func contains(rule map[string]bool, key string) bool {
	_, exists := rule[key]
	return exists
}

// display the rules
func printrules(rules map[string]map[string]bool) {
	for bag, containing := range rules {
		fmt.Printf("bag %s contains: ", bag)
		for key, _ := range containing {
			fmt.Printf("%s ", key)
		}
		fmt.Println()
	}
}
