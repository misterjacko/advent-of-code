package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func LoadInput(fileName string) map[string][]string {
	input := map[string][]string{}
	file, err := os.ReadFile(fileName)
	check(err)
	data := strings.Split(string(file), "\n")
	for _, line := range data {
		lineSlice := strings.Split(line, "-")
		input[lineSlice[0]] = append(input[lineSlice[0]], lineSlice[1])
		input[lineSlice[1]] = append(input[lineSlice[1]], lineSlice[0])
	}
	return input
}

func IsUpper(chars string) bool {
	return unicode.IsUpper(rune(chars[0]))
}

func NotVisited(cave string, visitedLower []string) bool {
	for _, lower := range visitedLower {
		if cave == lower {
			return false
		}
	}
	return true
}

func Explore(thisCave string, visitedLower []string, goneOnce bool) {
	if thisCave == "end" {
		Count++
		return
	} else {

		if !IsUpper(thisCave) {
			if goneOnce {
				visitedLower = append(visitedLower, thisCave)
			} else {
				if thisCave == LowerCave {
					goneOnce = true
				} else {
					visitedLower = append(visitedLower, thisCave)
				}
			}
		}

		for _, cave := range Data[thisCave] {
			if NotVisited(cave, visitedLower) {
				Explore(cave, visitedLower, goneOnce)
			}

		}
	}
}

var Count int
var Data map[string][]string

func Part1(data map[string][]string) int {
	Count = 0
	Data = data
	goneOnce := true
	for _, cave := range Data["start"] {
		visitedLower := []string{"start"}
		Explore(cave, visitedLower, goneOnce)
	}
	return Count
}
func MakeLowerList(data map[string][]string) []string {
	lowerList := []string{}
	for key, _ := range data {
		if key != "start" && key != "end" && !IsUpper(key) {
			lowerList = append(lowerList, key)
		}
	}
	return lowerList
}

var LowerCave string

func Part2(data map[string][]string) int {
	finalCount := 0
	Data = data
	lowerList := MakeLowerList(data)
	mainPathCount := Part1(data)
	finalCount += mainPathCount
	for _, lowerCave := range lowerList {
		LowerCave = lowerCave
		Count = 0
		goneOnce := false
		for _, cave := range Data["start"] {
			visitedLower := []string{"start"}
			Explore(cave, visitedLower, goneOnce)
		}
		thisCount := Count - mainPathCount
		finalCount += thisCount

	}

	return finalCount
}

func main() {
	startTime := time.Now()
	data1 := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data1))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data1))
	fmt.Println(time.Since(startTime))
}
