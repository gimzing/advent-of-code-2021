package main

import (
	"fmt"
	"os"
	"strconv"
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

func parseInput() ([]string, [][][]bingoNumber) {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	inputStrings := strings.Split(string(data), "\n\n")
	var numbers []string = strings.Split(inputStrings[0], ",")
	var boardsString []string = inputStrings[1:]

	var boards [][][]bingoNumber

	for i := range boardsString {
		var boardString string = boardsString[i]
		boardString = strings.Replace(boardString, "  ", " ", -1)
		boardString = strings.Replace(boardString, "\r ", "\r", -1)
		boardString = strings.Replace(boardString, "\n ", "\n", -1)
		boardString = strings.Replace(boardString, " \n ", "\n", -1)

		if string(boardString[0]) == " " {
			boardString = strings.Replace(boardString, " ", "", 1)
		}

		// This is where it goes wrong
		var boardSplit []string = strings.Split(boardString, "\n")

		var actualBoard [][]bingoNumber

		for j := range boardSplit {
			var currentChunk []bingoNumber
			var rows []string = strings.Split(boardSplit[j], " ")
			for k := range rows {
				var bingoNum bingoNumber = bingoNumber{number: rows[k], seen: false}
				currentChunk = append(currentChunk, bingoNum)
			}

			if len(currentChunk) == 5 {
				actualBoard = append(actualBoard, currentChunk)
				currentChunk = nil
			}
		}

		boards = append(boards, actualBoard)
	}

	fmt.Println(boards)

	return numbers, boards
}

func getWinner(board [][]bingoNumber, number string) {
	fmt.Println("board", board)
	fmt.Println("len", len(board))
	fmt.Println("number", number)

	var totalUnseen int = 0
	for i := range board {
		for j := range board[i] {

			if board[i][j].seen == false {
				value, err := strconv.Atoi(board[i][j].number)
				checkError(err)

				totalUnseen += value
			}
		}
	}

	value, err := strconv.Atoi(number)
	checkError(err)

	fmt.Println("totalUnseen", totalUnseen)
	fmt.Println("answer", value*totalUnseen)

	os.Exit(1)
}

func checkWinner(board [][]bingoNumber, number string) {
	column := checkColumn(board)
	if column == true {
		getWinner(board, number)
	}

	row := false
	for i := range board {
		row = checkRow(board[i])
		if row == true {
			getWinner(board, number)
		}
	}
}

func checkRow(row []bingoNumber) bool {
	var rowSeen bool = true

	for i := range row {
		if row[i].seen == false {
			rowSeen = false
		}
	}

	return rowSeen
}

func checkColumn(board [][]bingoNumber) bool {
	var columnSeen bool = true

	for i := 0; i < 5; i++ {
		for j := range board {
			if board[j][i].seen == false {
				columnSeen = false
			}
		}

		if columnSeen == true {
			return true
		}
	}

	return false
}

func part1() {
	numbers, boards := parseInput()

	for i := range numbers {
		for x := range boards {
			for y := range boards[x] {
				for z := range boards[x][y] {
					if numbers[i] == boards[x][y][z].number {
						boards[x][y][z].seen = true
					}
					checkWinner(boards[x], numbers[i])
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
