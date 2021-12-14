package main

import (
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

type PolyPlan struct {
	pair1, pair2 string
}

func LoadInput(fileName string) (map[string]int, map[string]PolyPlan, string, string) {
	templateMap := map[string]int{}
	polyPlans := map[string]PolyPlan{}
	file, err := os.ReadFile(fileName)
	check(err)
	data := strings.Split(string(file), "\n")
	template := strings.Split(data[0], "")
	beginning := template[0]
	end := template[len(template)-1]
	for i := 0; i < len(template)-1; i++ {
		templateMap[fmt.Sprintf("%s%s", template[i], template[i+1])]++
	}

	for i := 2; i < len(data); i++ {
		lineArr := strings.Split(data[i], " -> ")
		keySplit := strings.Split(lineArr[0], "")
		polyPlans[lineArr[0]] = PolyPlan{
			pair1: fmt.Sprintf("%s%s", keySplit[0], lineArr[1]),
			pair2: fmt.Sprintf("%s%s", lineArr[1], keySplit[1]),
		}
	}
	return templateMap, polyPlans, beginning, end
}

func CalcString(stringMap map[string]int, beginning, end string) map[string]int {
	resultMap := map[string]int{beginning: 1, end: 1}
	for key, value := range stringMap {
		keySlice := strings.Split(key, "")
		resultMap[keySlice[0]] += value
		resultMap[keySlice[1]] += value
	}
	for key, _ := range resultMap {
		resultMap[key] /= 2
	}
	return resultMap
}

func TakeTurn(templateMap map[string]int, polyPlans map[string]PolyPlan) map[string]int {
	newMap := map[string]int{}
	for key, value := range templateMap {
		newMap[polyPlans[key].pair1] += value
		newMap[polyPlans[key].pair2] += value
	}
	return newMap
}

func Part1(templateMap map[string]int, polyPlans map[string]PolyPlan, beginning, end string, cycles int) int {

	for i := 0; i < cycles; i++ {
		templateMap = TakeTurn(templateMap, polyPlans)
	}
	totalLetters := CalcString(templateMap, beginning, end)
	max := 0
	min := 999999999999999999
	for _, value := range totalLetters {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return max - min
}

func main() {
	startTime := time.Now()
	input, inserts, beginning, end := LoadInput("input.txt")
	part1Answer := Part1(input, inserts, beginning, end, 10)
	fmt.Printf("Part 1: %v\n", part1Answer)
	fmt.Println(time.Since(startTime))

	input2, inserts2, beginning2, end2 := LoadInput("input.txt")
	part2Answer := Part1(input2, inserts2, beginning2, end2, 40)
	fmt.Printf("Part 2: %v\n", part2Answer)
	fmt.Println(time.Since(startTime))
}
