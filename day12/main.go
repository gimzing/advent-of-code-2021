package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type cave struct {
	name         string
	timesVisited int
	isBig        bool
	connections  []string
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

func makeCaves(input [][]string) map[string]cave {
	pathMap := make(map[string]cave)
	for _, path := range input {
		if _, ok := pathMap[path[0]]; !ok {
			pathMap[path[0]] = cave{name: path[0]}
		}
	}

	// Adding all connections to each node
	for _, path := range input {
		firstCoord := pathMap[path[0]]
		firstCoordConnections := firstCoord.connections
		firstCoordConnections = append(firstCoordConnections, path[1])
		firstCoord.connections = firstCoordConnections
		pathMap[path[0]] = firstCoord

		secondCoord := pathMap[path[1]]
		secondCoordConnections := secondCoord.connections
		secondCoordConnections = append(secondCoordConnections, path[0])
		secondCoord.connections = secondCoordConnections
		pathMap[path[1]] = secondCoord
	}

	// Don't need to look at 'end' since the path ends there
	delete(pathMap, "end")

	// Figuring out which caves are big
	for key := range pathMap {
		if unicode.IsUpper(rune(key[0])) {
			tempValue := pathMap[key]
			tempValue.isBig = true
			pathMap[key] = tempValue
		}
	}

	return pathMap
}

func part1(input [][]string) {
	fmt.Println("Part 1")

	var paths []string

	caves := makeCaves(input)
	fmt.Println(caves)

	hasChoices := true
	currentCave := caves["start"]
	currentPath := ""
	for hasChoices {
		currentPath += currentCave.name + ","
		currentChoices := currentCave.connections
		fmt.Println(currentChoices)

		paths = append(paths, currentPath)
		hasChoices = false
	}

	fmt.Println(paths)

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
