package main

import (
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type bingoNumber struct {
	number string
	seen   bool
}

// func parseInput() {
func parseInput() ([]string, [][]bingoNumber) {
	data, err := os.ReadFile("input.txt")
	checkError(err)
	inputStrings := strings.Split(string(data), "\n\n")

	var numbers []string = strings.Split(inputStrings[0], ",")
	var boardsString []string = inputStrings[1:]
	var boards [][]bingoNumber

	for i := range boardsString {
		var boardString string = boardsString[i]
		boardString = strings.Replace(boardString, "\r ", "", -1)
		boardString = strings.Replace(boardString, "\n ", "\n", -1)
		boardString = strings.Replace(boardString, "  ", " ", -1)

		if string(boardString[0]) == " " {
			boardString = strings.Replace(boardString, " ", "", 1)
		}

		var boardSplit []string = strings.Split(boardString, " ")

		var currentChunk []bingoNumber
		for x := range boardSplit {
			var bingoNum bingoNumber = bingoNumber{number: boardSplit[x], seen: false}
			currentChunk = append(currentChunk, bingoNum)

			if len(currentChunk) == 5 {
				boards = append(boards, currentChunk)
				currentChunk = nil
			}
		}
	}

	return numbers, boards
}

func checkWinner(boards [][]bingoNumber, number string) {

}

func part1() {
	numbers, boards := parseInput()

	for i := range numbers {
		for x := range boards {
			for y := range boards[x] {
				if numbers[i] == boards[x][y].number {
					boards[x][y].seen = true

					checkWinner(boards, numbers[i])
				}
			}
		}
	}
}

// func part2() {

// }

func main() {
	part1()
	// part2()
}
