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

func GrowOne(data [][]int) [][]int {
	for i, row := range data {
		for j, _ := range row {
			data[i][j]++
		}
	}
	return data
}

type Coord struct {
	row int
	col int
}

func GetNeighbors(row, col int) [][]int {
	trueNeighbors := [][]int{}
	neighbors := [][]int{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}
	for _, neighbor := range neighbors {
		if neighbor[0] >= 0 && neighbor[0] <= 9 && neighbor[1] >= 0 && neighbor[1] <= 9 {
			trueNeighbors = append(trueNeighbors, neighbor)
		}
	}
	return trueNeighbors
}

func Bloom(data [][]int, row, col int) [][]int {
	if _, ok := flashed[Coord{
		row: row,
		col: col,
	}]; !ok {
		flashed[Coord{
			row: row,
			col: col,
		}] = true
		neighbors := GetNeighbors(row, col)
		for _, neighbor := range neighbors {
			data[neighbor[0]][neighbor[1]]++
			if data[neighbor[0]][neighbor[1]] > 9 {
				Bloom(data, neighbor[0], neighbor[1])
			}
		}
	}
	return data
}

var data1 [][]int = (LoadInput("input_test.txt"))
var data2 [][]int = (LoadInput("input_test.txt"))
var flashed = make(map[Coord]bool)

func GrowNeighbors(data [][]int) [][]int {
	flashed = map[Coord]bool{}

	for i, row := range data {
		for j, value := range row {
			if value > 9 {
				Bloom(data, i, j)
			}
		}
	}

	return data
}

var Counter int = 0

func Part1(data [][]int, cycles int) int {
	Counter = 0

	for i := 0; i < cycles; i++ {
		//do first solo growth
		data = GrowOne(data)

		// expand on
		data = GrowNeighbors(data)
		for i, row := range data {
			for j, value := range row {
				if value > 9 {
					data[i][j] = 0
					Counter++
				}
			}

		}
	}
	return Counter
}
func Part2(data [][]int) int {
	Counter = 0
	turn := 0
	for {
		turn++
		flashedThisTurn := 0
		//do first solo growth
		data = GrowOne(data)

		// expand on
		data = GrowNeighbors(data)
		for i, row := range data {
			for j, value := range row {
				if value > 9 {
					data[i][j] = 0
					flashedThisTurn++
				}
			}

		}
		if flashedThisTurn == 100 {
			return turn
		}
	}
}

func main() {
	startTime := time.Now()
	data1 = LoadInput("input.txt")
	data2 = LoadInput("input.txt")
	partOneAnswer := Part1(data1, 100)
	fmt.Printf("Part 1: %v\n", partOneAnswer)
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data2))
	fmt.Println(time.Since(startTime))
}
