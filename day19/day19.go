package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

func convertMap(workspaces []string) map[string]string {
	results := make(map[string]string)
	for _, workspace := range workspaces {
		strTab := strings.Split(workspace, "{")
		key := strTab[0]
		value := strings.Replace(strTab[1], "}", "", 1)
		results[key] = value
	}
	return results
}

func convertParts(parts string) map[string]int {
	results := make(map[string]int)
	parts = strings.Replace(parts, "}", "", 1)
	parts = strings.Replace(parts, "{", "", 1)
	partsTab := strings.Split(parts, ",")
	for _, part := range partsTab {
		partTab := strings.Split(part, "=")
		key := partTab[0]
		value, _ := strconv.Atoi(partTab[1])
		results[key] = value
	}
	return results
}
func isAccepted(part map[string]int, workflows map[string]string) bool {
	var name string = "in"
	var rules string
	for name != "A" && name != "R" {
		rules = workflows[name]
		name = newRule(part, rules)
	}
	if name == "A" {
		return true
	} else {
		return false
	}
}

func newRule(partMap map[string]int, rules string) string {
	rulesTab := strings.Split(rules, ",")

	for i, rule := range rulesTab {
		if i == len(rulesTab)-1 {
			return rulesTab[i]
		}
		condition := strings.Split(rule, ":")[0]
		if respectCondition(condition, partMap) {
			return strings.Split(rule, ":")[1]
		}
	}
	return ""
}

func respectCondition(condition string, partMap map[string]int) bool {
	if strings.Contains(condition, "<") {
		strTab := strings.Split(condition, "<")
		key := strTab[0]
		value, _ := strconv.Atoi(strTab[1])
		if partMap[key] < value {
			return true
		} else {
			return false
		}
	} else {
		strTab := strings.Split(condition, ">")
		key := strTab[0]
		value, _ := strconv.Atoi(strTab[1])
		if partMap[key] > value {
			return true
		} else {
			return false
		}
	}
}
func part1(str string) int {
	puzzle := strings.Split(str, "\r\n\r\n")
	workflows := convertMap(strings.Split(puzzle[0], "\r\n"))
	parts := strings.Split(puzzle[1], "\r\n")
	var accepted int = 0
	for _, part := range parts {
		partMap := convertParts(part)
		if isAccepted(partMap, workflows) {
			accepted += partMap["x"] + partMap["m"] + partMap["a"] + partMap["s"]
		}
	}
	return accepted
}

func part2(str string) int {
	puzzle := strings.Split(str, "\r\n\r\n")
	workflows := convertMap(strings.Split(puzzle[0], "\r\n"))
	partMap := make(map[string]int)
	var accepted int = 0
	for x := 0; x <= 4000; x++ {
		for m := 0; m <= 4000; m++ {
			for a := 0; a <= 4000; a++ {
				for s := 0; s <= 4000; s++ {
					partMap["x"] = x
					partMap["m"] = m
					partMap["a"] = a
					partMap["s"] = s
					if isAccepted(partMap, workflows) {
						accepted += 1
					}
				}
			}
		}
	}
	return accepted
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)
	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
