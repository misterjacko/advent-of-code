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

type Coord struct {
	col, row int
}

type Instruction struct {
	fold     string
	location int
}

func StringToInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
}

func LoadInput(fileName string) (map[Coord]int, []Instruction) {
	dots := map[Coord]int{}
	instructions := []Instruction{}
	file, err := os.ReadFile(fileName)
	instructionSwitch := false
	check(err)
	data := strings.Split(string(file), "\n")
	for _, line := range data {
		if line == "" {
			instructionSwitch = true
			continue
		}
		if !instructionSwitch { //dots
			lineSlice := strings.Split(line, ",")
			dots[Coord{
				col: StringToInt(lineSlice[0]),
				row: StringToInt(lineSlice[1]),
			}] = 1
		} else { // instructions
			line = strings.Replace(line, "fold along ", "", -1)
			lineSlice := strings.Split(line, "=")
			instructions = append(instructions, Instruction{
				fold:     lineSlice[0],
				location: StringToInt(lineSlice[1]),
			})
		}
	}
	return dots, instructions
}

func Part1(dots map[Coord]int, instructions []Instruction, cycles int) (int, map[Coord]int) {

	for i := 0; i < cycles; i++ {
		if instructions[i].fold == "y" {
			foldLocation := instructions[i].location
			for coord, dot := range dots {
				if coord.row > foldLocation {
					newRow := foldLocation - (coord.row - foldLocation)
					dots[Coord{
						col: coord.col,
						row: newRow,
					}] += dot
					delete(dots, coord)
				} else if coord.row == foldLocation {
					delete(dots, coord)

				}

			}
		}

		if instructions[i].fold == "x" {
			foldLocation := instructions[i].location
			for coord, dot := range dots {
				if coord.col > foldLocation {
					newCol := foldLocation - (coord.col - foldLocation)
					dots[Coord{
						col: newCol,
						row: coord.row,
					}] += dot
					delete(dots, coord)
				} else if coord.col == foldLocation {
					delete(dots, coord)

				}

			}
		}
	}
	return len(dots), dots
}

func Part2(olddots map[Coord]int, instructions []Instruction, cycles int) {
	_, dots := Part1(olddots, instructions, cycles)
	maxRow := 0
	maxCol := 0
	for dot, _ := range dots {
		if dot.col > maxCol {
			maxCol = dot.col
		}
		if dot.row > maxRow {
			maxRow = dot.row
		}
	}
	array := [][]string{}
	for i := 0; i <= maxRow; i++ {
		line := []string{}
		for j := 0; j <= maxCol; j++ {
			line = append(line, " ")
		}
		array = append(array, line)
	}
	for coord, _ := range dots {
		array[coord.row][coord.col] = "#"
	}
	for _, line := range array {
		fmt.Println(line)
	}
}

func main() {
	startTime := time.Now()
	dotInput, instructions := LoadInput("input.txt")
	part1Answer, _ := Part1(dotInput, instructions, 1)
	fmt.Printf("Part 1: %v\n", part1Answer)
	fmt.Println(time.Since(startTime))

	dotInput, instructions = LoadInput("input.txt")
	fmt.Println("Part 2:")
	Part2(dotInput, instructions, len(instructions))
	fmt.Println(time.Since(startTime))
}
