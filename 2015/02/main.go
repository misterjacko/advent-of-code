package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func StringToInt(startVal []string) []int {
	returnVal := []int{}
	for _, val := range startVal {
		intVal, err := strconv.Atoi(val)
		check(err)
		returnVal = append(returnVal, intVal)
	}
	return returnVal
}

func LoadInput(fileName string) [][]int {
	input := [][]int{}
	block := []int{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		block = StringToInt(strings.Split(scanner.Text(), "x"))
		sort.Ints(block)
		input = append(input, block)

	}
	return input
}

func Part1(data [][]int) int {
	totalPaper := 0
	for _, line := range data {
		side0 := line[0] * line[1]
		side1 := line[1] * line[2]
		side2 := line[2] * line[0]
		linePaper := side0*3 + side1*2 + side2*2
		totalPaper += linePaper
	}
	return totalPaper
}

func Part2(data [][]int) int {
	totalRibbon := 0
	for _, line := range data {
		wrap := 2 * (line[0] + line[1])
		bow := line[0] * line[1] * line[2]
		lineRibbon := wrap + bow
		totalRibbon += lineRibbon
	}
	return totalRibbon
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}
