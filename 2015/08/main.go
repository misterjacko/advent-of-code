package main

import (
	"bufio"
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
func CodeLength(str string) int {
	return len(str)
}

func Escape(char string) bool {
	return char == "\\"
}
func SkipNumber(nextChar string) int {
	if nextChar == "x" {
		return 3
	}
	return 1
}

func StringLength(str string) int {
	remove := 2
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if Escape(char) {
			next := string(str[i+1])
			jump := SkipNumber(next)
			remove += jump
			i += jump
		}
	}
	return len(str) - remove
}

func Part1(data []string) int {
	totalCode := 0
	totalInside := 0
	for _, line := range data {
		totalCode += CodeLength(line)
		totalInside += StringLength(line)
	}
	return (totalCode - totalInside)
}

func NeedsEscape(char string) bool {
	needsEscape := "\"\\"
	return strings.Contains(needsEscape, char)
}

func StringLengthAdded(str string) int {
	add := 2
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if NeedsEscape(char) {
			add += 1
		}
	}
	return len(str) + add
}

func Part2(data []string) int {
	totalCode := 0
	totalWithEscapes := 0
	for _, line := range data {
		totalCode += CodeLength(line)
		totalWithEscapes += StringLengthAdded(line)
	}
	return (totalWithEscapes - totalCode)
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}
