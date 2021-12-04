package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BingoCard struct {
	name       int
	cardData   [5][5]string
	rowWinTurn []int
	colWinTurn []int
	wins       []int
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}
func MakeInt(str string) int {
	intVal, err := strconv.Atoi(str)
	check(err)
	return intVal
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

func LoadInstructions(data []string) map[string]int {
	instructions := map[string]int{}
	instructionLine := strings.Split(data[0], ",")
	for index, key := range instructionLine {
		instructions[key] = index
	}
	return instructions
}

func MakeArray(data []string) [5][5]string {
	array := [5][5]string{}
	for i, line := range data {
		lineSlice := strings.Split(line, " ")
		j := 0
		for _, str := range lineSlice {
			if str == "" {
				continue
			} else {
				array[i][j] = str
				j++
			}
		}

	}
	return array
}

func FindWinners(cardData [5][5]string, instructions map[string]int) ([]int, []int) {
	rowWinners := []int{}
	colWinners := []int{0, 0, 0, 0, 0}
	for _, row := range cardData {
		rowMax := 0
		for index, col := range row {
			turn := instructions[col]
			if turn > rowMax {
				rowMax = turn
			}
			if turn+1 > colWinners[index] {
				colWinners[index] = turn + 1
			}
		}
		rowWinners = append(rowWinners, rowMax+1)
	}
	return rowWinners, colWinners
}

func NewBingoCard(cardLines []string, name int, instructions map[string]int) *BingoCard {
	cardArray := MakeArray(cardLines)
	row, col := FindWinners(cardArray, instructions)
	totalWins := append(row, col...)
	sort.Ints(totalWins)
	card := BingoCard{
		name:       name,
		cardData:   cardArray,
		rowWinTurn: row,
		colWinTurn: col,
		wins:       totalWins,
	}
	return &card
}

func MakeAllCards(data []string, instructions map[string]int) []*BingoCard {
	cardList := []*BingoCard{}

	for i := 1; i < len(data)-1; i++ {
		if data[i] == "" {
			continue
		}
		cardLines := []string{}
		for cardRow := 0; cardRow < 5; cardRow++ {
			cardLines = append(cardLines, data[i])
			i++
		}
		cardList = append(cardList, NewBingoCard(cardLines, i, instructions))
	}

	return cardList
}

func Part1(data []string) int {
	instructions := LoadInstructions(data)
	cards := MakeAllCards(data, instructions)
	earliestWinner := 1000
	winnerName := 1000
	for _, card := range cards {
		if card.wins[0] < earliestWinner {
			earliestWinner = card.wins[0]
			winnerName = card.name
		}
	}
	for _, card := range cards {
		if card.name == winnerName {
			missingTotal := 0
			lastNumber := 0
			for _, row := range card.cardData {
				for _, number := range row {
					if instructions[number] == earliestWinner-1 {
						lastNumber = MakeInt(number)
					}
					if instructions[number] >= earliestWinner {
						missingTotal += MakeInt(number)
					}

				}
			}
			return missingTotal * lastNumber
		}
	}
	return 999
}

func Part2(data []string) int {
	instructions := LoadInstructions(data)
	cards := MakeAllCards(data, instructions)
	latestWinner := 0
	winnerName := 1000
	for _, card := range cards {
		if card.wins[0] > latestWinner {
			latestWinner = card.wins[0]
			winnerName = card.name
		}
	}
	for _, card := range cards {
		if card.name == winnerName {
			missingTotal := 0
			lastNumber := 0
			for _, row := range card.cardData {
				for _, number := range row {
					if instructions[number] == latestWinner-1 {
						lastNumber = MakeInt(number)
					}
					if instructions[number] >= latestWinner {
						missingTotal += MakeInt(number)
					}

				}
			}
			return missingTotal * lastNumber
		}
	}
	return 999
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))
}
