package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func LoadInput(fileName string) []string {
	input := []string{}

	file, err := os.ReadFile(fileName)
	check(err)
	for _, data := range string(file) {
		input = append(input, string(data))
	}
	return input
}

func Part1(data []string) int {
	floor := 0
	for _, step := range data {
		if step == "(" {
			floor += 1
		}
		if step == ")" {
			floor -= 1
		}
	}
	return floor
}

func Part2(data []string) int {
	floor := 0
	for index, step := range data {
		if step == "(" {
			floor += 1
		}
		if step == ")" {
			floor -= 1
			if floor < 0 {
				return index + 1
			}
		}
	}
	return 0
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))
}
