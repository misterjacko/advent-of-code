package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
func LoadInput(fileName string) []int {
	input := []int{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		block := scanner.Text()
		input = append(input, MakeInt(block))

	}
	return input
}
func IsGreater(val1 int, val2 int) bool {
	return val1 < val2
}
func Part1(data []int) int {
	counter := 0
	for i := 1; i < len(data); i++ {
		if IsGreater(data[i-1], data[i]) {
			counter++
		}
	}
	return counter
}

func Combine3(data []int) []int {
	threeMW := []int{}
	for i := 0; i < len(data)-2; i++ {
		three := data[i] + data[i+1] + data[i+2]
		threeMW = append(threeMW, three)
	}
	return threeMW
}

func Part2(data []int) int {
	newData := Combine3(data)
	return Part1(newData)
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}
