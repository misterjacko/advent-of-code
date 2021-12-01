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

type deer struct {
	name      string
	raceSlice []int
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

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func TotalCycles(fly int, rest int, duration int) []int {
	cycleTime := fly + rest
	completeCycles := duration / cycleTime
	extraSeconds := duration % cycleTime
	if extraSeconds >= fly {
		return []int{completeCycles, fly}
	} else {
		return []int{completeCycles, extraSeconds}
	}
}

func Part1(data []string, duration int) int {
	farthestDistance := 0

	for _, line := range data {
		reindeer := Array(line)
		sprint := StringToInt(string(reindeer[3]))
		fly := StringToInt(string(reindeer[6]))
		rest := StringToInt(string(reindeer[13]))
		cycles := TotalCycles(fly, rest, duration)
		distance := (sprint * fly * cycles[0]) + (sprint * cycles[1])
		if distance > farthestDistance {
			farthestDistance = distance
		}
	}
	return farthestDistance
}

func Pt2Cycles(fly int, rest int, duration int) int {
	cycleTime := fly + rest
	Cycles := duration / cycleTime
	extraSeconds := duration % cycleTime
	if extraSeconds > 0 {
		return Cycles + 1
	} else {
		return Cycles
	}
}

func MakeSlice(dist int, fly int, rest int, cycles int) []int {
	returnSlice := []int{0}
	distance := 0
	for cycle := 0; cycle < cycles; cycle++ {
		for f := 0; f < fly; f++ {
			distance += dist
			returnSlice = append(returnSlice, distance)
		}
		for r := 0; r < rest; r++ {
			returnSlice = append(returnSlice, distance)
		}
	}
	return returnSlice
}

func RaceDeer(data []string, duration int) []*deer {
	allTheDeer := []*deer{}
	for _, reindeer := range data {
		arrayReindeer := Array(reindeer)
		dist := StringToInt(string(arrayReindeer[3]))
		fly := StringToInt(string(arrayReindeer[6]))
		rest := StringToInt(string(arrayReindeer[13]))
		cycles := Pt2Cycles(fly, rest, duration)
		d := deer{
			name:      arrayReindeer[0],
			raceSlice: MakeSlice(dist, fly, rest, cycles),
		}
		allTheDeer = append(allTheDeer, &d)

	}
	return allTheDeer

}

func Part2(data []string, duration int) int {
	winnersMap := map[string]int{}
	results := RaceDeer(data, duration)

	for i := 1; i < (duration + 1); i++ {
		maxDistance := 0
		maxDeer := []string{}
		for _, reindeer := range results {
			if reindeer.raceSlice[i] > maxDistance {
				maxDistance = reindeer.raceSlice[i]
				maxDeer = []string{reindeer.name}
			} else if reindeer.raceSlice[i] == maxDistance {
				maxDeer = append(maxDeer, reindeer.name)
			}
		}
		for _, reindeer := range maxDeer {
			if _, ok := winnersMap[reindeer]; ok {
				winnersMap[reindeer] += 1
			} else {
				winnersMap[reindeer] = 1
			}
		}
	}
	maxWins := 0
	for _, deer := range winnersMap {
		if deer > maxWins {
			maxWins = deer
		}
	}
	return maxWins
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data, 2503))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data, 2503))
	fmt.Println(time.Since(startTime))
}
