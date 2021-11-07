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

func Part1(i []string) bool {
	return true
}

func Part2(i []string) bool {
	return true
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	// fmt.Printf("Part 1: %v\n", Part2(data))
	// fmt.Println(time.Since(startTime))

}
