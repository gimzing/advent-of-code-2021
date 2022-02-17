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

func parseInput(filename string) []int {
	data, err := os.ReadFile(filename)
	checkError(err)

	var numbers []int

	numbersStrings := strings.Split(string(data), ",")
	for _, x := range numbersStrings {
		numberInt, err := strconv.Atoi(x)
		checkError(err)
		numbers = append(numbers, numberInt)
	}

	return numbers
}

func part1(input) {
	fmt.Println("Part 1")

}

func part2(input) {
	fmt.Println("Part 2")

}

func main() {
	input := parseInput("test.txt")

	part1(input)
	part2(input)
}
