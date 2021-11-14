package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
func NoForbiddenSubString(str string) bool {
	response := true
	forbiddenStrings := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, fs := range forbiddenStrings {
		if strings.Contains(str, fs) {
			return false
		}
	}
	return response
}

func IsVowel(char string) int {
	vowels := "aeiou"
	if strings.Contains(vowels, char) {
		return 1
	} else {
		return 0
	}
}

func IsDupe(pastChar string, charStr string) bool {
	return pastChar == charStr
}

func Iterate(str string) bool {
	vowelCount := 0
	hasDouble := false
	pastChar := ""
	for _, char := range str {
		charStr := string(char)
		if vowelCount < 3 {
			vowelCount += IsVowel(charStr)
		}
		if !hasDouble {
			hasDouble = IsDupe(pastChar, charStr)
		}
		pastChar = charStr

	}
	if vowelCount >= 3 && hasDouble {
		return true
	}
	return false
}

func Part1(data []string) int {
	niceCount := 0
	for _, str := range data {
		if Iterate(str) && NoForbiddenSubString(str) {
			niceCount += 1
		}
	}
	return niceCount
}
func Array(str string) []string {
	returnArray := []string{}
	for _, char := range str {
		returnArray = append(returnArray, string(char))
	}
	return returnArray
}

func Sandwich(input string) bool {
	arrInput := Array(input)
	for i := 0; i < len(arrInput)-2; i++ {
		if arrInput[i] == arrInput[i+2] {
			return true
		}
	}
	return false
}

func TwoRepeat(input string) bool {
	arrInput := Array(input)
	charMap := map[string]int{}
	for i := 0; i < len(arrInput)-1; i++ {
		key := fmt.Sprintf("%s%s", arrInput[i], arrInput[i+1])
		if index, ok := charMap[key]; ok {
			if i-index > 1 {
				return true
			}
		} else {
			charMap[key] = i
		}
	}
	return false
}

func Part2(data []string) int {
	niceCount := 0
	for _, str := range data {
		if Sandwich(str) && TwoRepeat(str) {
			niceCount += 1
		}
	}
	return niceCount
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}
