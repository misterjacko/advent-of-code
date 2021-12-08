package main

import (
	"fmt"
	"log"
	"math"
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
		input[MakeInt(num)]++
	}

	return input
}

func OrderedSlice(data map[int]int) []int {
	orderedSlice := []int{}
	for key, _ := range data {
		orderedSlice = append(orderedSlice, key)
	}
	sort.Ints(orderedSlice)
	return orderedSlice
}

func Part1(data map[int]int) int {
	totalFuel := 1000000000000000000
	dataRange := OrderedSlice(data)
	for i := dataRange[0]; i <= dataRange[len(dataRange)-1]; i++ {
		thisFuel := 0
		for dataKey, dataValue := range data {
			absFuel := math.Abs(float64(dataKey - i))
			thisFuel += (dataValue * int(absFuel))
		}

		if thisFuel < totalFuel {
			totalFuel = thisFuel
		}
	}
	return totalFuel
}

func Part2(data map[int]int) int {
	totalFuel := 1000000000000000000
	dataRange := OrderedSlice(data)
	for i := dataRange[0]; i <= dataRange[len(dataRange)-1]; i++ {
		thisFuel := 0
		for dataKey, dataValue := range data {
			absFuel := math.Abs(float64(dataKey - i))
			eachFuel := (int(absFuel) * (int(absFuel) + 1)) / 2
			thisFuel += (dataValue * eachFuel)
		}

		if thisFuel < totalFuel {
			totalFuel = thisFuel
		}
	}
	return totalFuel
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))
}
