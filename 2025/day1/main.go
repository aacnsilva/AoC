package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const DIAL_START = 50
const DIAL_MAX = 99
const DIAL_MIN = 0

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("Error opening file: %e", err)
	}
	defer file.Close()

	dial := DIAL_START
	password := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch scanner.Text()[0] {
		case 'L':
			{
				rotations, err := strconv.Atoi(strings.Split(scanner.Text(), "L")[1])
				if err != nil {
					log.Fatal("cabum")
				}
				if rotations > DIAL_MAX {
					rotations = rotations % (DIAL_MAX + 1)
				}
				if (dial - rotations) < DIAL_MIN {
					dial = DIAL_MAX - rotations + dial + 1
				} else {
					dial = dial - rotations
				}
			}
		case 'R':
			{
				rotations, err := strconv.Atoi(strings.Split(scanner.Text(), "R")[1])
				if err != nil {
					log.Fatal("cabum")
				}
				if rotations > DIAL_MAX {
					rotations = rotations % (DIAL_MAX + 1)
				}
				if (dial + rotations) > DIAL_MAX {
					dial = dial + rotations - DIAL_MAX - 1
				} else {
					dial = dial + rotations
				}
			}
		}
		if dial == 0 {
			password = password + 1
		}
	}
	println(password)
}
