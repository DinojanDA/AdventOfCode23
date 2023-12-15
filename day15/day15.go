package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputday string

func calculate(word string) int {
	var c int = 0
	for _, char := range word {
		c += int(char)
		c = c * 17
		c = c % 256
	}
	return c
}
func part1(str string) int {
	var words []string = strings.Split(str, ",")
	var sum int = 0
	for _, word := range words {
		sum += calculate(word)
	}
	return sum
}

func part2(str string) int {
	return 0
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
