package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	maxid := 0
	maxrow := 0
	maxcol := 0
	maxline := ""
	for _, line := range lines {
		row, col, id := convert(line)
		if id > maxid {
			maxline = line
			maxid = id
			maxrow = row
			maxcol = col
		}
	}
	fmt.Printf("Maximum seatid is %d at row %d and col %d for input %s\n", maxid, maxrow, maxcol, maxline)
}

// take a boarding pass string and return row, column and seat id
func convert(boarding string) (int, int, int) {
	v1 := strings.ReplaceAll(boarding, "F", "0")
	v2 := strings.ReplaceAll(v1, "B", "1")
	v3 := strings.ReplaceAll(v2, "R", "1")
	v4 := strings.ReplaceAll(v3, "L", "0")
	seatid, err := strconv.ParseInt(v4, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	rownum := seatid >> 3 // highest bits were the rownum
	column := seatid & 7  // bottom 3 bits the column
	return int(rownum), int(column), int(seatid)
}
