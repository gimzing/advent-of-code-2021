package main

import (
	"fmt"
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

		fmt.Println(len(boardSplit))

		// for j := range boardSplit {
		// 	fmt.Println(boardSplit[j])
		// }

		// var actualBoard []bingoNumber

		// boards = append(boards, board)
	}

	return numbers, boards
}

func part1() {
	// numbers, boards = parseInput()
	// fmt.Println("numbers", numbers)
	// fmt.Println("boards", boards)
}

// func part2() {

// }

func main() {
	parseInput()
	// part1()
	// part2()
}
