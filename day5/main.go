package main

import (
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type oceanFloor struct {
	floor [10][10]int
}

func parseInput() []line {
	data, err := os.ReadFile("test.txt")
	checkError(err)

	var lines []line

	inputString := strings.Replace(string(data), " -> ", ",", -1)
	inputStringSplit := strings.Split(inputString, "\n")

	for x := range inputStringSplit {
		lineString := strings.Split(inputStringSplit[x], ",")
		x1, err := strconv.Atoi(lineString[0])
		checkError(err)
		y1, err := strconv.Atoi(lineString[1])
		checkError(err)
		x2, err := strconv.Atoi(lineString[2])
		checkError(err)
		y2, err := strconv.Atoi(lineString[3])
		checkError(err)

		line := line{x1: x1, y1: y1, x2: x2, y2: y2}
		lines = append(lines, line)
	}

	return lines
}

func main() {
	parseInput()
	// part1()
	// part2()
}
