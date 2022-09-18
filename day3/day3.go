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
	inputLines := strings.Split(string(input), "\n")
	w := len(inputLines[0])
	fmt.Printf("nr of input lines: %d of length %d\n", len(inputLines), w)
	count := 0
	col := 0
	for i := 0; i < len(inputLines); i++ {
		log.Printf("check line %d col %d", i+1, col)
		if inputLines[i][col] == '#' {
			log.Printf("'#' at line %d col %d", i+1, col)
			count++
		}
		col += 3
		if col >= w {
			col = col - w
		}
	}
	fmt.Printf("Encountered %d trees\n", count)

	count11 := check(1, 1, inputLines)
	count13 := check(1, 3, inputLines)
	count15 := check(1, 5, inputLines)
	count17 := check(1, 7, inputLines)
	count21 := check(2, 1, inputLines)
	log.Printf("Found %d, %d, %d, %d and %d\n", count11, count13, count15, count17, count21)
	log.Printf("result is %d\n", count11*count13*count15*count17*count21)

}

// check the route obtained by stepping hor horizontal, ver vertical.
// it will return the number of trees encountered on this route.
func check(vert int, hor int, inputLines []string) int {
	count := 0
	col := 0
	width := len(inputLines[0])
	for i := 0; i < len(inputLines); i = i + vert {
		if inputLines[i][col] == '#' {
			count++
		}
		col += hor
		if col >= width {
			col = col - width
		}
	}
	return count
}
