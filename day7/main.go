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

		if lowestFuel == 0 {
			lowestFuel = total
		} else {
			if total < lowestFuel {
				lowestFuel = total
			}
		}
	}

	fmt.Println(lowestFuel)
	fmt.Println()
}

func part2(input []int) {

	var lowestFuel int

	for x := range input {
		total := 0

		for y := range input {
			steps := int(math.Abs(float64(x) - float64(input[y])))

			// Triangular numbers formula
			total += steps * (steps + 1) / 2

			// Or the manual way
			// fuelTotal := 0
			// for i := 1; i < steps+1; i++ {
			// 	fuelTotal += i
			// }
			// total += fuelTotal
		}

		if lowestFuel == 0 {
			lowestFuel = total
		} else {
			if total < lowestFuel {
				lowestFuel = total
			}
		}
	}

	fmt.Println(lowestFuel)
}

func main() {
	input := parseInput("input.txt")
	part1(input)
	part2(input)
}
