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
func MakeSlice(str string) []string {
	str = strings.Replace(str, " -> ", ",", -1)
	return strings.Split(str, ",")
}

func MakeInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

func MakeIntSlice(str []string) []int {
	intSlice := []int{}
	for _, num := range str {
		intSlice = append(intSlice, MakeInt(num))
	}
	return intSlice
}
func LoadInput(fileName string) [][]int {
	input := [][]int{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		block := scanner.Text()
		intSlice := MakeIntSlice(MakeSlice(block))
		input = append(input, intSlice)

	}
	return input
}

func Make10Array() [10][10]int {
	return [10][10]int{}
}

func Make1000Array() [1000][1000]int {
	return [1000][1000]int{}
}

func IsHorizontal(coords []int) bool {
	return coords[1] == coords[3]
}

func MarkHorizontal(coords []int, array [1000][1000]int) [1000][1000]int {
	var start, end int
	if coords[0] <= coords[2] {
		start = coords[0]
		end = coords[2]
	} else {
		start = coords[2]
		end = coords[0]
	}

	for i := start; i <= end; i++ {
		array[coords[1]][i]++
	}
	return array

}

func IsVertical(coords []int) bool {
	return coords[0] == coords[2]
}

func MarkVertical(coords []int, array [1000][1000]int) [1000][1000]int {
	var start, end int
	if coords[1] <= coords[3] {
		start = coords[1]
		end = coords[3]
	} else {
		start = coords[3]
		end = coords[1]
	}

	for i := start; i <= end; i++ {
		array[i][coords[0]]++
	}
	return array

}

func Part1(data [][]int, array [1000][1000]int) int {
	overlap := 0

	for _, line := range data {
		if IsHorizontal(line) {
			array = MarkHorizontal(line, array)
		} else if IsVertical(line) {
			array = MarkVertical(line, array)
		}
	}
	for _, row := range array {
		for _, col := range row {
			if col > 1 {
				overlap++
			}
		}
	}

	return overlap
}

func MarkDiagonal(coords []int, array [1000][1000]int) [1000][1000]int {

	if coords[0] > coords[2] {
		coords = []int{coords[2], coords[3], coords[0], coords[1]}
	}

	if coords[1] < coords[3] {
		j := coords[1]
		for i := coords[0]; i <= coords[2]; i++ {
			array[j][i]++
			j++
		}
	} else {
		j := coords[1]
		for i := coords[0]; i <= coords[2]; i++ {
			array[j][i]++
			j--
		}
	}
	return array
}

func Part2(data [][]int, array [1000][1000]int) int {
	overlap := 0

	for _, line := range data {
		if IsHorizontal(line) {
			array = MarkHorizontal(line, array)
		} else if IsVertical(line) {
			array = MarkVertical(line, array)
		} else {
			array = MarkDiagonal(line, array)
		}
	}
	for _, row := range array {
		for _, col := range row {
			if col >= 2 {
				overlap++
			}
		}
	}

	return overlap
}

func main() {
	array := Make1000Array()
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data, array))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data, array))
	fmt.Println(time.Since(startTime))
}
