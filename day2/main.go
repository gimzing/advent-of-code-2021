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

func parseInput() []string {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	inputStrings := strings.Split(string(data), "\n")
	return inputStrings
}

func part1() {
	var depths []string = parseInput()

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
}

func part2() {
	var depths []string = parseInput()

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
	part1()
	part2()
}
