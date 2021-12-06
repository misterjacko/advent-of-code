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
func MakeSlice(str string) []string {
	return strings.Split(str, ",")
}

func MakeInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

func LoadInput(fileName string) map[int]int {
	input := map[int]int{}

	file, err := os.ReadFile(fileName)
	check(err)
	dataSlice := MakeSlice(string(file))
	for _, num := range dataSlice {
		if _, ok := input[MakeInt(num)]; ok {
			input[MakeInt(num)] += 1
		} else {
			input[MakeInt(num)] = 1
		}
	}

	for i := 0; i <= 8; i++ {
		if _, ok := input[i]; !ok {
			input[i] = 0
		}
	}
	return input
}

func NextDay(data map[int]int) map[int]int {
	nextDay := map[int]int{}
	for i := 0; i < 9; i++ {
		if i == 8 {
			nextDay[i] = data[0]
		} else if i == 6 {
			nextDay[i] = data[0] + data[7]
		} else {
			nextDay[i] = data[i+1]
		}

	}
	return nextDay
}

func Part1(data map[int]int, days int) int {
	fishTotal := 0
	fishMap := map[int]int{}
	for k, v := range data {
		fishMap[k] = v
	}
	for i := 0; i < days; i++ {
		fishMap = NextDay(fishMap)
	}
	for _, v := range fishMap {
		fishTotal += v
	}

	return fishTotal
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data, 80))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part1(data, 256))
	fmt.Println(time.Since(startTime))
}
