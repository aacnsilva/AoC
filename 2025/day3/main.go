package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	input, err := os.Open("input2")
	if err != nil {
		log.Fatalf("Error opening file: %e", err)
	}
	defer input.Close()

	totalJoltage := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		nums := make([]int, len(scanner.Text()))
		for _, n := range scanner.Text() {
			num, _ := strconv.Atoi(string(n))
			nums = append(nums, num)
		}
		// part 1
		// max1 := slices.Max(nums[:len(nums)-1])
		// max2 := slices.Max(nums[slices.Index(nums, max1)+1:])
		// totalJoltage = totalJoltage + max1*10 + max2

		// part 2
		max := slices.Max(nums[:len(nums)-11])
		maxJoltage := max + int(math.Pow(10, 12))
		fmt.Printf("max: %v\n", max)
		for i := 11; i > 0; i-- {
			max2 := slices.Max(nums[slices.Index(nums, max)+1 : len(nums)-i-1])
			maxJoltage = maxJoltage + max2 + int(math.Pow(10, float64(i+1)))
			max = max2
			fmt.Printf("max: %v\n", max2)
		}
		totalJoltage = totalJoltage + maxJoltage
	}
	fmt.Println(totalJoltage)
}
