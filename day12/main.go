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

func parseInput(filename string) []string {
	data, err := os.ReadFile(filename)
	checkError(err)

	lines := strings.Split(string(data), "\n")
	return lines
}

func part1(input []string) {
	fmt.Println("Part 1")

	fmt.Println()
}

func part2(input []string) {
	fmt.Println("Part 2")

	fmt.Println()
}

func main() {
	input := parseInput("test1.txt")

	part1(input)
	part2(input)
}
