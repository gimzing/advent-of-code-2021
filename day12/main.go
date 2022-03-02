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

type cave struct {
	name        string
	connections []cave
}

func parseInput(filename string) [][]string {
	data, err := os.ReadFile(filename)
	checkError(err)

	var paths [][]string

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		splitLine := strings.Split(line, "-")
		paths = append(paths, splitLine)
	}

	return paths
}

func makeCaves(input [][]string) map[string][]string {

	pathMap := make(map[string][]string)
	for _, path := range input {
		if _, ok := pathMap[path[0]]; !ok {
			pathMap[path[0]] = nil
		}
	}

	for _, path := range input {
		pathMap[path[0]] = append(pathMap[path[0]], path[1])
		pathMap[path[1]] = append(pathMap[path[1]], path[0])
	}

	return pathMap
}

func part1(input [][]string) {
	fmt.Println("Part 1")

	pathMap := makeCaves(input)
	fmt.Println(pathMap)

	var results []string

	endReached := false
	for !endReached {

		for _, i := range pathMap["start"] {
			fmt.Println(i)
			for _, j := range i {
				fmt.Println(pathMap[string(j)])
			}
		}

		endReached = true
	}

	fmt.Println(results)
	fmt.Println()
}

func part2(input [][]string) {
	fmt.Println("Part 2")

	fmt.Println()
}

func main() {
	input := parseInput("test1.txt")

	part1(input)
	// part2(input)
}
