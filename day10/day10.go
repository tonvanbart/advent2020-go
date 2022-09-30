package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := "input"
	input, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	nums := make([]int, 1+len(lines))
	nums[0] = 0
	for i, nr := range lines {
		nums[i+1], err = strconv.Atoi(nr)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("input file '%s' length %d", file, len(nums))
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	fmt.Println(nums)
	ones := 0
	threes := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			ones++
		} else if nums[i]-nums[i-1] == 3 {
			threes++
		}
	}
	fmt.Printf("ones: %d, threes: %d, product %d\n", ones, threes, ones*threes)
}
