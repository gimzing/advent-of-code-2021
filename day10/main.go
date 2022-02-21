package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var CHAR_MAP = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var ERROR_SCORES = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var COMPLETION_SCORES = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

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

func checkMatchingCharacter(buffer *[]string, errors *[]string, character string) bool {
	hadError := false
	if (*buffer)[len((*buffer))-1] != character {
		*errors = append((*errors), CHAR_MAP[character])
		hadError = true
	} else {
		*buffer = (*buffer)[:len((*buffer))-1]
	}
	return hadError
}

func part1(lines []string) []string {
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
				hadError = checkMatchingCharacter(&currentBuffer, &errors, "(")
			case "]":
				hadError = checkMatchingCharacter(&currentBuffer, &errors, "[")
			case "}":
				hadError = checkMatchingCharacter(&currentBuffer, &errors, "{")
			case ">":
				hadError = checkMatchingCharacter(&currentBuffer, &errors, "<")
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
		points += ERROR_SCORES[character]
	}

	fmt.Println(points)
	fmt.Println()

	return incompleteLines
}

func part2(lines []string) {
	fmt.Println("Part 2")

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
			currentPoints *= 5
			currentPoints += COMPLETION_SCORES[currentBuffer[i]]
		}
		points = append(points, currentPoints)
	}

	sort.Ints(points)
	fmt.Println(points[len(points)/2])
	fmt.Println()
}

func main() {
	input := parseInput("input.txt")

	incompleteLines := part1(input)
	part2(incompleteLines)
}
