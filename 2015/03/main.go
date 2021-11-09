package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

func makeKey(position [2]int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(position), " ", "x", -1), "[]")
}

func changePosition(direction string, position [2]int) [2]int {
	switch direction {
	case "^":
		position[0] += 1
	case "v":
		position[0] -= 1
	case ">":
		position[1] += 1
	case "<":
		position[1] -= 1
	}
	return position
}

func Part1(data []string) int {
	visitedHouses := map[string]int{}
	// [x,y]
	position := [2]int{0, 0}
	visitedHouses[makeKey(position)] = 1
	for _, direction := range data {
		position = changePosition(direction, position)
		key := makeKey(position)
		if value, ok := visitedHouses[key]; ok {
			visitedHouses[key] = value + 1
		} else {
			visitedHouses[key] = 1
		}
	}
	return len(visitedHouses)
}

func Part2(data []string) int {
	visitedHouses := map[string]int{}
	santaPosition := [2]int{0, 0}
	visitedHouses[makeKey(santaPosition)] = 1

	roboPosition := [2]int{0, 0}
	chooser := true
	for _, direction := range data {
		if chooser {
			santaPosition = changePosition(direction, santaPosition)
			key := makeKey(santaPosition)
			if value, ok := visitedHouses[key]; ok {
				visitedHouses[key] = value + 1
			} else {
				visitedHouses[key] = 1
			}
		} else {
			roboPosition = changePosition(direction, roboPosition)
			key := makeKey(roboPosition)
			if value, ok := visitedHouses[key]; ok {
				visitedHouses[key] = value + 1
			} else {
				visitedHouses[key] = 1
			}
		}
		chooser = !chooser
	}
	totalHouses := len(visitedHouses)
	return totalHouses
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}
