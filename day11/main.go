package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type octopus struct {
	energy     int
	x          int
	y          int
	hasFlashed bool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(filename string) [][]octopus {
	data, err := os.ReadFile(filename)
	checkError(err)

	var octopuses [][]octopus

	lines := strings.Split(string(data), "\n")
	for i := range lines {
		var octopusRow []octopus
		for j := range lines[i] {

			value, err := strconv.Atoi(string(lines[i][j]))
			checkError(err)

			var octopus octopus = octopus{x: j, y: i, energy: value}
			octopusRow = append(octopusRow, octopus)
		}
		octopuses = append(octopuses, octopusRow)
	}

	return octopuses
}

func printOctopuses(octopuses [][]octopus) {
	for i := range octopuses {
		for j := range octopuses[i] {
			fmt.Print(octopuses[i][j].energy)
		}
		fmt.Println()
	}
	fmt.Println()
}

func increaseEnergyAll(octopuses *[][]octopus) {
	for i := range *octopuses {
		for j := range (*octopuses)[i] {
			(*octopuses)[i][j].energy++
		}
	}
}

func resetEnergyAll(octopuses *[][]octopus) {
	for i := range *octopuses {
		for j := range (*octopuses)[i] {
			if (*octopuses)[i][j].hasFlashed {
				(*octopuses)[i][j].energy = 0
				(*octopuses)[i][j].hasFlashed = false
			}
		}
	}
}

func increaseEnergyAdjacent(octopuses *[][]octopus, x int, y int) {
	maxHeight := len(*octopuses)
	maxWidth := len((*octopuses)[0])

	// UP
	if y-1 >= 0 {
		(*octopuses)[y-1][x].energy++
	}
	// DOWN
	if y+1 < maxHeight {
		(*octopuses)[y+1][x].energy++
	}
	// RIGHT
	if x+1 < maxWidth {
		(*octopuses)[y][x+1].energy++
	}
	// LEFT
	if x-1 >= 0 {
		(*octopuses)[y][x-1].energy++
	}
	// UP-LEFT
	if y-1 >= 0 && x-1 >= 0 {
		(*octopuses)[y-1][x-1].energy++
	}
	// UP-RIGHT
	if y-1 >= 0 && x+1 < maxWidth {
		(*octopuses)[y-1][x+1].energy++
	}
	// DOWN-LEFT
	if y+1 < maxHeight && x-1 >= 0 {
		(*octopuses)[y+1][x-1].energy++
	}
	// DOWN-RIGHT
	if y+1 < maxHeight && x+1 < maxWidth {
		(*octopuses)[y+1][x+1].energy++
	}
}

func calculateFlashes(octopuses *[][]octopus) (int, bool) {
	atLeastOneFlash := false
	flashes := 0
	for i := range *octopuses {
		for j := range (*octopuses)[i] {
			if (*octopuses)[i][j].energy > 9 {
				if !(*octopuses)[i][j].hasFlashed {
					increaseEnergyAdjacent(octopuses, j, i)
					(*octopuses)[i][j].hasFlashed = true
					flashes++
					atLeastOneFlash = true
				}
			}
		}
	}

	return flashes, atLeastOneFlash
}

func checkAllFlashed(octopuses *[][]octopus) bool {
	allFlashed := true
	for i := range *octopuses {
		for j := range (*octopuses)[i] {
			if !(*octopuses)[i][j].hasFlashed {
				allFlashed = false
			}
		}
	}

	return allFlashed
}

func part1(octopuses [][]octopus) {
	fmt.Println("Part 1")

	totalFlashes := 0
	step := 0
	for step < 100 {
		increaseEnergyAll(&octopuses)

		atLeastOneFlash := true
		for atLeastOneFlash {
			atLeastOneFlash = false
			flashes, atLeastOneFlashOccured := calculateFlashes(&octopuses)
			if atLeastOneFlashOccured {
				atLeastOneFlash = true
			}
			totalFlashes += flashes
		}

		resetEnergyAll(&octopuses)
		step++
	}

	fmt.Println(totalFlashes)
	fmt.Println()
}

func part2(octopuses [][]octopus) {
	fmt.Println("Part 2")

	totalFlashes := 0
	step := 0
	allFlashed := false
	for !allFlashed {
		increaseEnergyAll(&octopuses)

		atLeastOneFlash := true
		for atLeastOneFlash {
			atLeastOneFlash = false
			flashes, atLeastOneFlashOccured := calculateFlashes(&octopuses)
			if atLeastOneFlashOccured {
				atLeastOneFlash = true
			}
			totalFlashes += flashes
		}

		allFlashed = checkAllFlashed(&octopuses)
		if allFlashed {
			fmt.Println(step + 1)
		}
		resetEnergyAll(&octopuses)
		step++
	}
}

func main() {
	input := parseInput("input.txt")

	part1(input)
	part2(input)
}
