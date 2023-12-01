package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

func convert(str string) []string {
	var s []string
	for _, char := range str {
		s = append(s, string(char))
	}
	return s
}

func find(str string) int {
	var tab []string = convert(str)
	var compteur int = 0
	var compteurInv int = len(str) - 1
	var a string
	var b string

	_, err1 := strconv.Atoi(tab[compteur])
	for err1 != nil {
		compteur++
		_, err1 = strconv.Atoi(tab[compteur])
	}
	a = tab[compteur]

	_, err2 := strconv.Atoi(tab[compteurInv])
	for err2 != nil {
		compteurInv--
		_, err2 = strconv.Atoi(tab[compteurInv])
	}
	b = tab[compteurInv]

	var strInt string = a + b
	sol, _ := strconv.Atoi(strInt)
	return sol
}

func part1(str string) int {
	var sum int = 0
	var strTab []string = strings.Split(str, "\n")
	for _, p := range strTab {
		sum += find(p)
	}
	return sum
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)
}
