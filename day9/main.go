package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cavePoint struct {
	x       int
	y       int
	height  int
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
		if floor[i][j].height < floor[i+1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j+1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if i == len(floor)-1 && j == len(floor[0])-1 {
		if floor[i][j].height < floor[i-1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j-1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if i == 0 && j == len(floor[0])-1 {
		if floor[i][j].height < floor[i+1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j-1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if j == 0 && i == len(floor[0])-1 {
		if floor[i][j].height < floor[i][j+1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i-1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if i == len(floor)-1 && j == 0 {
		if floor[i][j].height < floor[i][j+1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i-1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if j == len(floor[0])-1 && i == 0 {
		if floor[i][j].height < floor[i+1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j-1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if i == 0 {
		if floor[i][j].height < floor[i+1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j+1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j-1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if j == 0 {
		if floor[i][j].height < floor[i+1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j+1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i-1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if i == len(floor)-1 {
		if floor[i][j].height < floor[i][j+1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i-1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j-1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else if j == len(floor[0])-1 {
		if floor[i][j].height < floor[i+1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i-1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j-1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	} else {
		if floor[i][j].height < floor[i+1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j+1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i-1][j].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
		if floor[i][j].height < floor[i][j-1].height {
			if !floor[i][j].inBasin {
				floor[i][j].inBasin = true
				newPoint = true
			}
		}
	}

	return floor, newPoint
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

	getLowestCavePoints(floor)

	hadNewBasinPoint := true
	for hadNewBasinPoint {
		atLeastOneNewPoint := false
		for i := range floor {
			for j := range floor[i] {
				var hadNewPoint bool
				floor, hadNewPoint = getAdjacentBasinPoints(floor, floor[i][j])
				if hadNewPoint {
					atLeastOneNewPoint = true
				}
			}
		}

		hadNewBasinPoint = atLeastOneNewPoint
	}

	printFloor(floor)
	fmt.Println()
}

func main() {
	input := parseInput("test.txt")

	// part1(input)
	part2(input)
}
