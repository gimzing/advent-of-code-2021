package main

// Part 1
// 452

// Part 2
// 1263735

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cavePoint struct {
	height  int
	x       int
	y       int
	inBasin bool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(filename string) [][]cavePoint {
	data, err := os.ReadFile(filename)
	checkError(err)

	var cavePoints [][]cavePoint

	rows := strings.Split(string(data), "\n")
	for i := range rows {
		var cavePointArray []cavePoint
		for j := range rows[i] {
			number, err := strconv.Atoi(string(rows[i][j]))
			checkError(err)
			cavePoint := cavePoint{x: j, y: i, height: number, inBasin: false}

			cavePointArray = append(cavePointArray, cavePoint)
		}

		cavePoints = append(cavePoints, cavePointArray)
	}

	return cavePoints
}

func getLowestCavePoints(floor [][]cavePoint) []cavePoint {
	var lowestCavePoints []cavePoint
	maxHeight := len(floor)
	maxWidth := len(floor[0])

	for i := range floor {
		for j := range floor[i] {
			// UP
			lowestPoint := true
			if i-1 >= 0 {
				if floor[i][j].height >= floor[i-1][j].height {
					lowestPoint = false
				}
			}
			// DOWN
			if i+1 < maxHeight {
				if floor[i][j].height >= floor[i+1][j].height {
					lowestPoint = false
				}
			}
			// LEFT
			if j-1 >= 0 {
				if floor[i][j].height >= floor[i][j-1].height {
					lowestPoint = false
				}
			}
			// RIGHT
			if j+1 < maxWidth {
				if floor[i][j].height >= floor[i][j+1].height {
					lowestPoint = false
				}
			}

			if lowestPoint {
				floor[i][j].inBasin = true
				lowestCavePoints = append(lowestCavePoints, floor[i][j])
			}
		}
	}

	return lowestCavePoints
}

func getAdjacentBasinPoints(floor [][]cavePoint, cavePoint cavePoint) ([][]cavePoint, bool) {
	i := cavePoint.y
	j := cavePoint.x
	maxHeight := len(floor)
	maxWidth := len(floor[0])
	newPoint := false

	// UP
	if i-1 >= 0 {
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
	}
	// DOWN
	if i+1 < maxHeight {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
	}
	// LEFT
	if j-1 >= 0 {
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	}
	// RIGHT
	if j+1 < maxWidth {
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
	}

	return floor, newPoint
}

func findAdjacentBasinPoints(floor [][]cavePoint, point cavePoint) []cavePoint {
	var points []cavePoint
	i := point.y
	j := point.x
	maxHeight := len(floor)
	maxWidth := len(floor[0])

	// UP
	if i-1 >= 0 {
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
	}
	// DOWN
	if i+1 < maxHeight {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
	}
	// LEFT
	if j-1 >= 0 {
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
		}
	}
	// RIGHT
	if j+1 < maxWidth {
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
	}

	return points
}

func printFloor(floor [][]cavePoint) {
	for i := range floor {
		for j := range floor[i] {
			if floor[i][j].inBasin {
				fmt.Printf("[%d]", floor[i][j].height)
			} else {
				fmt.Printf(" %d ", floor[i][j].height)
			}
		}
		fmt.Println()
	}
}

func part1(floor [][]cavePoint) {
	fmt.Println("Part 1")

	cavePoints := getLowestCavePoints(floor)

	total := 0
	for _, x := range cavePoints {
		total += x.height + 1
	}

	fmt.Println(total)
	fmt.Println()
}

func part2(floor [][]cavePoint) {
	fmt.Println("Part 2")

	lowestPoints := getLowestCavePoints(floor)

	hadNewBasinPoint := true
	for hadNewBasinPoint {
		atLeastOneNewPoint := false
		for i := range floor {
			for j := range floor[i] {
				if floor[i][j].inBasin {
					var hadNewPoint bool
					floor, hadNewPoint = getAdjacentBasinPoints(floor, floor[i][j])
					if hadNewPoint {
						atLeastOneNewPoint = true
					}
				}
			}
		}

		hadNewBasinPoint = atLeastOneNewPoint
	}

	var basinSizes []int
	for i := range lowestPoints {
		currentPoint := lowestPoints[i]
		var basinPoints []cavePoint
		basinPoints = append(basinPoints, currentPoint)
		hadNewPoint := true
		for hadNewPoint {
			hadNewPoint = false
			var points []cavePoint

			for x := range basinPoints {
				newAdjacentPoints := findAdjacentBasinPoints(floor, basinPoints[x])
				for y := range newAdjacentPoints {
					points = append(points, newAdjacentPoints[y])
				}
			}

			for y := range points {
				unique := true
				for z := range basinPoints {
					if basinPoints[z].x == points[y].x && basinPoints[z].y == points[y].y {
						unique = false
					}
				}

				if unique {
					basinPoints = append(basinPoints, points[y])
					hadNewPoint = true
				}
			}
		}

		basinSizes = append(basinSizes, len(basinPoints))
	}

	sort.Ints(basinSizes)
	threeLargestBasins := basinSizes[len(basinSizes)-3:]
	fmt.Println(uint64(threeLargestBasins[0]) * uint64(threeLargestBasins[1]) * uint64(threeLargestBasins[2]))
	fmt.Println()
}

func main() {
	input := parseInput("input.txt")

	part1(input)
	part2(input)
}
