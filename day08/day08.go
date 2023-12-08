package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputday string

func giveKey(str string) string {
	key := strings.Split(str, " =")
	return key[0]
}

func giveLeft(str string) string {
	strTab := strings.Split(str, "(")
	strTab2 := strings.Split(strTab[1], ",")
	return strTab2[0]
}

func giveRight(str string) string {
	strTab := strings.Split(str, ", ")
	strTab2 := strings.Split(strTab[1], ")")
	return strTab2[0]
}

func part1(str string) int {
	var steps int = 0
	var currentStr string = "AAA"
	var strTab []string = strings.Split(str, "\r\n")
	var instructions []string = strings.Split(strTab[0], "")
	var instructionsIndex int = 0
	for currentStr != "ZZZ" {
		if instructionsIndex == len(instructions) {
			instructionsIndex = 0
		}
		var i int = 2
		for giveKey(strTab[i]) != currentStr {
			i++
		}
		if instructions[instructionsIndex] == "L" {
			currentStr = giveLeft(strTab[i])
			instructionsIndex++
		} else if instructions[instructionsIndex] == "R" {
			currentStr = giveRight(strTab[i])
			instructionsIndex++
		}
		steps++
	}
	return steps
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)
}
