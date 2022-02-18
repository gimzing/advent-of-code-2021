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

func parseInputPart1(filename string) [][]string {
	data, err := os.ReadFile(filename)
	checkError(err)

	var outputs [][]string

	linesString := strings.Split(string(data), "\n")

	for _, lineString := range linesString {
		splitLineString := strings.Split(lineString, " | ")
		outputString := splitLineString[1]
		output := strings.Split(outputString, " ")
		outputs = append(outputs, output)
	}

	return outputs
}

func parseInputPart2(filename string) [][]string {
	data, err := os.ReadFile(filename)
	checkError(err)

	var outputs [][]string

	linesString := strings.Split(string(data), "\n")

	for _, lineString := range linesString {
		splitLineString := strings.Replace(lineString, " | ", " ", -1)
		output := strings.Split(splitLineString, " ")

		outputs = append(outputs, output)
	}

	return outputs
}

func getDigit(entry string, signal map[int]string) string {
	if len(entry) == 2 {
		return "1"
	}

	if len(entry) == 3 {
		return "7"
	}

	if len(entry) == 4 {
		return "4"
	}

	if len(entry) == 7 {
		return "8"
	}

	// TODO: WRONG HERE
	if len(entry) == 5 {
		// 2 3 5
		counter := 0
		for i := range entry {
			if string(entry[i]) == signal[3] ||
				string(entry[i]) == signal[6] {
				counter++
			}
		}

		if counter == 2 {
			return "3"
		} else {
			for i := range entry {
				if string(entry[i]) == signal[2] {
					return "5"
				}

				if string(entry[i]) == signal[5] {
					return "2"
				}
			}
		}
	}

	if len(entry) == 6 {
		// 0 6 9
		hasMiddle := false
		for i := range entry {
			if string(entry[i]) == signal[4] {
				hasMiddle = true
			}
		}

		if !hasMiddle {
			return "0"
		} else {
			for i := range entry {
				if string(entry[i]) == signal[3] {
					return "9"
				}

				if string(entry[i]) == signal[5] {
					return "6"
				}
			}
		}
	}

	return entry
}

func part1(input [][]string) {
	fmt.Println("Part 1")

	count := 0
	for display := range input {
		for digit := range input[display] {
			segment := input[display][digit]

			if len(segment) == 2 ||
				len(segment) == 3 ||
				len(segment) == 4 ||
				len(segment) == 7 {
				count++
			}
		}
	}

	fmt.Println(count)
	fmt.Println()
}

func part2(input [][]string) {
	fmt.Println("Part 2")

	signal := make(map[int]string)
	signal[1] = ""
	signal[2] = ""
	signal[3] = ""
	signal[4] = ""
	signal[5] = ""
	signal[6] = ""
	signal[7] = ""

	//	 1111
	//	2    3
	//	2    3
	//	 4444
	//	5    6
	//	5    6
	//	 7777

	// 0: 6
	// 1: 2
	// 2: 5
	// 3: 5
	// 4: 4
	// 5: 5
	// 6: 6
	// 7: 3
	// 8: 7
	// 9: 6

	// Segment 1
	var one string
	var seven string
	var four string
	var eight string

	for j := range input[0] {
		if len(input[0][j]) == 2 {
			one = input[0][j]
		} else if len(input[0][j]) == 3 {
			seven = input[0][j]
		} else if len(input[0][j]) == 4 {
			four = input[0][j]
		} else if len(input[0][j]) == 7 {
			eight = input[0][j]
		}
	}

	remainingSegment1 := seven
	for _, x := range one {
		remainingSegment1 = strings.Replace(remainingSegment1, string(x), "", -1)
	}
	signal[1] = remainingSegment1

	// Segment 6
	// var six string
	for j := range input[0] {
		if len(input[0][j]) == 6 {
			matchingSegmentCount := 0
			for _, k := range input[0][j] {
				if string(k) == string(one[0]) || string(k) == string(one[1]) {
					matchingSegmentCount++
				}
			}

			if matchingSegmentCount == 1 {
				// six = input[0][j]
				remainingSegment6 := one
				for _, x := range input[0][j] {
					remainingSegment6 = strings.Replace(remainingSegment6, string(x), "", -1)
				}
				signal[6] = remainingSegment6
			}
		}
	}

	// Segment 3
	remainingSegment3 := one
	remainingSegment3 = strings.Replace(remainingSegment3, signal[6], "", -1)
	signal[3] = remainingSegment3

	// Segment 4
	remainingFour := four
	remainingFour = strings.Replace(remainingFour, signal[3], "", -1)
	remainingFour = strings.Replace(remainingFour, signal[6], "", -1)
	for j := range input[0] {
		if len(input[0][j]) == 6 {
			matchingSegmentCount := 0
			for _, k := range input[0][j] {
				if string(k) == string(remainingFour[0]) || string(k) == string(remainingFour[1]) {
					matchingSegmentCount++
				}
			}

			if matchingSegmentCount == 1 {
				remainingSegment4 := remainingFour
				for _, x := range input[0][j] {
					remainingSegment4 = strings.Replace(remainingSegment4, string(x), "", -1)
				}
				signal[4] = remainingSegment4
			}
		}
	}

	// Segment 2
	remainingSegment2 := four
	remainingSegment2 = strings.Replace(remainingSegment2, signal[3], "", -1)
	remainingSegment2 = strings.Replace(remainingSegment2, signal[4], "", -1)
	remainingSegment2 = strings.Replace(remainingSegment2, signal[6], "", -1)
	signal[2] = remainingSegment2

	// Segment 7
	remainingEight := eight
	remainingEight = strings.Replace(remainingEight, signal[1], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[2], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[3], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[4], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[6], "", -1)
	for j := range input[0] {
		if len(input[0][j]) == 6 {
			matchingSegmentCount := 0
			for _, k := range input[0][j] {
				if string(k) == string(remainingEight[0]) || string(k) == string(remainingEight[1]) {
					matchingSegmentCount++
				}
			}

			if matchingSegmentCount == 1 {
				remainingSegment7 := remainingEight
				for _, x := range input[0][j] {
					remainingSegment7 = strings.Replace(remainingSegment7, string(x), "", -1)
				}
				signal[7] = remainingSegment7
			}
		}
	}

	// Segment 5
	remainingEight = eight
	remainingEight = strings.Replace(remainingEight, signal[1], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[2], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[3], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[4], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[6], "", -1)
	remainingEight = strings.Replace(remainingEight, signal[7], "", -1)
	signal[5] = remainingEight

	lastFour := input[0][len(input[0])-4:]

	var result string
	for i := range lastFour {
		result += getDigit(lastFour[i], signal)
	}
	value, err := strconv.Atoi(result)
	checkError(err)

	fmt.Println(value, "8394")
}

func main() {
	input1 := parseInputPart1("test.txt")
	input2 := parseInputPart2("test.txt")

	part1(input1)
	part2(input2)
}
