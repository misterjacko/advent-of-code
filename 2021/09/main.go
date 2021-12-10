package main

import (
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
	return strings.Split(str, "")
}

func MakeInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

func LoadInput(fileName string) [][]int {
	input := [][]int{}

	file, err := os.ReadFile(fileName)
	check(err)
	data := strings.Split(string(file), "\n")

	for _, nums := range data {
		lineSlice := []int{}
		stringSlice := MakeSlice(nums)
		for _, str := range stringSlice {
			lineSlice = append(lineSlice, MakeInt(str))
		}
		input = append(input, lineSlice)
	}
	return input
}

func IsLowest(point int, neighbors []int) bool {
	for _, neighbor := range neighbors {
		if neighbor <= point {
			return false
		}
	}
	return true
}

func MapEnds(targetLine, refrenceLine []int) (int, []int) {
	total := 0
	column := []int{}
	if IsLowest(targetLine[0], []int{targetLine[1], refrenceLine[0]}) {
		total += targetLine[0] + 1
		column = append(column, 0)
	}
	for i := 1; i < len(targetLine)-1; i++ {
		if IsLowest(targetLine[i], []int{targetLine[i+1], targetLine[i-1], refrenceLine[i]}) {
			total += targetLine[i] + 1
			column = append(column, i)
		}
	}
	if IsLowest(targetLine[len(targetLine)-1], []int{targetLine[len(targetLine)-2], refrenceLine[len(targetLine)-1]}) {
		total += targetLine[len(targetLine)-1] + 1
		column = append(column, len(targetLine)-1)

	}
	return total, column
}

func MapCenter(targetLine, aboveLine, belowLine []int) (int, []int) {
	total := 0
	column := []int{}

	if IsLowest(targetLine[0], []int{targetLine[1], aboveLine[0], belowLine[0]}) {
		total += targetLine[0] + 1
		column = append(column, 0)

	}
	for i := 1; i < len(targetLine)-1; i++ {
		if IsLowest(targetLine[i], []int{targetLine[i-1], targetLine[i+1], aboveLine[i], belowLine[i]}) {
			total += targetLine[i] + 1
			column = append(column, i)

		}
	}
	if IsLowest(targetLine[len(targetLine)-1], []int{targetLine[len(targetLine)-2], aboveLine[len(targetLine)-1], belowLine[len(targetLine)-1]}) {
		total += targetLine[len(targetLine)-1] + 1
		column = append(column, len(targetLine)-1)

	}

	return total, column
}

type Coord struct {
	row int
	col int
}

func Part1(data [][]int) (int, []Coord) {
	riskLevel := 0
	lowestPoints := []Coord{}

	newRisk, newPoints := MapEnds(data[0], data[1])
	riskLevel += newRisk
	for _, column := range newPoints {
		lowestPoints = append(lowestPoints, Coord{
			row: 0,
			col: column,
		})
	}

	for i := 1; i < len(data)-1; i++ {
		newRisk, newPoints = MapCenter(data[i], data[i-1], data[i+1])
		riskLevel += newRisk
		for _, column := range newPoints {
			lowestPoints = append(lowestPoints, Coord{
				row: i,
				col: column,
			})
		}
	}
	newRisk, newPoints = MapEnds(data[len(data)-1], data[len(data)-2])
	riskLevel += newRisk
	for _, column := range newPoints {
		lowestPoints = append(lowestPoints, Coord{
			row: len(data) - 1,
			col: column,
		})
	}
	return riskLevel, lowestPoints
}

var GlobalPoints = make(map[Coord]bool)

var data = LoadInput("input.txt")

// var data = LoadInput("input_test.txt")
var count int

func CheckAdjacent(point Coord) int {
	// fmt.Println(point)
	count++
	totalSize := 0
	newNeighbors := []Coord{}
	GlobalPoints[point] = true
	// top and bottom
	if point.row == 0 {
		if data[point.row+1][point.col] < 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row + 1,
				col: point.col,
			})
		}
	} else if point.row == len(data)-1 {
		if data[point.row-1][point.col] < 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row - 1,
				col: point.col,
			})
		}
	} else {
		if data[point.row-1][point.col] < 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row - 1,
				col: point.col,
			})
		}
		if data[point.row+1][point.col] < 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row + 1,
				col: point.col,
			})
		}
	}

	// left and right
	if point.col == 0 {
		if data[point.row][point.col+1] < 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row,
				col: point.col + 1,
			})
		}
	} else if point.col == len(data[0])-1 {
		if data[point.row][point.col-1] < 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row,
				col: point.col - 1,
			})
		}
	} else {
		if data[point.row][point.col-1] != 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row,
				col: point.col - 1,
			})
		}
		if data[point.row][point.col+1] != 9 {
			newNeighbors = append(newNeighbors, Coord{
				row: point.row,
				col: point.col + 1,
			})
		}
	}
	// fmt.Println(newNeighbors)
	for _, point := range newNeighbors {

		if _, ok := GlobalPoints[point]; !ok {
			CheckAdjacent(point)
			totalSize++
		}
	}
	// fmt.Println(totalSize)
	return count

}

func Part2(data [][]int, lowestPoints []Coord) int {
	allSizes := []int{}
	for _, basin := range lowestPoints {
		count = 0
		allSizes = append(allSizes, CheckAdjacent(basin))
	}
	sort.Ints(allSizes)
	length := len(allSizes)
	total := allSizes[length-1] * allSizes[length-2] * allSizes[length-3]
	return total
}

func main() {
	// GlobalPoints = make(map[Coord]bool)
	startTime := time.Now()
	// data = LoadInput("input.txt")
	partOneAnswer, lowestPoints := Part1(data)
	fmt.Printf("Part 1: %v\n", partOneAnswer)
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data, lowestPoints))
	fmt.Println(time.Since(startTime))
}
