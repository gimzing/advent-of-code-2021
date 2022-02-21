package main

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

	for i := range floor {
		for j := range floor[i] {
			if i == 0 && j == 0 {
				if floor[i][j].height < floor[i+1][j].height &&
					floor[i][j].height < floor[i][j+1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if i == len(floor)-1 && j == len(floor[0])-1 {
				if floor[i][j].height < floor[i-1][j].height &&
					floor[i][j].height < floor[i][j-1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if i == 0 && j == len(floor[0])-1 {
				if floor[i][j].height < floor[i+1][j].height &&
					floor[i][j].height < floor[i][j-1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if j == 0 && i == len(floor[0])-1 {
				if floor[i][j].height < floor[i][j+1].height &&
					floor[i][j].height < floor[i-1][j].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if i == len(floor)-1 && j == 0 {
				if floor[i][j].height < floor[i][j+1].height &&
					floor[i][j].height < floor[i-1][j].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if j == len(floor[0])-1 && i == 0 {
				if floor[i][j].height < floor[i+1][j].height &&
					floor[i][j].height < floor[i][j-1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if i == 0 {
				if floor[i][j].height < floor[i+1][j].height &&
					floor[i][j].height < floor[i][j+1].height &&
					floor[i][j].height < floor[i][j-1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if j == 0 {
				if floor[i][j].height < floor[i+1][j].height &&
					floor[i][j].height < floor[i][j+1].height &&
					floor[i][j].height < floor[i-1][j].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if i == len(floor)-1 {
				if floor[i][j].height < floor[i][j+1].height &&
					floor[i][j].height < floor[i-1][j].height &&
					floor[i][j].height < floor[i][j-1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else if j == len(floor[0])-1 {
				if floor[i][j].height < floor[i+1][j].height &&
					floor[i][j].height < floor[i-1][j].height &&
					floor[i][j].height < floor[i][j-1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			} else {
				if floor[i][j].height < floor[i+1][j].height &&
					floor[i][j].height < floor[i][j+1].height &&
					floor[i][j].height < floor[i-1][j].height &&
					floor[i][j].height < floor[i][j-1].height {
					floor[i][j].inBasin = true
					lowestCavePoints = append(lowestCavePoints, floor[i][j])
				}
			}
		}
	}

	return lowestCavePoints
}

func getAdjacentBasinPoints(floor [][]cavePoint, cavePoint cavePoint) ([][]cavePoint, bool) {
	i := cavePoint.y
	j := cavePoint.x
	newPoint := false

	if i == 0 && j == 0 {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
	} else if i == len(floor)-1 && j == len(floor[0])-1 {
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	} else if i == 0 && j == len(floor[0])-1 {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	} else if j == 0 && i == len(floor[0])-1 {
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
	} else if i == len(floor)-1 && j == 0 {
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
	} else if j == len(floor[0])-1 && i == 0 {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	} else if i == 0 {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	} else if j == 0 {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
	} else if i == len(floor)-1 {
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	} else if j == len(floor[0])-1 {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	} else {
		if floor[i+1][j].height != 9 && !floor[i+1][j].inBasin {
			floor[i+1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j+1].height != 9 && !floor[i][j+1].inBasin {
			floor[i][j+1].inBasin = true
			newPoint = true
		}
		if floor[i-1][j].height != 9 && !floor[i-1][j].inBasin {
			floor[i-1][j].inBasin = true
			newPoint = true
		}
		if floor[i][j-1].height != 9 && !floor[i][j-1].inBasin {
			floor[i][j-1].inBasin = true
			newPoint = true
		}
	}

	return floor, newPoint
}

// func appendAdjacentPoint(floor [][]cavePoint, currentPoint cavePoint, adjacentPoint cavePoint) [][]cavePoint {
// 	i := currentPoint.y
// 	j := currentPoint.x

// 	floor[i][j].adjacentBasinPoints = append(floor[i][j].adjacentBasinPoints, adjacentPoint)
// 	return floor
// }

func findAdjacentBasinPoints(floor [][]cavePoint, point cavePoint) []cavePoint {
	var points []cavePoint
	i := point.y
	j := point.x

	if i == 0 && j == 0 {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
	} else if i == len(floor)-1 && j == len(floor[0])-1 {
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
		}
	} else if i == 0 && j == len(floor[0])-1 {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
		}
	} else if j == 0 && i == len(floor[0])-1 {
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
	} else if i == len(floor)-1 && j == 0 {
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
	} else if j == len(floor[0])-1 && i == 0 {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
		}
	} else if i == 0 {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
		}
	} else if j == 0 {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
	} else if i == len(floor)-1 {
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
		}
	} else if j == len(floor[0])-1 {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
		}
	} else {
		if floor[i+1][j].inBasin {
			points = append(points, floor[i+1][j])
		}
		if floor[i][j+1].inBasin {
			points = append(points, floor[i][j+1])
		}
		if floor[i-1][j].inBasin {
			points = append(points, floor[i-1][j])
		}
		if floor[i][j-1].inBasin {
			points = append(points, floor[i][j-1])
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

	// printFloor(floor)
	fmt.Println(total)
	fmt.Println()
}

func part2(floor [][]cavePoint) {
	fmt.Println("Part 2")

	lowestPoints := getLowestCavePoints(floor)
	fmt.Println(len(lowestPoints))

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
	fmt.Println()
	// printFloor(floor)

	// lowestPoints := getLowestCavePoints(floor)

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

		if len(basinPoints) == 108 {
			fmt.Println(basinPoints)
		}

		basinSizes = append(basinSizes, len(basinPoints))
	}

	printFloor(floor)

	fmt.Println(len(basinSizes))
	fmt.Println()
	sort.Ints(basinSizes)
	fmt.Println(basinSizes)
	threeLargestBasins := basinSizes[len(basinSizes)-3:]
	fmt.Println(threeLargestBasins)

	fmt.Println(uint64(threeLargestBasins[0]) * uint64(threeLargestBasins[1]) * uint64(threeLargestBasins[2]))
	fmt.Println()
}

func main() {
	input := parseInput("input.txt")

	// part1(input)
	part2(input)
}
