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

func parseInput(filename string) []int {
	data, err := os.ReadFile(filename)
	checkError(err)

	var numbers []int

	numbersStrings := strings.Split(string(data), ",")
	for _, x := range numbersStrings {
		numberInt, err := strconv.Atoi(x)
		checkError(err)
		numbers = append(numbers, numberInt)
	}

	return numbers
}

func part1(input []int, days int) {
	fmt.Println("Part 1")
	timers := input

	day := 0
	for day < days {
		var newTimers []int

		for _, timer := range timers {
			if timer == 0 {
				newTimers = append(newTimers, 6)
				newTimers = append(newTimers, 8)
			} else {
				newTimers = append(newTimers, timer-1)
			}
		}

		timers = newTimers

		day++
	}
	fmt.Println(len(timers))
	fmt.Println()
}

func part2(input []int, days int) {
	fmt.Println("Part 2")

	timers := input

	m := make(map[int]int)

	for i := 0; i < 9; i++ {
		m[i] = 0
	}

	for _, timer := range timers {
		m[timer]++
	}

	day := 0
	for day < days {
		newMap := make(map[int]int)
		for i := 0; i < 9; i++ {
			newMap[i] = 0
		}

		newMap[7] = m[8]
		newMap[6] = m[7]
		newMap[5] = m[6]
		newMap[4] = m[5]
		newMap[3] = m[4]
		newMap[2] = m[3]
		newMap[1] = m[2]
		newMap[0] = m[1]

		newMap[6] += m[0]
		newMap[8] += m[0]

		m = newMap
		day++
	}

	count := 0
	for _, value := range m {
		count += value
	}

	fmt.Println(count)

}

func main() {
	input := parseInput("input.txt")

	part1(input, 80)
	part2(input, 256)
}
