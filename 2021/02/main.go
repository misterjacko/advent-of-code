package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}
func MakeInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

// multiple lines
func LoadInput(fileName string) []string {
	input := []string{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		block := scanner.Text()
		input = append(input, block)

	}
	return input
}

func MakeSlice(str string) []string {
	return strings.Split(str, " ")
}

func Part1(data []string) int {
	hPosition := 0
	depth := 0
	for _, instruction := range data {
		instructionSlice := MakeSlice(instruction)
		if instructionSlice[0] == "forward" {
			hPosition += MakeInt(instructionSlice[1])
		} else {
			if instructionSlice[0] == "down" {
				depth += MakeInt(instructionSlice[1])
			} else {
				depth -= MakeInt(instructionSlice[1])
			}
		}

	}
	return hPosition * depth
}

func Part2(data []string) int {
	hPosition := 0
	depth := 0
	aim := 0
	for _, instruction := range data {
		instructionSlice := MakeSlice(instruction)
		if instructionSlice[0] == "forward" {
			hPosition += MakeInt(instructionSlice[1])
			depth += aim * MakeInt(instructionSlice[1])

		} else {
			if instructionSlice[0] == "down" {
				aim += MakeInt(instructionSlice[1])
			} else {
				aim -= MakeInt(instructionSlice[1])
			}
		}

	}
	return hPosition * depth
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}
