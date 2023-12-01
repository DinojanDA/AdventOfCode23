package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

//go:embed input2.txt
var inputday2 string

var nbTab [10]string = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var nbInvTab [10]string = [10]string{"orez", "eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"} // utile pour gérer le cas "one2one"

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

func reverseString(s string) string {
	runes := []rune(s)
	var reversedRunes []rune

	for i := len(runes) - 1; i >= 0; i-- {
		reversedRunes = append(reversedRunes, runes[i])
	}

	return string(reversedRunes)
}

func present(str string) [10]int {
	var indexPresent [10]int
	for i, number := range nbTab {
		index := strings.Index(str, number)
		indexPresent[i] = index
	}
	return indexPresent
}

func invPresent(str string) [10]int {
	var indexPresent [10]int
	for i, number := range nbInvTab {
		index := strings.Index(str, number)
		indexPresent[i] = index
	}
	return indexPresent
}

func indexMinMax(str string) (int, int, int, int) {
	numbers := present(str)
	invNumbers := invPresent(reverseString(str))
	min := len(str)
	max := len(str)
	l := 10
	m := 10

	for i, value := range numbers {
		if value < min && value != -1 {
			min = value
			l = i
		}
	}

	for j, value := range invNumbers {
		if value < max && value != -1 {
			max = value
			m = j
		}
	}
	return min, l, len(str) - max - 1, m
}

func findV2(str string) int {
	var tab []string = convert(str)
	var indexMin, min, indexMax, max = indexMinMax(str)
	var a string = strconv.Itoa(min)
	var b string = strconv.Itoa(max)

	for i := 0; i < indexMin; i++ {
		if _, err1 := strconv.Atoi(tab[i]); err1 == nil {
			a = tab[i]
			break
		}
	}

	for i := len(str) - 1; i > indexMax; i-- {
		if _, err2 := strconv.Atoi(tab[i]); err2 == nil {
			b = tab[i]
			break
		}
	}

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

func part2(str string) int {
	var sum int = 0
	var strTab []string = strings.Split(str, "\n")
	for _, p := range strTab {
		sum += findV2(p)
	}
	return sum
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday2)
	fmt.Println(sol2)
}
