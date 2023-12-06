package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	maxi := numbers[0]
	for _, num := range numbers {
		if num > maxi {
			maxi = num
		}
	}
	return maxi
}
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

func arrive(hold int, time int, distance int) bool {
	if hold > time {
		return false
	}
	return hold*(time-hold) > distance
}

func multTab(tab []int) int {
	var len int = len(tab)
	var sum int = 1
	var i int = 0
	for i < len {
		sum *= tab[i]
		i++
	}
	return sum
}

func part1(str string) int {
	var categories []string = strings.Split(str, "\r\n")

	var strTimes string = strings.Split(categories[0], ":")[1]
	var strDistances string = strings.Split(categories[1], ":")[1]
	var times []int = convertToTabInt(strTimes)
	var distances []int = convertToTabInt(strDistances)
	var winWays []int = make([]int, len(times))
	maxTime := findMax(times)

	for i := 0; i < maxTime+1; i++ {
		for j := 0; j < len(times); j++ {
			if arrive(i, times[j], distances[j]) {
				winWays[j] += 1
			}
		}
	}
	return multTab(winWays)
}

func part2(str string) int {
	var categories []string = strings.Split(str, "\r\n")

	var strTimes string = strings.Split(categories[0], ":")[1]
	var strDistances string = strings.Split(categories[1], ":")[1]
	var strTime string = strings.ReplaceAll(strTimes, " ", "")
	var strDistance string = strings.ReplaceAll(strDistances, " ", "")
	time, _ := strconv.Atoi(strTime)
	distance, _ := strconv.Atoi(strDistance)
	var winWays int = 0

	for i := 0; i < time; i++ {
		if arrive(i, time, distance) {
			winWays += 1
		}
	}
	return winWays
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
