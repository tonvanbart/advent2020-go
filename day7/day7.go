package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.ReadFile("example")
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
	for bag, containing := range rules {
		fmt.Printf("bag %s contains: ", bag)
		for key, _ := range containing {
			fmt.Printf("%s ", key)
		}
		fmt.Println()
	}
}

// return a map of string to bool, containing all input strings as keys.
func makeset(bags []string) map[string]bool {
	result := make(map[string]bool)
	for _, bag := range bags {
		result[bag] = true
	}
	return result

}
