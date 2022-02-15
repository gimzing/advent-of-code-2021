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

type coordinate struct {
	x int
	y int
}

type line struct {
	coordinate1 coordinate
	coordinate2 coordinate
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

		line := line{coordinate1: coordinate{x: x1, y: y1}, coordinate2: coordinate{x: x2, y: y2}}
		// line := line{x1: x1, y1: y1, x2: x2, y2: y2}
		lines = append(lines, line)
	}

	return lines
}

func findLineCoordinates(line line, axis string) []coordinate {
	var coordinates []coordinate

	if axis == "x" {

		if line.coordinate1.y > line.coordinate2.y {

			for i := line.coordinate2.y; i < line.coordinate1.y; i++ {
				fmt.Println("i", i)
			}

		} else if line.coordinate1.y < line.coordinate2.y {

			for i := line.coordinate1.y; i < line.coordinate2.y; i++ {
				fmt.Println("i", i)
			}
		}

	}
	// else if axis == "y" {

	// }

	return coordinates
}

func main() {
	lines := parseInput()
	var oceanFloor [10][10]int
	fmt.Println(oceanFloor)
	// floor := oceanFloor{}
	// fmt.Println(lines)
	// fmt.Println(floor)

	for _, line := range lines {
		if line.coordinate1.x == line.coordinate2.x {
			findLineCoordinates(line, "x")
		}

		if line.coordinate1.y == line.coordinate2.y {
			findLineCoordinates(line, "y")
		}
	}
}
