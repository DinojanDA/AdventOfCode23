package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

/*var games[]string = strings.Split(inputtest, "\n")
var cards []string = strings.Split(games, ": ")
var hands []string = strings.Split(cards[1], " | ")
var winningCards []string = strings.Split(hands[0], " ")
var myCards []string = strings.Split(hands[1], " ")*/

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 13
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}
