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
	// split input into passports (separated by blank line)
	passports := strings.Split(string(input), "\n\n")
	fmt.Printf("nr of input passports: %d \n", len(passports))
	validcount := 0

	for i := 0; i < len(passports); i++ {
		// split passport in key:value pairs
		passport := strings.Fields(passports[i])
		// log.Printf("%d fields in passport\n", len(passport))

		// put in map
		fields := make(map[string]string)
		for j := 0; j < len(passport); j++ {
			kv := strings.Split(passport[j], ":")
			fields[kv[0]] = kv[1]
		}
		// check required key presence
		_, byr := fields["byr"]
		_, iyr := fields["iyr"]
		_, eyr := fields["eyr"]
		_, hgt := fields["hgt"]
		_, hcl := fields["hcl"]
		_, ecl := fields["ecl"]
		_, pid := fields["pid"]

		// count if valid
		if byr && iyr && eyr && hgt && hcl && ecl && pid {
			validcount++
		}
	}
	fmt.Printf("Valid passports: %d\n", validcount)
}
