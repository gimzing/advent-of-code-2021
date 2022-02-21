package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type octopus struct {
	energy int
	x      int
	y      int
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
}

func increaseEnergy(octopuses *[][]octopus) {
	for i := range *octopuses {
		for j := range (*octopuses)[i] {
			(*octopuses)[i][j].energy++
		}
	}
}

func part1(octopuses [][]octopus) {
	fmt.Println("Part 1")

	// printOctopuses(octopuses)
	increaseEnergy(&octopuses)

	printOctopuses(octopuses)

}

func part2(octopuses [][]octopus) {
	fmt.Println("Part 2")

	fmt.Println()
}

func main() {
	input := parseInput("test.txt")

	part1(input)
	part2(input)
}
