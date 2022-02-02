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

type boardWithHasWon struct {
	board  [][]bingoNumber
	hasWon bool
}

func parseInput() ([]string, []boardWithHasWon) {
	data, err := os.ReadFile("test.txt")
	checkError(err)

	inputStrings := strings.Split(string(data), "\n\n")
	var numbers []string = strings.Split(inputStrings[0], ",")
	var boardsString []string = inputStrings[1:]

	var boards []boardWithHasWon

	for i := range boardsString {
		var boardString string = boardsString[i]
		boardString = strings.Replace(boardString, "  ", " ", -1)
		boardString = strings.Replace(boardString, "\r ", "\r", -1)
		boardString = strings.Replace(boardString, "\n ", "\n", -1)
		boardString = strings.Replace(boardString, " \n ", "\n", -1)

		if string(boardString[0]) == " " {
			boardString = strings.Replace(boardString, " ", "", 1)
		}

		var boardSplit []string = strings.Split(boardString, "\n")
		var actualBoard boardWithHasWon

		var actualBoardBoard [][]bingoNumber

		for j := range boardSplit {
			var currentChunk []bingoNumber
			var rows []string = strings.Split(boardSplit[j], " ")
			for k := range rows {
				var bingoNum bingoNumber = bingoNumber{number: rows[k], seen: false}
				currentChunk = append(currentChunk, bingoNum)
			}

			if len(currentChunk) == 5 {
				actualBoardBoard = append(actualBoardBoard, currentChunk)
				currentChunk = nil
			}
		}

		actualBoard = boardWithHasWon{board: actualBoardBoard, hasWon: false}

		boards = append(boards, actualBoard)
	}

	return numbers, boards
}

func getWinner(board boardWithHasWon, currentNumber string) {
	var totalUnseen int = 0
	for i := range board.board {
		for j := range board.board[i] {

			if !board.board[i][j].seen {
				value, err := strconv.Atoi(board.board[i][j].number)
				checkError(err)

				totalUnseen += value
			}
		}
	}

	value, err := strconv.Atoi(currentNumber)
	checkError(err)

	fmt.Println("totalUnseen", totalUnseen)
	fmt.Println("currentNumber", currentNumber)
	fmt.Println("answer", value*totalUnseen)
	fmt.Println()

	os.Exit(1)
}

func checkWinner(board boardWithHasWon, number string) bool {
	column := checkColumn(board)
	if column {
		return true
	}

	for i := range board.board {
		row := checkRow(board.board[i])
		if row {
			return true
		}
	}

	return false
}

func checkRow(row []bingoNumber) bool {
	var rowSeen bool = true

	for i := range row {
		if !row[i].seen {
			rowSeen = false
		}
	}

	return rowSeen
}

func checkColumn(board boardWithHasWon) bool {
	var columnSeen bool = true

	for i := 0; i < 5; i++ {
		for j := range board.board {
			if !board.board[j][i].seen {
				columnSeen = false
			}
		}

		if columnSeen {
			return true
		}
	}

	return false
}

func part1() {
	numbers, boards := parseInput()

	for i := range numbers {
		for x := range boards {
			for y := range boards[x].board {
				for z := range boards[x].board[y] {
					if numbers[i] == boards[x].board[y][z].number {
						boards[x].board[y][z].seen = true
					}
					if checkWinner(boards[x], numbers[i]) {
						getWinner(boards[x], numbers[i])
					}
				}
			}
		}
	}
}

func part2() {
	numbers, boards := parseInput()

	for i := range numbers {
		var lastBoardToWin boardWithHasWon
		for x := range boards {
			for y := range boards[x].board {
				for z := range boards[x].board[y] {
					if boards[x].board[y][z].number == numbers[i] {
						boards[x].board[y][z].seen = true

						if checkWinner(boards[x], numbers[i]) {
							boards[x].hasWon = true
							lastBoardToWin = boards[x]
						}
					}

				}
			}
		}

		var totalWinningBoards int = 0
		for x := range boards {
			if boards[x].hasWon {
				totalWinningBoards++
			}
		}

		if totalWinningBoards == len(boards) {
			getWinner(lastBoardToWin, numbers[i])
		}
	}
}

func main() {
	// part1()
	part2()
}
