package main

import (
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput() {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	// var lines []line

	// inputString := strings.Replace(string(data), " -> ", ",", -1)
	// inputStringSplit := strings.Split(inputString, "\n")

	// for x := range inputStringSplit {
	// 	lineString := strings.Split(inputStringSplit[x], ",")
	// 	x1, err := strconv.Atoi(lineString[0])
	// 	checkError(err)
	// 	y1, err := strconv.Atoi(lineString[1])
	// 	checkError(err)
	// 	x2, err := strconv.Atoi(lineString[2])
	// 	checkError(err)
	// 	y2, err := strconv.Atoi(lineString[3])
	// 	checkError(err)

	// 	line := line{coordinate1: coordinate{x: x1, y: y1}, coordinate2: coordinate{x: x2, y: y2}}
	// 	lines = append(lines, line)
	// }

	// return lines
}

func part1() {

}

func part2() {

}

func main() {
	part1()
	part2()
}
