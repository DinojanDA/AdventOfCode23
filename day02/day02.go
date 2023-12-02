package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

var colors [3]string = [3]string{"red", "green", "blue"}

func gameNb(str string) (int, string) {
	var strTab []string = strings.Split(str, ": ")
	var nbStr []string = strings.Split(strTab[0], " ")
	nb, _ := strconv.Atoi(nbStr[1])
	return nb, strTab[1]
}

func TransformMap(str string) map[string]int {
	str = strings.Replace(str, "\r", "", -1)
	var myMap map[string]int = make(map[string]int)
	var results []string = strings.Split(str, ", ")
	for _, element := range results {
		var count []string = strings.Split(element, " ")
		myMap[count[1]], _ = strconv.Atoi(count[0])
	}
	return myMap
}

func possible(myMap map[string]int, redMax int, greenMax int, blueMax int) map[string]bool {
	maxMap := map[string]int{"red": redMax, "green": greenMax, "blue": blueMax}
	respect := map[string]bool{"red": true, "green": true, "blue": true}
	for _, color := range colors {
		_, ok := myMap[color]
		if ok {
			if myMap[color] > maxMap[color] {
				respect[color] = false
			}
		}
	}
	return respect
}

func possibleSet(set string, redMax int, greenMax int, blueMax int) bool {
	myMap := TransformMap(set)
	respect := possible(myMap, redMax, greenMax, blueMax)
	for _, color := range colors {
		if !respect[color] {
			return false
		}
	}
	return true
}

func possibleGame(game string, redMax int, greenMax int, blueMax int) bool {
	var sets []string = strings.Split(game, "; ")
	for _, set := range sets {
		if !possibleSet(set, redMax, greenMax, blueMax) {
			return false
		}
	}
	return true
}

func Power(game string) int {
	var pow int = 1
	myMap := map[string]int{"red": 0, "green": 0, "blue": 0}
	var sets []string = strings.Split(game, "; ")
	for _, set := range sets {
		var setMap map[string]int = TransformMap(set)
		for _, color := range colors {
			if setMap[color] > myMap[color] {
				myMap[color] = setMap[color]
			}
		}
	}
	for _, color := range colors {
		pow *= myMap[color]
	}
	return pow
}

func part1(str string) int {
	var sum int = 0
	var games []string = strings.Split(str, "\n")
	for _, game := range games {
		numGame, resultGame := gameNb(game)
		if possibleGame(resultGame, 12, 13, 14) {
			sum += numGame
		}
	}
	return sum
}

func part2(str string) int {
	var sum int = 0
	var games []string = strings.Split(str, "\n")
	for _, game := range games {
		var resultGame []string = strings.Split(game, ": ")
		var a int = Power(resultGame[1])
		sum += a
	}
	return sum
}

func main() {

	/*
		var test string = "Game 10: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		nb, resultGame := gameNb(test)
		fmt.Println(nb)
		fmt.Println(resultGame)

		fmt.Println(possibleGame(resultGame, 10, 12, 13))

		fmt.Println(possibleSet("3 blue, 4 red", 3, 12, 13))
	*/

	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
