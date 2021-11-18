package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func StringToInt(str string) int {
	iVal, err := strconv.Atoi(str)
	check(err)
	return iVal

}

func LookAndSay(input []int) []int {
	output := []int{}
	var currentChar int
	var currentCount int
	for i := 0; i < len(input); i++ {
		currentChar = input[i]
		currentCount = 1
		for j := i + 1; j < len(input); j++ {
			if currentChar == input[j] {
				currentCount += 1
			} else {
				break
			}
		}
		output = append(output, currentCount)
		output = append(output, currentChar)
		i += currentCount - 1
	}
	return output
}

func Part1(data string, iterations int) int {
	arrData := []int{}
	for _, char := range data {
		arrData = append(arrData, StringToInt(string(char)))
	}
	for i := 0; i < iterations; i++ {
		arrData = LookAndSay(arrData)
	}
	return len(arrData)
}

func main() {
	startTime := time.Now()
	data := "1113122113"
	fmt.Printf("Part 1: %v\n", Part1(data, 40))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part1(data, 50))
	fmt.Println(time.Since(startTime))
}
