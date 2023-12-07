package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	_ "strings"
)

//go:embed input.txt
var inputday string

var cardValues map[string]int = map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}
var cardValuesWithJoker map[string]int = map[string]int{"A": 14, "K": 13, "Q": 12, "J": 1, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}

func countRepetitions(s string) []int {
	counts := make(map[string]int)
	for _, char := range s {
		counts[string(char)]++
	}

	var repetitions []int
	for _, count := range counts {
		repetitions = append(repetitions, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(repetitions)))
	return repetitions
}

func countRepetitionsWithJoker(s string) []int {
	counts := make(map[string]int)
	jCount := 0

	for _, char := range s {
		if string(char) == "J" {
			jCount++
		} else {
			counts[string(char)]++
		}
	}

	var maxCard string
	maxCount := 0
	for card, count := range counts {
		if count > maxCount {
			maxCount = count
			maxCard = card
		}
	}
	if maxCard != "" {
		counts[maxCard] += jCount
	}

	var repetitions []int
	for _, count := range counts {
		repetitions = append(repetitions, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(repetitions)))
	return repetitions
}

func handStrengthValue(repetitions []int) int {
	if len(repetitions) == 1 && repetitions[0] == 5 {
		return 6 // Five of a kind
	} else if len(repetitions) == 2 && repetitions[0] == 4 && repetitions[1] == 1 {
		return 5 // Four of a kind
	} else if len(repetitions) == 2 && repetitions[0] == 3 && repetitions[1] == 2 {
		return 4 // Full house
	} else if len(repetitions) == 3 && repetitions[0] == 3 && repetitions[1] == 1 && repetitions[2] == 1 {
		return 3 // Three of a kind
	} else if len(repetitions) == 3 && repetitions[0] == 2 && repetitions[1] == 2 && repetitions[2] == 1 {
		return 2 // Two pair
	} else if len(repetitions) == 4 && repetitions[0] == 2 && repetitions[1] == 1 && repetitions[2] == 1 && repetitions[3] == 1 {
		return 1 // One pair
	} else if len(repetitions) == 5 {
		return 0 // High card
	}
	return -1
}

func typeHand(strHands string) int {
	return handStrengthValue(countRepetitions(strHands))
}

func typeHandWithJoker(strHands string) int {
	return handStrengthValue(countRepetitionsWithJoker(strHands))
}

func compareHands(hand1, hand2 []string) int {
	for i := 0; i < len(hand1); i++ {
		if cardValues[hand1[i]] > cardValues[hand2[i]] {
			return 1
		} else if cardValues[hand1[i]] < cardValues[hand2[i]] {
			return -1
		}
	}
	return 0
}

func compareHandsWithJoker(hand1, hand2 []string) int {
	for i := 0; i < len(hand1); i++ {
		if cardValuesWithJoker[hand1[i]] > cardValuesWithJoker[hand2[i]] {
			return 1
		} else if cardValuesWithJoker[hand1[i]] < cardValuesWithJoker[hand2[i]] {
			return -1
		}
	}
	return 0
}

func rankHands(hands []string, handTypes []int, handTypeToRank int) []int {
	var indicesToRank []int

	for i, handType := range handTypes {
		if handType == handTypeToRank {
			indicesToRank = append(indicesToRank, i)
		}
	}

	sort.Slice(indicesToRank, func(i, j int) bool {
		strHand1 := strings.Split(hands[indicesToRank[i]], " ")
		hand1 := strings.Split(strHand1[0], "")
		strHand2 := strings.Split(hands[indicesToRank[j]], " ")
		hand2 := strings.Split(strHand2[0], "")
		return compareHands(hand1, hand2) < 0
	})

	return indicesToRank
}

func rankHandsWithJoker(hands []string, handTypes []int, handTypeToRank int) []int {
	var indicesToRank []int

	for i, handType := range handTypes {
		if handType == handTypeToRank {
			indicesToRank = append(indicesToRank, i)
		}
	}

	sort.Slice(indicesToRank, func(i, j int) bool {
		strHand1 := strings.Split(hands[indicesToRank[i]], " ")
		hand1 := strings.Split(strHand1[0], "")
		strHand2 := strings.Split(hands[indicesToRank[j]], " ")
		hand2 := strings.Split(strHand2[0], "")
		return compareHandsWithJoker(hand1, hand2) < 0
	})
	return indicesToRank
}

func part1(str string) int {
	var hands []string = strings.Split(str, "\r\n")
	var typeHands []int = make([]int, len(hands))
	for i, hand := range hands {
		var handCards []string = strings.Split(hand, " ")
		typeHands[i] = typeHand(handCards[0])
	}
	var rankingHands []int = make([]int, len(hands))
	var rank int = 1
	for i := 0; i < 7; i++ {
		var indices []int = rankHands(hands, typeHands, i)
		for _, i := range indices {
			rankingHands[i] = rank
			rank++
		}
	}
	var sum int = 0
	for j, r := range rankingHands {
		var handCards []string = strings.Split(hands[j], " ")
		points, _ := strconv.Atoi(handCards[1])
		sum += r * points
	}
	return sum
}

func part2(str string) int {
	var hands []string = strings.Split(str, "\r\n")
	var typeHands []int = make([]int, len(hands))
	for i, hand := range hands {
		var handCards []string = strings.Split(hand, " ")
		typeHands[i] = typeHandWithJoker(handCards[0])
	}
	var rankingHands []int = make([]int, len(hands))
	var rank int = 1
	for i := 0; i < 7; i++ {
		var indices []int = rankHandsWithJoker(hands, typeHands, i)
		for _, i := range indices {
			rankingHands[i] = rank
			rank++
		}
	}
	var sum int = 0
	for j, r := range rankingHands {
		var handCards []string = strings.Split(hands[j], " ")
		points, _ := strconv.Atoi(handCards[1])
		sum += r * points
	}
	return sum
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
