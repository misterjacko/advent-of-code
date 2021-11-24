package main

import (
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

func LoadInput(fileName string) []string {
	input := []string{}

	file, err := os.ReadFile(fileName)
	check(err)
	for _, data := range string(file) {
		input = append(input, string(data))
	}
	return input
}
func StringToInt(str string) int {
	val, err := strconv.Atoi(str)
	check(err)
	return val
}

func MakeNeg(val int) int {
	return val * -1
}

func IsNegative(str string) bool {
	return str == "-"
}

func Part1(data []string) int {
	numberSlice := []int{}
	for i := 0; i < len(data); i++ {
		if _, err := strconv.Atoi(data[i]); err == nil {
			isNegative := IsNegative(data[i-1])
			numString := []string{data[i]}
			i++
			for {
				if _, err := strconv.Atoi(data[i]); err == nil {
					numString = append(numString, data[i])
					i++
				} else {
					break
				}
			}
			intVal := StringToInt(strings.Join(numString, ""))
			if isNegative {
				numberSlice = append(numberSlice, MakeNeg(intVal))
			} else {
				numberSlice = append(numberSlice, intVal)
			}
		}
	}
	total := 0
	for _, val := range numberSlice {
		total += val
	}
	return total
}

// func Part2(data []string) int {

// 	return 999
// }

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	// fmt.Printf("Part 2: %v\n", Part2(data))
	// fmt.Println(time.Since(startTime))
}
