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

func parseInput() []int {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	inputStrings := strings.Split(string(data), "\n")
	var inputInts []int

	for index := range inputStrings {
		value, err := strconv.Atoi(inputStrings[index])
		checkError(err)

		inputInts = append(inputInts, value)
	}

	return inputInts
}

func part1() {
	var depths []int = parseInput()
	var counter int = 0

	for i := range depths {

		if i != 0 {
			if depths[i] > depths[i-1] {
				counter = counter + 1
			}
		}
	}

	fmt.Println(counter)
}

func part2() {
	var depths []int = parseInput()
	var counter int = 0

	for i := range depths {

		if i > 2 {
			if depths[i]+depths[i-1]+depths[i-2] > depths[i-1]+depths[i-2]+depths[i-3] {
				counter = counter + 1
			}
		}
	}

	fmt.Println(counter)
}

func main() {
	part1()
	part2()
}
