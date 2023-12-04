package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"unicode"
)

//go:embed input.txt
var inputday string

func betterString(str string) string {
	str = strings.TrimLeftFunc(str, unicode.IsSpace)
	return strings.Replace(str, "\r", "", -1)
}

func search(winningCards []string, myCards []string) []string {
	var cards []string
	for _, myCard := range myCards {
		for _, winningCard := range winningCards {
			if myCard == winningCard {
				cards = append(cards, myCard)
			}
		}
	}
	return cards
}

func points(cards []string) int {
	var l float64 = float64(len(cards))
	return int(math.Pow(2, l-1))
}

func part1(str string) int {
	var sum int = 0
	var games []string = strings.Split(str, "\n")
	for _, game := range games {
		var cards []string = strings.Split(game, ": ")
		cards[1] = strings.ReplaceAll(cards[1], "  ", " ")
		var hands []string = strings.Split(cards[1], " | ")
		var winningCards []string = strings.Split(betterString(hands[0]), " ")
		var myCards []string = strings.Split(betterString(hands[1]), " ")
		var point int = points(search(winningCards, myCards))
		sum += point
	}
	return sum
}

func part2(str string) int {
	return 0
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

}
