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
		lines = append(lines, line)
	}

	return lines
}

func findStraightLineCoordinates(line line, axis string) []coordinate {
	var coordinates []coordinate

	if axis == "x" {
		if line.coordinate1.y > line.coordinate2.y {
			for i := line.coordinate2.y; i < line.coordinate1.y+1; i++ {
				newCoordinate := coordinate{x: line.coordinate1.x, y: i}
				coordinates = append(coordinates, newCoordinate)
			}
		} else if line.coordinate1.y < line.coordinate2.y {
			for i := line.coordinate1.y; i < line.coordinate2.y+1; i++ {
				newCoordinate := coordinate{x: line.coordinate1.x, y: i}
				coordinates = append(coordinates, newCoordinate)
			}
		}
	} else if axis == "y" {
		if line.coordinate1.x > line.coordinate2.x {
			for i := line.coordinate2.x; i < line.coordinate1.x+1; i++ {
				newCoordinate := coordinate{x: i, y: line.coordinate1.y}
				coordinates = append(coordinates, newCoordinate)
			}
		} else if line.coordinate1.x < line.coordinate2.x {
			for i := line.coordinate1.x; i < line.coordinate2.x+1; i++ {
				newCoordinate := coordinate{x: i, y: line.coordinate1.y}
				coordinates = append(coordinates, newCoordinate)
			}
		}
	}

	// fmt.Println("\n", line)
	// fmt.Println(coordinates)
	return coordinates
}

func findDiagonalLineCoordinates(coordinates []coordinate, line line, axis string) []coordinate {
	var diagonalCoordinates []coordinate

	if axis == "x" {
		if line.coordinate1.y > line.coordinate2.y {
			for i := line.coordinate2.y; i < line.coordinate1.y+1; i++ {
				newCoordinate := coordinate{x: line.coordinate1.x, y: i}
				diagonalCoordinates = append(diagonalCoordinates, newCoordinate)
			}
		} else if line.coordinate1.y < line.coordinate2.y {
			for i := line.coordinate1.y; i < line.coordinate2.y+1; i++ {
				newCoordinate := coordinate{x: line.coordinate1.x, y: i}
				diagonalCoordinates = append(diagonalCoordinates, newCoordinate)
			}
		}
	} else if axis == "y" {
		if line.coordinate1.x > line.coordinate2.x {
			for i := line.coordinate2.x; i < line.coordinate1.x+1; i++ {
				newCoordinate := coordinate{x: i, y: line.coordinate1.y}
				diagonalCoordinates = append(diagonalCoordinates, newCoordinate)
			}
		} else if line.coordinate1.x < line.coordinate2.x {
			for i := line.coordinate1.x; i < line.coordinate2.x+1; i++ {
				newCoordinate := coordinate{x: i, y: line.coordinate1.y}
				diagonalCoordinates = append(diagonalCoordinates, newCoordinate)
			}
		}
	}

	// fmt.Println("\n", line)
	// fmt.Println(diagonalCoordinates)
	return coordinates
}

func part1() {
	fmt.Println("PART 1")
	lines := parseInput()
	// var oceanFloor [10][10]int
	var oceanFloor [1000][1000]int
	// fmt.Println(oceanFloor)

	for _, line := range lines {
		if line.coordinate1.x == line.coordinate2.x {
			coordinates := findStraightLineCoordinates(line, "x")

			for i := range coordinates {
				oceanFloor[coordinates[i].x][coordinates[i].y]++
			}
		}

		if line.coordinate1.y == line.coordinate2.y {
			coordinates := findStraightLineCoordinates(line, "y")

			for i := range coordinates {
				oceanFloor[coordinates[i].x][coordinates[i].y]++
			}
		}
	}

	var count int = 0
	for row := range oceanFloor {
		for column := range oceanFloor {
			if oceanFloor[row][column] > 1 {
				count++
			}
		}
	}

	// fmt.Println(oceanFloor)
	fmt.Println(count)
	fmt.Println()
}

func part2() {
	fmt.Println("PART 2")
	lines := parseInput()
	// var oceanFloor [10][10]int
	var oceanFloor [1000][1000]int
	// fmt.Println(oceanFloor)

	for _, line := range lines {
		if line.coordinate1.x == line.coordinate2.x {
			coordinates := findStraightLineCoordinates(line, "x")

			for i := range coordinates {
				oceanFloor[coordinates[i].x][coordinates[i].y]++
			}
		}

		if line.coordinate1.y == line.coordinate2.y {
			coordinates := findStraightLineCoordinates(line, "y")

			for i := range coordinates {
				oceanFloor[coordinates[i].x][coordinates[i].y]++
			}
		}
	}

	var count int = 0
	for row := range oceanFloor {
		for column := range oceanFloor {
			if oceanFloor[row][column] > 1 {
				count++
			}
		}
	}

	// fmt.Println(oceanFloor)
	fmt.Println(count)
	fmt.Println()
}

func main() {
	part1()
	part2()
}
