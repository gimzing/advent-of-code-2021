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

	total := 0

	for i := range input {
		signal := make(map[int]string)
		signal[1] = ""
		signal[2] = ""
		signal[3] = ""
		signal[4] = ""
		signal[5] = ""
		signal[6] = ""
		signal[7] = ""

		// Segment 1
		var one string
		var seven string
		var four string
		var eight string

		for j := range input[i] {
			if len(input[i][j]) == 2 {
				one = input[i][j]
			} else if len(input[i][j]) == 3 {
				seven = input[i][j]
			} else if len(input[i][j]) == 4 {
				four = input[i][j]
			} else if len(input[i][j]) == 7 {
				eight = input[i][j]
			}
		}

		remainingSegment1 := seven
		for _, x := range one {
			remainingSegment1 = strings.Replace(remainingSegment1, string(x), "", -1)
		}
		signal[1] = remainingSegment1

		// Segment 6
		for j := range input[i] {
			if len(input[i][j]) == 6 {
				matchingSegmentCount := 0
				for _, k := range input[i][j] {
					if string(k) == string(one[0]) || string(k) == string(one[1]) {
						matchingSegmentCount++
					}
				}

				if matchingSegmentCount == 1 {
					remainingSegment3 := one
					for _, x := range input[i][j] {
						remainingSegment3 = strings.Replace(remainingSegment3, string(x), "", -1)
					}
					signal[3] = remainingSegment3
				}
			}
		}

		// Segment 3
		remainingSegment6 := one
		remainingSegment6 = strings.Replace(remainingSegment6, signal[3], "", -1)
		signal[6] = remainingSegment6

		// Segment 4
		remainingFour := four
		remainingFour = strings.Replace(remainingFour, signal[3], "", -1)
		remainingFour = strings.Replace(remainingFour, signal[6], "", -1)
		for j := range input[i] {
			if len(input[i][j]) == 6 {
				matchingSegmentCount := 0
				for _, k := range input[i][j] {
					if string(k) == string(remainingFour[0]) || string(k) == string(remainingFour[1]) {
						matchingSegmentCount++
					}
				}

				if matchingSegmentCount == 1 {
					remainingSegment4 := remainingFour
					for _, x := range input[i][j] {
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

		// Segment 5
		remainingEight := eight
		remainingEight = strings.Replace(remainingEight, signal[1], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[2], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[3], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[4], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[6], "", -1)
		for j := range input[i] {
			if len(input[i][j]) == 6 {
				matchingSegmentCount := 0
				for _, k := range input[i][j] {
					if string(k) == string(remainingEight[0]) || string(k) == string(remainingEight[1]) {
						matchingSegmentCount++
					}
				}
				if matchingSegmentCount == 1 {
					remainingSegment5 := remainingEight

					for _, x := range input[i][j] {

						remainingSegment5 = strings.Replace(remainingSegment5, string(x), "", -1)
					}
					signal[5] = remainingSegment5
				}
			}
		}

		// Segment 7
		remainingEight = eight
		remainingEight = strings.Replace(remainingEight, signal[1], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[2], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[3], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[4], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[5], "", -1)
		remainingEight = strings.Replace(remainingEight, signal[6], "", -1)
		signal[7] = remainingEight

		lastFour := input[i][10:]

		var result string
		for index := range lastFour {
			result += getDigit(lastFour[index], signal)
		}

		value, err := strconv.Atoi(result)
		checkError(err)

		total += value
	}

	fmt.Println(total)
	fmt.Println()
}

func main() {
	input1 := parseInputPart1("input.txt")
	input2 := parseInputPart2("input.txt")

	part1(input1)
	part2(input2)
}
