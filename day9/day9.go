package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	length := 25 // 5 for example, 25 for input
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	nums := make([]int, len(lines))
	for i, nr := range lines {
		nums[i], err = strconv.Atoi(nr)
		if err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Println(nums)
	for pos := length; pos < len(nums); pos++ {
		previous := nums[pos-length : pos]
		if !isSum(nums[pos], previous) {
			fmt.Printf("%d not a sum of any previous %d\n", nums[pos], length)
			break
		}
	}
}

// return true if nr is a sum of any two numbers in nums
func isSum(nr int, nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nr == nums[i]+nums[j] {
				return true
			}
		}
	}
	return false
}
