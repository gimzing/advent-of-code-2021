package main

import (
	"fmt"
	"math"
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

func part1(input []int) {

	var lowestFuel int

	for x := range input {
		total := 0
		for y := range input {
			total += int(math.Abs(float64(x) - float64(input[y])))
		}
		fmt.Println("total", total)

		if lowestFuel == 0 {
			lowestFuel = total
		} else {
			if total < lowestFuel {
				lowestFuel = total
			}
		}
	}

	fmt.Println("lowestFuel", lowestFuel)

}

// func part2() {

// }

func main() {
	input := parseInput("input.txt")
	part1(input)
	// part2()
}
