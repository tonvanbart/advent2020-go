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

	// part 2
	for lower := 0; lower < len(nums); lower++ {
		for upper := lower + 1; upper <= len(nums); upper++ {
			sum, min, max := addall(nums[lower:upper])
			if sum == 22406676 {
				fmt.Println(nums[lower:upper])
				fmt.Printf("%d + %d = %d\n", min, max, min+max)
			}
			if sum > 22406676 {
				break
			}
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

func addall(nums []int) (int, int, int) {
	sum := 0
	min := nums[0]
	max := nums[0]
	for _, nr := range nums {
		sum += nr
		if nr < min {
			min = nr
		}
		if nr > max {
			max = nr
		}
	}
	return sum, min, max
}
