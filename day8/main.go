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

func parseInput(filename string) [][]string {
	data, err := os.ReadFile(filename)
	checkError(err)

	var outputs [][]string

	linesString := strings.Split(string(data), "\n")

	for _, lineString := range linesString {
		splitLineString := strings.Split(lineString, " | ")
		outputString := splitLineString[1]
		output := strings.Split(outputString, " ")
		outputs = append(outputs, output)
	}

	return outputs
}

func part1(input [][]string) {
	fmt.Println("Part 1")

	count := 0
	for display := range input {
		for digit := range input[display] {
			segment := input[display][digit]

			if len(segment) == 2 ||
				len(segment) == 3 ||
				len(segment) == 4 ||
				len(segment) == 7 {
				count++
			}
		}
	}

	fmt.Println(count)
	fmt.Println()
}

func part2(input [][]string) {
	fmt.Println("Part 2")
}

func main() {
	input := parseInput("input.txt")

	part1(input)
	// part2(input)
}
