package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	test := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	// data, err := os.ReadFile("input")
	// if err != nil {
	// 	log.Fatalf("Error opening file: %e", err)
	// }
	log.Print(part1(string(test)))
}

func part1(data string) int {
	invalidIdsCount := 0
	ranges := strings.Split(data, ",")
	for _, el := range ranges {
		interval := strings.Split(el, "-")
		start, err := strconv.Atoi(interval[0])
		if err != nil {
			log.Fatal("cabum")
		}
		end, err := strconv.Atoi(interval[1])
		if err != nil {
			log.Fatal("cabum")
		}
		for {
			startDigits := digitsUsingLog(start)
			if startDigits%2 != 0 {
				newStart := "1"
				for {
					newStart = newStart + "0"
					if len(newStart) == startDigits+1 {
						start, err = strconv.Atoi(newStart)
						if err != nil {
							log.Fatalf("New start cabum %e", err)
						}
						break
					}
				}
			}
			if start > end {
				break
			}
			id := strconv.Itoa(start)
			if id[:len(id)/2] == id[len(id)/2:] {
				invalidIdsCount = invalidIdsCount + start
			}
			start = start + 1
		}
	}
	return invalidIdsCount
}

func digitsUsingLog(n int) int {
	NumberOfDigits := math.Floor(math.Log10(float64(n))) + 1
	return int(NumberOfDigits)
}
