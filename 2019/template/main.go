package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LoadInput(fileName string) []int {
	input := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		input = append(input, line)
	}
	return input
}

func Part1(i []int) bool {
	return true
}

func Part2(i []int) bool {
	return true
}

func main() {
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Printf("Part 1: %v\n", Part2(data))

}
