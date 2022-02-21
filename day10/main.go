package main

import (
	"fmt"
	"os"
	"sort"
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

func checkMatchingCharacter(buffer *[]string, errors *[]string, charmap map[string]string, character string) bool {
	hadError := false
	if (*buffer)[len((*buffer))-1] != character {
		*errors = append((*errors), charmap[character])
		hadError = true
	} else {
		*buffer = (*buffer)[:len((*buffer))-1]
	}
	return hadError
}

func part1(lines []string, charmap map[string]string) []string {
	fmt.Println("Part 1")

	var incompleteLines []string
	var errors []string
	for _, line := range lines {
		var currentBuffer []string
		hadError := false
		for _, char := range line {
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
				hadError = checkMatchingCharacter(&currentBuffer, &errors, charmap, "(")
			case "]":
				hadError = checkMatchingCharacter(&currentBuffer, &errors, charmap, "[")
			case "}":
				hadError = checkMatchingCharacter(&currentBuffer, &errors, charmap, "{")
			case ">":
				hadError = checkMatchingCharacter(&currentBuffer, &errors, charmap, "<")
			}

			if hadError {
				break
			}
		}

		if !hadError {
			incompleteLines = append(incompleteLines, line)
		}
	}

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

	return incompleteLines
}

func part2(lines []string, charmap map[string]string) {
	fmt.Println("Part 2")

	// fmt.Println(lines)

	var points []int
	for _, line := range lines {
		var currentBuffer []string
		for _, char := range line {
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
				currentBuffer = currentBuffer[:len(currentBuffer)-1]
			case "]":
				currentBuffer = currentBuffer[:len(currentBuffer)-1]
			case "}":
				currentBuffer = currentBuffer[:len(currentBuffer)-1]
			case ">":
				currentBuffer = currentBuffer[:len(currentBuffer)-1]
			}
		}

		currentPoints := 0
		for i := len(currentBuffer) - 1; i >= 0; i-- {
			currentPoints = currentPoints * 5
			if currentBuffer[i] == "(" {
				currentPoints += 1
			} else if currentBuffer[i] == "[" {
				currentPoints += 2
			} else if currentBuffer[i] == "{" {
				currentPoints += 3
			} else if currentBuffer[i] == "<" {
				currentPoints += 4
			}
		}
		points = append(points, currentPoints)
	}

	sort.Ints(points)
	fmt.Println(points[len(points)/2])
	fmt.Println()
}

func main() {
	input := parseInput("input.txt")

	charmap := make(map[string]string)
	charmap["("] = ")"
	charmap["["] = "]"
	charmap["{"] = "}"
	charmap["<"] = ">"

	incompleteLines := part1(input, charmap)
	part2(incompleteLines, charmap)
}
