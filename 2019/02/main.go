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

func LoadInput(fileName string) []int {
	input := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = strings.Split(scanner.Text(), ",")
	}
	intInput := []int{}
	for _, val := range input {
		intInput = append(intInput, MakeInt(val))
	}
	return intInput
}

func Add(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) int {
	return a * b
}

func MakeInt(str string) int {
	intVal, _ := strconv.Atoi(str)
	return intVal
}

func Execute(cleanData []int, noun int, verb int) int {
	data := make([]int, len(cleanData))
	copy(data, cleanData)
	index := 0
	data[1] = noun
	data[2] = verb

	for {
		operator := data[index]
		if operator == 99 {
			return data[0]
		}
		if operator == 1 {
			sum := Add(data[data[index+1]], data[data[index+2]])
			data[data[index+3]] = sum
		}
		if operator == 2 {
			prod := Multiply(data[data[index+1]], data[data[index+2]])
			data[data[index+3]] = prod
		}
		index += 4
	}
}
func Part1(cleanData []int) int {
	noun := 12
	verb := 2
	return Execute(cleanData, noun, verb)
}
func Part2(cleanData []int) int {
	var noun, verb int
	for noun = 0; noun < 99; noun++ {
		for verb = 0; verb < 99; verb++ {
			output := Execute(cleanData, noun, verb)
			if output == 19690720 {
				return (100*noun + verb)
			}
		}
	}
	return 0
}

func main() {
	startTime := time.Now()
	cleanData := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(cleanData))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(cleanData))
	fmt.Println(time.Since(startTime))
}
