package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lowPoint struct {
	height int
	x      int
	y      int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(filename string) [][]int {
	data, err := os.ReadFile(filename)
	checkError(err)

	var numbers [][]int

	rows := strings.Split(string(data), "\n")
	for i := range rows {
		var intArray []int
		for j := range rows[i] {
			number, err := strconv.Atoi(string(rows[i][j]))
			checkError(err)

			intArray = append(intArray, number)
		}

		numbers = append(numbers, intArray)
	}

	return numbers
}

func getLowPoints(floor [][]int) []lowPoint {
	var lowPoints []lowPoint

	for i := range floor {
		for j := range floor[i] {
			if i == 0 && j == 0 {
				if floor[i][j] < floor[i+1][j] &&
					floor[i][j] < floor[i][j+1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if i == len(floor)-1 && j == len(floor[0])-1 {
				if floor[i][j] < floor[i-1][j] &&
					floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if i == 0 && j == len(floor[0])-1 {
				if floor[i][j] < floor[i+1][j] &&
					floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if j == 0 && i == len(floor[0])-1 {
				if floor[i][j] < floor[i][j+1] &&
					floor[i][j] < floor[i-1][j] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if i == len(floor)-1 && j == 0 {
				if floor[i][j] < floor[i][j+1] &&
					floor[i][j] < floor[i-1][j] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if j == len(floor[0])-1 && i == 0 {
				if floor[i][j] < floor[i+1][j] &&
					floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if i == 0 {
				if floor[i][j] < floor[i+1][j] &&
					floor[i][j] < floor[i][j+1] &&
					floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if j == 0 {
				if floor[i][j] < floor[i+1][j] &&
					floor[i][j] < floor[i][j+1] &&
					floor[i][j] < floor[i-1][j] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if i == len(floor)-1 {
				if floor[i][j] < floor[i][j+1] &&
					floor[i][j] < floor[i-1][j] &&
					floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else if j == len(floor[0])-1 {
				if floor[i][j] < floor[i+1][j] &&
					floor[i][j] < floor[i-1][j] &&
					floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			} else {
				if floor[i][j] < floor[i+1][j] &&
					floor[i][j] < floor[i][j+1] &&
					floor[i][j] < floor[i-1][j] &&
					floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, lowPoint{height: floor[i][j], x: j, y: i})
				}
			}
		}
	}

	return lowPoints
}

func part1(input [][]int) {
	fmt.Println("Part 1")

	lowPoints := getLowPoints(input)

	total := 0
	for _, x := range lowPoints {
		total += x.height + 1
	}

	fmt.Println(total)
	fmt.Println()
}

func part2(input [][]int) {
	fmt.Println("Part 2")

	lowPoints := getLowPoints(input)

	for _, i := range lowPoints {
		fmt.Println(i)
	}
}

func main() {
	input := parseInput("test.txt")

	// part1(input)
	part2(input)
}
