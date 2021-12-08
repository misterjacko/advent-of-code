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
func MakeSlice(str string) []string {
	return strings.Split(str, " ")
}
func StringToInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

func LoadInput(fileName string) [][][]string {
	input := [][][]string{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		add := [][]string{}
		block := scanner.Text()
		blockA := strings.Split(block, " | ")
		for _, str := range blockA {
			add = append(add, MakeSlice(str))
		}
		input = append(input, add)

	}
	return input
}

func Part1(data [][][]string) int {
	uniqueCount := 0
	for _, line := range data {
		for _, digit := range line[1] {
			if len(digit) <= 4 {
				uniqueCount++
			} else if len(digit) == 7 {
				uniqueCount++
			}
		}
	}
	return uniqueCount
}

func SortSignalPattern(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		return len(data[i]) < len(data[j])
	})
	return data
}

func FourMinusOne(four string, one string) string {
	for _, char := range one {
		four = strings.Replace(four, string(char), "", -1)
	}
	return four
}
func HasOne(pattern, one string) bool {
	for _, char := range one {
		if !strings.Contains(pattern, string(char)) {
			return false
		}
	}
	return true
}

func HasFourMinusOne(pattern, fourMinusOne string) bool {
	for _, char := range fourMinusOne {
		if !strings.Contains(pattern, string(char)) {
			return false
		}
	}
	return true
}
func SortKeys(key string) string {
	newKey := strings.Split(key, "")
	sort.Strings(newKey)
	return strings.Join(newKey, "")
}

func Calculate(signalMap map[string]string, codes []string) int {
	total := []string{}
	for _, code := range codes {
		code = SortKeys(code)
		total = append(total, signalMap[code])
	}

	return StringToInt(strings.Join(total, ""))
}

func Part2(data [][][]string) int {
	output := 0
	for _, line := range data {
		signalMap := map[string]string{}
		one := ""
		four := ""
		for _, signalPattern := range SortSignalPattern(line[0]) {
			signalPattern = SortKeys(signalPattern)
			length := len(signalPattern)
			if length == 2 {
				signalMap[signalPattern] = "1"
				one = signalPattern
				continue
			}
			if length == 3 {
				signalMap[signalPattern] = "7"
				continue
			}
			if length == 4 {
				signalMap[signalPattern] = "4"
				four = signalPattern
				continue
			}
			if length == 7 {
				signalMap[signalPattern] = "8"
				continue
			}
			if length == 5 {
				if HasOne(signalPattern, one) {
					signalMap[signalPattern] = "3"
					continue
				}
				fourMinusOne := FourMinusOne(four, one)
				if HasFourMinusOne(signalPattern, fourMinusOne) {
					signalMap[signalPattern] = "5"
					continue
				} else {
					signalMap[signalPattern] = "2"
					continue
				}
			}
			if length == 6 {
				if !HasOne(signalPattern, one) {
					signalMap[signalPattern] = "6"
					continue
				}
				fourMinusOne := FourMinusOne(four, one)
				if HasFourMinusOne(signalPattern, fourMinusOne) {
					signalMap[signalPattern] = "9"
					continue
				} else {
					signalMap[signalPattern] = "0"
					continue
				}
			}
		}
		output += Calculate(signalMap, line[1])

	}
	return output
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))
}
