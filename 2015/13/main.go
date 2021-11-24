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
	str = strings.Replace(str, ".", "", -1)
	return strings.Split(str, " ")
}

func MakeSet(locations [][]string) []string {
	aMap := map[string]bool{}
	for _, line := range locations {
		if _, ok := aMap[line[0]]; !ok {
			aMap[line[0]] = true
		}
		if _, ok := aMap[line[10]]; !ok {
			aMap[line[10]] = true
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

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func MakeRelationships(instruction [][]string) map[string]int {
	relationshipMap := map[string]int{}
	for _, line := range instruction {
		happyValue := StringToInt(line[3])
		if line[2] == "lose" {
			happyValue *= -1
		}
		keyVals := MakeMapKey(line[0], line[10])
		if val, ok := relationshipMap[keyVals[0]]; !ok {
			relationshipMap[keyVals[0]] = happyValue
			relationshipMap[keyVals[1]] = happyValue
		} else {
			relationshipMap[keyVals[0]] = happyValue + val
			relationshipMap[keyVals[1]] = happyValue + val
		}
	}
	return relationshipMap
}

func Part1(data []string) int {
	arrayData := [][]string{}
	for _, line := range data {
		arrayData = append(arrayData, Array(line))
	}
	GuestSet := MakeSet(arrayData)
	arrangementCombos := MakeCombinations(GuestSet)
	relationships := MakeRelationships(arrayData)
	bestArrangement := 0
	for i, combo := range arrangementCombos {
		arrangementHappiness := 0
		for j := 0; j < len(combo)-1; j++ {
			key := MakeKey(combo[j], combo[j+1])
			arrangementHappiness += relationships[key]
		}
		// last set is last index to first index (table is a circle)
		key := MakeKey(combo[len(combo)-1], combo[0])
		arrangementHappiness += relationships[key]
		if i == 0 {
			bestArrangement = arrangementHappiness
		} else {
			if arrangementHappiness > bestArrangement {
				bestArrangement = arrangementHappiness
			}
		}
	}
	return bestArrangement
}

func Part2(data []string) int {
	arrayData := [][]string{}
	for _, line := range data {
		arrayData = append(arrayData, Array(line))
	}
	GuestSet := MakeSet(arrayData)
	arrangementCombos := MakeCombinations(GuestSet)
	relationships := MakeRelationships(arrayData)
	bestArrangement := 0
	for i, combo := range arrangementCombos {
		arrangementHappiness := 0
		for j := 0; j < len(combo)-1; j++ {
			key := MakeKey(combo[j], combo[j+1])
			arrangementHappiness += relationships[key]
		}
		if i == 0 {
			bestArrangement = arrangementHappiness
		} else {
			if arrangementHappiness > bestArrangement {
				bestArrangement = arrangementHappiness
			}
		}
	}
	return bestArrangement
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))
}
