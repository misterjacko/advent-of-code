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

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	// check(err)
	return i, err
}

func IntToBinary(i int64) string {
	return strconv.FormatInt(i, 2)
}

func AND(num1 int, num2 int) int {
	return (num1 & num2)
}

func OR(num1 int, num2 int) int {
	return (num1 | num2)
}
func LSHIFT(num int, shift int) int {
	return (num << shift)
}
func RSHIFT(num int, shift int) int {
	return (num >> shift)
}

func NOT(num int) int {
	val := 65535
	return (val ^ num)
}

func Part1(data []string, wire string) int {
	wireMap := map[string]int{}
	// pendingCommands := []string{}
	count := 0
	for {
		count += 1
		// fmt.Println(wireMap)
		fmt.Println(count)
		if wireValue, ok := wireMap[wire]; ok {
			return wireValue
		} else {
			for index, line := range data {

				if strings.Contains(line, "skip") {
					continue
				}
				arrLine := Array(line)
				if len(arrLine) == 3 {
					intValue, err := StringToInt(arrLine[0])
					if err != nil {
						if value, ok := wireMap[arrLine[0]]; ok {
							wireMap[arrLine[2]] = value
							data[index] = "skip"
						} else {
							continue
						}
					} else {
						wireMap[arrLine[2]] = intValue
						data[index] = "skip"
					}
				}
				if len(arrLine) == 4 {
					if value, ok := wireMap[arrLine[1]]; ok {
						wireMap[arrLine[3]] = NOT(value)
						data[index] = "skip"
					}

				}
				if len(arrLine) == 5 {
					if (arrLine[1] == "AND") || (arrLine[1] == "OR") {
						if value1, ok1 := wireMap[arrLine[0]]; ok1 {
							if value2, ok2 := wireMap[arrLine[2]]; ok2 {
								if arrLine[1] == "AND" {
									wireMap[arrLine[4]] = AND(value1, value2)
								}
								if arrLine[1] == "OR" {
									wireMap[arrLine[4]] = OR(value1, value2)
								}
								data[index] = "skip"
							} else {
								continue
							}
						} else {
							continue
						}
					} else {
						if value, ok := wireMap[arrLine[0]]; ok {
							if arrLine[1] == "LSHIFT" {
								intValue, err := StringToInt(arrLine[2])
								check(err)
								wireMap[arrLine[4]] = LSHIFT(value, intValue)
							}
							if arrLine[1] == "RSHIFT" {
								intValue, err := StringToInt(arrLine[2])
								check(err)
								wireMap[arrLine[4]] = RSHIFT(value, intValue)
							}
							data[index] = "skip"
						}
					}
				}
			}
		}
		fmt.Println(data)
		os.Exit(1)
	}
	// run everything that has 3 length and remove those instructions from a pending instruction list
	// iterate over pending instructions. IF one or both of the wires on the left are in the map execute and remove.
	// recursively do this. or go until we get a value for A
}

func Part2(data []string) int {

	return 1
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data, "a"))
	fmt.Println(time.Since(startTime))
	// fmt.Printf("Part 2: %v\n", Part2(data))
	// fmt.Println(time.Since(startTime))

}
