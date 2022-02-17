package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	magnitude int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(filename string) []string {
	data, err := os.ReadFile(filename)
	checkError(err)

	inputStrings := strings.Split(string(data), "\n")
	return inputStrings
}

func part1(input []string) {
	fmt.Println("Part 1")

	depths := input
	var forward int
	var depth int

	for i := range depths {
		fields := strings.Fields(depths[i])
		magnitude, err := strconv.Atoi(fields[1])
		checkError(err)

		command := command{direction: fields[0], magnitude: magnitude}

		if command.direction == "forward" {
			forward += command.magnitude
		} else if command.direction == "up" {
			depth -= command.magnitude
		} else if command.direction == "down" {
			depth += command.magnitude
		}
	}

	fmt.Println(depth, forward, depth*forward)
	fmt.Println()
}

func part2(input []string) {
	fmt.Println("Part 2")

	depths := input
	var forward int
	var depth int
	var aim int

	for i := range depths {
		fields := strings.Fields(depths[i])
		magnitude, err := strconv.Atoi(fields[1])
		checkError(err)
		command := command{direction: fields[0], magnitude: magnitude}

		if command.direction == "forward" {
			forward += command.magnitude
			depth += aim * command.magnitude
		} else if command.direction == "up" {
			aim -= command.magnitude
		} else if command.direction == "down" {
			aim += command.magnitude
		}
	}

	fmt.Println(depth, forward, depth*forward)
}

func main() {
	input := parseInput("input.txt")
	part1(input)
	part2(input)
}
