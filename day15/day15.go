package main

import (
	_ "embed"
	"fmt"
	"strconv"
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

func parseInstruction(instruction string) (string, string, int) {
	var label string
	var operation string
	var value int
	if strings.Contains(instruction, "=") {
		parts := strings.Split(instruction, "=")
		label = parts[0]
		operation = "="
		value, _ = strconv.Atoi(parts[1])
	} else {
		label = strings.TrimSuffix(instruction, "-")
		operation = "-"
		value = -1
	}
	return label, operation, value
}

func isPresent(matrix [256][]string, st string) int {
	var i int = calculate(st)
	var j int = 0
	for j < len(matrix[i]) {
		s := strings.Split(matrix[i][j], " ")
		if s[0] == st {
			return j
		}
		j++
	}
	return -1
}

func remove(matrix *[256][]string, st string) {
	i := calculate(st)
	j := isPresent(*matrix, st)
	if j != -1 {
		matrix[i] = append(matrix[i][:j], matrix[i][j+1:]...)
	}
}

func calculateMatrix(matrix [256][]string) int {
	var sum int
	for i, row := range matrix {
		for j, cell := range row {
			if cell != "" {
				parts := strings.Split(cell, " ")
				value, _ := strconv.Atoi(parts[1])
				sum += (i + 1) * (j + 1) * value
			}
		}
	}
	return sum
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
	var boxes [256][]string
	var words []string = strings.Split(str, ",")
	for _, word := range words {
		st, op, a := parseInstruction(word)
		c := calculate(st)
		if op == "=" {
			var l int = isPresent(boxes, st)
			if l != -1 {
				s := strconv.Itoa(a)
				boxes[c][l] = st + " " + s
			} else {
				s := strconv.Itoa(a)
				boxes[c] = append(boxes[c], st+" "+s)
			}
		} else {
			var l int = isPresent(boxes, st)
			if l != -1 {
				remove(&boxes, st)
			}
		}
	}
	return calculateMatrix(boxes)
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
