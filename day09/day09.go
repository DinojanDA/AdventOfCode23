package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

func convertToTabInt(str string) []int {
	numbersStr := strings.Fields(str)
	var numbers []int
	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}
func nonZeros(history []int) bool {
	for _, value := range history {
		if value != 0 {
			return true
		}
	}
	return false
}

func difference(history []int) []int {
	diffs := make([]int, 0, len(history)-1)
	for i := 0; i < len(history)-1; i++ {
		diffs = append(diffs, history[i+1]-history[i])
	}
	return diffs
}

func part1(str string) int {
	var sum int = 0
	var histories []string = strings.Split(str, "\r\n")
	for _, strHistory := range histories {
		var history []int = convertToTabInt(strHistory)
		var stockNb int
		for nonZeros(history) {
			stockNb += history[len(history)-1]
			history = difference(history)
		}
		sum += stockNb
	}
	return sum
}
func part2(str string) int {
	var sum int = 0
	var histories []string = strings.Split(str, "\r\n")
	for _, strHistory := range histories {
		var history []int = convertToTabInt(strHistory)
		var stockNb int
		var signe int = 1
		for nonZeros(history) {
			stockNb += signe * history[0]
			signe = -1 * signe
			history = difference(history)
		}
		sum += stockNb
	}
	return sum
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
