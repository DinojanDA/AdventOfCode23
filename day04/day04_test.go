package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestSearch(t *testing.T) {
	var winningCards []string = []string{"2", "3", "4"}
	var myCards []string = []string{"4", "1", "8", "3", "7"}
	var expected []string = []string{"4", "3"}
	sols := search(winningCards, myCards)
	for i, sol := range sols {
		if sol != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], sol)
		}
	}
}

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 13
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}
func TestPart2(t *testing.T) {
	var sol2 int = part2(inputtest)
	var expected int = 30
	if sol2 != expected {
		t.Errorf("Expected %d, got %d", expected, sol2)
	}
}
