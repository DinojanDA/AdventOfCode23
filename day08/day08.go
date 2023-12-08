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

func findFinalA(strTab []string) []string {
	var tabFinalA []string
	for i := 2; i < len(strTab); i++ {
		key := giveKey(strTab[i])
		var lastLetter string = strings.Split(key, "")[2]
		if lastLetter == "A" {
			tabFinalA = append(tabFinalA, key)
		}
	}
	return tabFinalA
}

func finalZ(strTab []string) bool {
	for _, str := range strTab {
		if !strings.HasSuffix(str, "Z") {
			return false
		}
	}
	return true
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

func part2(str string) int {
	var steps int = 0
	var strTab []string = strings.Split(str, "\r\n")
	var entries []string = findFinalA(strTab)
	var instructions []string = strings.Split(strTab[0], "")
	var instructionsIndex int = 0
	for !finalZ(entries) {
		if instructionsIndex == len(instructions) {
			instructionsIndex = 0
		}
		for j := 0; j < len(entries); j++ {
			var i int = 2
			for giveKey(strTab[i]) != entries[j] {
				i++
			}
			if instructions[instructionsIndex] == "L" {
				entries[j] = giveLeft(strTab[i])
			} else if instructions[instructionsIndex] == "R" {
				entries[j] = giveRight(strTab[i])
			}
		}
		instructionsIndex++
		steps++
	}
	return steps
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
