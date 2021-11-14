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

type Instruction struct {
	Command  string
	RowStart int
	RowEnd   int
	ColStart int
	ColEnd   int
}

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
func Array(str string) []string {
	str = strings.Replace(str, "turn ", "", -1)
	str = strings.Replace(str, "through ", "", -1)
	return strings.Split(str, " ")
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func MakeInstruction(iLine []string) *Instruction {
	set1 := strings.Split(iLine[1], ",")
	set2 := strings.Split(iLine[2], ",")
	return &Instruction{
		Command:  iLine[0],
		RowStart: StringToInt(set1[0]),
		RowEnd:   StringToInt(set2[0]),
		ColStart: StringToInt(set1[1]),
		ColEnd:   StringToInt(set2[1]),
	}
}

func ChangeLight(lightState bool, instruction string) bool {
	switch instruction {
	case "on":
		return true
	case "off":
		return false
	case "toggle":
		return !lightState
	}
	return lightState
}

func MakeArray1000() [1000][1000]bool {
	myArray := [1000][1000]bool{}
	return myArray
}

func MakeNumberArray1000() [1000][1000]int {
	myArray := [1000][1000]int{}
	return myArray
}

func Part1(data []string) int {
	lightGrid := MakeArray1000()
	for _, line := range data {
		instruction := MakeInstruction(Array((line)))
		for row := instruction.RowStart; row <= instruction.RowEnd; row++ {
			for col := instruction.ColStart; col <= instruction.ColEnd; col++ {
				lightGrid[row][col] = ChangeLight(lightGrid[row][col], instruction.Command)
			}
		}
	}
	dict := map[bool]int{true: 0, false: 0}
	for row := range lightGrid {
		for col := range lightGrid[row] {
			dict[lightGrid[row][col]] += 1
		}
	}
	return dict[true]
}

func ChangeLightBrightness(lightState int, instruction string) int {
	switch instruction {
	case "on":
		lightState += 1
		return lightState
	case "off":
		if lightState <= 0 {
			return 0
		} else {
			lightState -= 1
			return lightState
		}
	case "toggle":
		lightState += 2
		return lightState
	}
	return lightState
}

func Part2(data []string) int {
	lightGrid := MakeNumberArray1000()
	for _, line := range data {
		instruction := MakeInstruction(Array((line)))
		for row := instruction.RowStart; row <= instruction.RowEnd; row++ {
			for col := instruction.ColStart; col <= instruction.ColEnd; col++ {
				lightGrid[row][col] = ChangeLightBrightness(lightGrid[row][col], instruction.Command)
			}
		}
	}
	sum := 0
	for row := range lightGrid {
		for col := range lightGrid[row] {
			sum += lightGrid[row][col]
		}
	}
	return sum
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}
