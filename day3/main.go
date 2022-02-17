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

func parseInput(filename string) []string {
	data, err := os.ReadFile(filename)
	checkError(err)

	inputStrings := strings.Split(string(data), "\n")

	return inputStrings
}

func binaryToDecimal(binary string) float64 {
	var result float64

	for i := len(binary) - 1; i > -1; i-- {
		bit, err := strconv.Atoi(string(binary[i]))
		checkError(err)
		result += float64(bit) * math.Pow(2, float64(len(binary)-i-1))
	}

	return result
}

func findOneAndZeroCounts(array []string, index int) (int, int) {
	oneCount := 0
	zeroCount := 0

	for i := range array {
		var bit string = string(array[i][index])
		if bit == "1" {
			oneCount++
		} else {
			zeroCount++
		}
	}

	return oneCount, zeroCount
}

func part1(input []string) {
	fmt.Println("Part 1")

	binaries := input
	length := len(binaries[0])
	var gamma_bits string = ""
	var epsilon_bits string = ""

	for i := 0; i < length; i++ {
		var oneCount int = 0
		var zeroCount int = 0
		for j := range binaries {
			var bit string = string(binaries[j][i])
			if bit == "1" {
				oneCount++
			} else {
				zeroCount++
			}
		}

		if oneCount >= zeroCount {
			gamma_bits += "1"
			epsilon_bits += "0"
		} else {
			gamma_bits += "0"
			epsilon_bits += "1"
		}
	}

	var gamma_rate float64 = binaryToDecimal(gamma_bits)
	var epsilon_rate float64 = binaryToDecimal(epsilon_bits)

	fmt.Println(gamma_rate * epsilon_rate)
	fmt.Println()
}

func part2(input []string) {
	fmt.Println("Part 2")

	binaries := input
	length := len(binaries[0])
	var currentArray []string = binaries
	i := 0

	for length > 1 {
		oneCount, zeroCount := findOneAndZeroCounts(currentArray, i)

		var newArray []string

		if oneCount >= zeroCount {
			for x := range currentArray {
				var bit string = string(currentArray[x][i])
				if bit == "1" {
					newArray = append(newArray, currentArray[x])
				}
			}
		} else {
			for x := range currentArray {
				var bit string = string(currentArray[x][i])
				if bit == "0" {
					newArray = append(newArray, currentArray[x])
				}
			}
		}

		currentArray = newArray
		length = len(currentArray)
		i++
	}

	var oxygen float64 = binaryToDecimal(currentArray[0])

	length = len(binaries[0])
	currentArray = binaries
	i = 0

	for length > 1 {
		oneCount, zeroCount := findOneAndZeroCounts(currentArray, i)
		var newArray []string

		if zeroCount <= oneCount {
			for x := range currentArray {
				var bit string = string(currentArray[x][i])
				if bit == "0" {
					newArray = append(newArray, currentArray[x])
				}
			}
		} else {
			for x := range currentArray {
				var bit string = string(currentArray[x][i])
				if bit == "1" {
					newArray = append(newArray, currentArray[x])
				}
			}
		}

		currentArray = newArray
		length = len(currentArray)
		i++
	}

	var c02 float64 = binaryToDecimal(currentArray[0])

	fmt.Println(oxygen * c02)
}

func main() {
	input := parseInput("input.txt")
	part1(input)
	part2(input)
}
