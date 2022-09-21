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
	// split input into passports (separated by blank line)
	passports := strings.Split(string(input), "\n\n")
	fmt.Printf("nr of input passports: %d \n", len(passports))
	validcount := 0
	validcount2 := 0

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

		// part 2
		if valid2(fields) {
			validcount2++
		}

	}
	fmt.Printf("Valid passports: %d\n", validcount)
	fmt.Printf("Out of %d passports %d were valid\n", len(passports), validcount2)

}

func valid2(fields map[string]string) bool {
	// valid eye colors
	colors := map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
	}
	// check required key presence
	vbyr, byr := fields["byr"]
	viyr, iyr := fields["iyr"]
	veyr, eyr := fields["eyr"]
	vhgt, hgt := fields["hgt"]
	vhcl, hcl := fields["hcl"]
	vecl, ecl := fields["ecl"]
	vpid, pid := fields["pid"]

	if !(byr && iyr && eyr && hgt && hcl && ecl && pid) {
		return false
	}

	birth, err := strconv.Atoi(vbyr)
	if err != nil || birth < 1920 || birth > 2002 {
		log.Printf("Birth %d did not match\n", birth)
		return false
	}

	issue, err := strconv.Atoi(viyr)
	if err != nil || issue < 2010 || issue > 2020 {
		log.Printf("iyr %d wrong\n", issue)
		return false
	}

	expire, err := strconv.Atoi(veyr)
	if err != nil || expire < 2020 || expire > 2030 {
		log.Printf("eyr %s wrong\n", veyr)
		return false
	}

	// check eye color present in valid colors
	_, found := colors[vecl]
	if !found {
		log.Printf("ecl %s wrong\n", vecl)
		return false
	}

	// check hair color on regex
	matched, err := regexp.MatchString("#[0-9a-f]{6}", vhcl)
	if err != nil || !matched {
		log.Printf("hcl %s wrong\n", vhcl)
		return false
	}

	// check vpid on regexp
	matchedpid, err := regexp.MatchString("\\d{9}", vpid)
	if err != nil || !matchedpid {
		log.Printf("pid %s wrong\n", vpid)
		return false
	}

	// match height on regex and keep subgroups for height and cm/in
	re, err := regexp.Compile(`(\d+)(cm|in)`)
	if err != nil {
		log.Fatal(err)
	}
	parts := re.FindStringSubmatch(vhgt)
	if len(parts) == 0 {
		log.Printf("height %s not parsed\n", vhgt)
		return false
	} else {
		height, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		if parts[2] == "cm" && (height < 150 || height > 193) {
			log.Printf("height %s out of bounds\n", vhgt)
			return false
		}
		if parts[2] == "in" && (height < 59 || height > 79) {
			log.Printf("height %s out of bounds\n", vhgt)
			return false
		}
	}

	return true

}
