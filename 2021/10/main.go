package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	data := strings.Split(string(file), "\n")

	for _, line := range data {
		input = append(input, line)
	}
	return input
}

func Score(char string) int {
	switch char {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	return 0
}

func Remove(str string) string {
	removeList := []string{"()", "[]", "{}", "<>"}

	for _, set := range removeList {
		str = strings.ReplaceAll(str, set, "")
	}
	return str
}

func RemoveAllPairs(str string) string {
	tempString := ""
	for {
		tempString = Remove(str)
		if len(tempString) == len(str) {
			return tempString
		} else {
			str = tempString
		}
	}
}

func Closer(char string) bool {
	switch char {
	case ")", "]", "}", ">":
		return true
	}
	return false
}
func CompletionScore(char string) int {
	switch char {
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	}
	return 0
}
func Part1(data []string) (int, int) {
	total := 0
	completeScores := []int{}

	for _, line := range data {
		corrupt := false
		line = RemoveAllPairs(line)
		for _, char := range line {
			if Closer(string(char)) {
				corrupt = true
				total += Score(string(char))
				break
			}
		}
		if !corrupt {
			// incomomplete
			completeScore := 0
			for i := len(line) - 1; i >= 0; i-- {
				completeScore = (completeScore * 5) + CompletionScore(string(line[i]))
			}
			completeScores = append(completeScores, completeScore)
		}
	}
	sort.Ints(completeScores)
	completeLength := len(completeScores)
	return total, completeScores[completeLength/2]
}
func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	partOneAnswer, partTwoAnswer := Part1(data)
	fmt.Printf("Part 1: %v\n", partOneAnswer)
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", partTwoAnswer)
	fmt.Println(time.Since(startTime))
}
