package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	combo "github.com/ernestosuarez/itertools"
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
	return strings.Split(str, " ")
}

func MakeSet(locations [][]string) []string {
	aMap := map[string]bool{}
	for _, line := range locations {
		if _, ok := aMap[line[0]]; !ok {
			aMap[line[0]] = true
		}
		if _, ok := aMap[line[2]]; !ok {
			aMap[line[2]] = true
		}
	}
	set := []string{}
	for key, _ := range aMap {
		set = append(set, key)
	}
	return set
}

func MakeCombinations(locationList []string) [][]string {
	r := len(locationList)
	combinations := [][]string{}
	for v := range combo.PermutationsStr(locationList, r) {
		combinations = append(combinations, v)
	}
	return combinations
}

func MakeMapKey(str1 string, str2 string) []string {
	names := []string{
		fmt.Sprintf("%s-%s", str1, str2), fmt.Sprintf("%s-%s", str2, str1),
	}
	return names
}

func MakeKey(str1 string, str2 string) string {
	return fmt.Sprintf("%s-%s", str1, str2)
}

func MakeRoutes(instruction [][]string) map[string]int {
	routeMap := map[string]int{}
	for _, line := range instruction {
		keyVals := MakeMapKey(line[0], line[2])
		if _, ok := routeMap[keyVals[0]]; !ok {
			routeMap[keyVals[0]] = StringToInt(line[4])
			routeMap[keyVals[1]] = StringToInt(line[4])
		}
	}
	return routeMap
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func Part1(data []string) (int, int) {
	arrayData := [][]string{}
	for _, line := range data {
		arrayData = append(arrayData, Array(line))
	}
	locationSet := MakeSet(arrayData)
	locationCombos := MakeCombinations(locationSet)
	routes := MakeRoutes(arrayData)
	shortestRoute := 0
	longestRoute := 0
	for i, combo := range locationCombos {
		comboLength := 0
		for j := 0; j < len(combo)-1; j++ {
			key := MakeKey(combo[j], combo[j+1])
			comboLength += routes[key]
		}
		if i == 0 {
			shortestRoute = comboLength
			longestRoute = comboLength
		} else {
			if comboLength < shortestRoute {
				shortestRoute = comboLength
			}
			if comboLength > longestRoute {
				longestRoute = comboLength
			}
		}
	}
	return shortestRoute, longestRoute
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	pt1, pt2 := Part1(data)
	fmt.Printf("Part 1: %v\n", pt1)
	fmt.Printf("Part 2: %v\n", pt2)
	fmt.Println(time.Since(startTime))
}
