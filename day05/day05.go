package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

func findMin(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func transformSet(set map[int]int, str string) {
	str = strings.Replace(str, "\r", "", -1)
	var s []string = strings.Split(str, " ")
	key, _ := strconv.Atoi(s[1])
	value, _ := strconv.Atoi(s[0])
	len, _ := strconv.Atoi(s[2])
	var i int = 0
	for i < len {
		set[key+i] = value + i
		i++
	}
}

func returnCorrespondance(str string, i int) int {
	lignes := strings.Split(str, "\r\n")
	for j, ligne := range lignes {
		if j != 0 {
			var s []string = strings.Split(ligne, " ")
			entry, _ := strconv.Atoi(s[1])
			exit, _ := strconv.Atoi(s[0])
			len, _ := strconv.Atoi(s[2])
			if entry <= i && i < entry+len {
				return exit - entry + i
			}
		}
	}
	return i
}

func part1(str string) int {
	var categories []string = strings.Split(str, "\r\n\r\n")
	var s []string = strings.Split(categories[0], ": ")
	var seeds []string = strings.Split(s[1], " ")
	var locationTab []int
	for _, seed_ := range seeds {
		seed, _ := strconv.Atoi(seed_)
		var soil int = returnCorrespondance(categories[1], seed)
		var fertilizer int = returnCorrespondance(categories[2], soil)
		var water int = returnCorrespondance(categories[3], fertilizer)
		var light int = returnCorrespondance(categories[4], water)
		var temperature int = returnCorrespondance(categories[5], light)
		var humidity int = returnCorrespondance(categories[6], temperature)
		var location int = returnCorrespondance(categories[7], humidity)
		locationTab = append(locationTab, location)
	}
	return findMin(locationTab)
}

func part2(str string) int {
	var categories []string = strings.Split(str, "\r\n\r\n")
	var s []string = strings.Split(categories[0], ": ")
	var seeds []string = strings.Split(s[1], " ")
	var locationTab []int
	for i := 0; i < len(seeds)-1; i += 2 {
		seedi, _ := strconv.Atoi(seeds[i])
		seedi1, _ := strconv.Atoi(seeds[i+1])
		for j := 0; j < seedi1; j++ {
			var soil int = returnCorrespondance(categories[1], seedi+j)
			var fertilizer int = returnCorrespondance(categories[2], soil)
			var water int = returnCorrespondance(categories[3], fertilizer)
			var light int = returnCorrespondance(categories[4], water)
			var temperature int = returnCorrespondance(categories[5], light)
			var humidity int = returnCorrespondance(categories[6], temperature)
			var location int = returnCorrespondance(categories[7], humidity)
			locationTab = append(locationTab, location)
		}
	}
	return findMin(locationTab)
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
