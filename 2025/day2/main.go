package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isInvalid(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)
	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}
		pattern := s[:l]
		repeated := strings.Repeat(pattern, n/l)
		if repeated == s {
			return true
		}
	}
	return false
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("Error opening file: %e", err)
	}

	ranges := strings.Split(string(input), ",")
	var total int64 = 0

	for _, r := range ranges {
		bounds := strings.Split(strings.TrimSpace(r), "-")
		if len(bounds) != 2 {
			continue
		}

		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])

		for i := start; i <= end; i++ {
			if isInvalid(i) {
				total += int64(i)
			}
		}
	}

	fmt.Println(total)
}
