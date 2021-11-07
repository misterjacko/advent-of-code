package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func FindFuel(mass int) int {
	floatMass := ConvertIntToFlo64(mass)
	// var fuel int = int(math.Floor(mass/3) - 2)
	fuel := int(math.Floor(floatMass/3) - 2)
	return fuel
}

func ConvertIntToFlo64(i int) float64 {
	var flo64 float64 = float64(i)
	return flo64
}
func LoadInput(fileName string) []int {
	input := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		input = append(input, line)
	}

	return input
}

func FuelForFuel(unaccountedForFuel int) int {
	totalNewFuel := 0

	newFuel := FindFuel(unaccountedForFuel)
	for newFuel >= 0 {
		totalNewFuel += newFuel
		newFuel = FindFuel(newFuel)

	}
	return totalNewFuel
}

func Part1(data []int) int {
	totalFuel := 0

	for _, module := range data {
		totalFuel += FindFuel(module)
	}
	return totalFuel
}

func Part2(data []int) int {
	totalFuel := 0

	for _, module := range data {
		totalFuel += FuelForFuel(module)
	}
	return totalFuel
}

func main() {
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Printf("Part 1: %v\n", Part2(data))
}
