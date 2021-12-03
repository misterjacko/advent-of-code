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
	return strings.Split(str, "")
}

func InitSlice(str string) []int {
	onesTotal := []int{}
	for i := 0; i < len(str); i++ {
		onesTotal = append(onesTotal, 0)
	}
	return onesTotal
}

func ConvertBinary(str string) int {
	intVal, err := strconv.ParseInt(str, 2, 64)
	check(err)
	return int(intVal)
}

func Part1(data []string) int {
	onesTotal := InitSlice(data[0])
	half := len(data) / 2
	for _, binaryNumber := range data {
		for index, number := range MakeSlice(binaryNumber) {
			if number == "1" {
				onesTotal[index] += 1
			}
		}
	}
	gamma := ""
	epsilon := ""
	for _, total := range onesTotal {
		if total > half {
			gamma = fmt.Sprintf("%s1", gamma)
			epsilon = fmt.Sprintf("%s0", epsilon)
		} else {
			gamma = fmt.Sprintf("%s0", gamma)
			epsilon = fmt.Sprintf("%s1", epsilon)
		}
	}
	return ConvertBinary(gamma) * ConvertBinary(epsilon)
}

func MostCommon(data []string, index int) string {
	oneCount := float32(0)
	length := float32(len(data)) / float32(2)
	for _, binary := range data {
		if string(binary[index]) == "1" {
			oneCount += 1
		}
	}
	if oneCount >= length {
		return "1"
	} else {
		return "0"
	}
}

func LeastCommon(data []string, index int) string {
	oneCount := float32(0)
	length := float32(len(data)) / float32(2)
	for _, binary := range data {
		if string(binary[index]) == "1" {
			oneCount += 1
		}
	}
	if oneCount >= length {
		return "0"
	} else {
		return "1"
	}
}

func KeepOnly(data []string, keepValue string, index int) []string {
	newData := []string{}
	for _, line := range data {
		if string(line[index]) == keepValue {
			newData = append(newData, line)
		}
	}
	return newData
}

func DoTheWork(data []string, even string) string {
	binaryLength := len(data[0])
	for i := 0; i < binaryLength; i++ {
		lookFor := ""
		if even == "1" {
			lookFor = MostCommon(data, i)
		} else if even == "0" {
			lookFor = LeastCommon(data, i)
		}

		data = KeepOnly(data, lookFor, i)
		if len(data) == 1 {
			return data[0]
		}
	}
	return data[0]
}

func Part2(oxygen []string) int {
	co2 := make([]string, len(oxygen))
	copy(co2, oxygen)
	oxygenVal := DoTheWork(oxygen, "1")
	co2Val := DoTheWork(co2, "0")

	// find Oxygen
	return ConvertBinary(oxygenVal) * ConvertBinary(co2Val)
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))
}
