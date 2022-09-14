package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../data/day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataStr := string(data)
	dataLines := strings.Split(dataStr, "\n")
	fmt.Printf("length = %d\n", len(dataLines))
	values := make([]int, len(dataLines))
	for i := 0; i < len(dataLines); i++ {
		values[i], _ = strconv.Atoi(dataLines[i])
	}

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == 2020 {
				fmt.Printf("Found %d, %d : sum %d product %d\n", values[i], values[j], values[i]+values[j], values[i]*values[j])
			}
		}
	}

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			vi := values[i]
			vj := values[j]
			if vi+vj < 2020 {
				for k := j + 1; k < len(values); k++ {
					if vi+vj+values[k] == 2020 {
						fmt.Printf("found %d, %d, %d : %d %d\n", vi, vj, values[k], vi+vj+values[k], vi*vj*values[k])
					}
				}
			}
		}
	}
	fmt.Println("Done.")

}
