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

func part1(lines []string) {
	fmt.Println("Part 1")

	var errors []string
	for _, line := range lines {
		fmt.Println(line)
		var currentBuffer []string
		for _, char := range line {
			hadError := false
			i := string(char)
			switch i {
			case "(":
				currentBuffer = append(currentBuffer, i)
			case "[":
				currentBuffer = append(currentBuffer, i)
			case "{":
				currentBuffer = append(currentBuffer, i)
			case "<":
				currentBuffer = append(currentBuffer, i)

			case ")":
				if currentBuffer[len(currentBuffer)-1] != "(" {
					errors = append(errors, ")")
					hadError = true
				} else {
					currentBuffer = currentBuffer[:len(currentBuffer)-1]
				}
			case "]":
				if currentBuffer[len(currentBuffer)-1] != "[" {
					errors = append(errors, "]")
					hadError = true
				} else {
					currentBuffer = currentBuffer[:len(currentBuffer)-1]
				}
			case "}":
				if currentBuffer[len(currentBuffer)-1] != "{" {
					errors = append(errors, "}")
					hadError = true
				} else {
					currentBuffer = currentBuffer[:len(currentBuffer)-1]
				}
			case ">":
				if currentBuffer[len(currentBuffer)-1] != "<" {
					errors = append(errors, ">")
					hadError = true
				} else {
					currentBuffer = currentBuffer[:len(currentBuffer)-1]
				}
			}

			if hadError {
				break
			}
		}
	}
	fmt.Println(errors)

	points := 0
	for _, character := range errors {
		if character == ")" {
			points += 3
		} else if character == "]" {
			points += 57
		} else if character == "}" {
			points += 1197
		} else if character == ">" {
			points += 25137
		}
	}

	fmt.Println(points)
	fmt.Println()
}

func part2(lines []string) {
	fmt.Println("Part 2")

}

func main() {
	input := parseInput("test.txt")

	part1(input)
	part2(input)
}
