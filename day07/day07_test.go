package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestCompareHands(t *testing.T) {
	var hand1 []string = strings.Split("KK677", "")
	var hand2 []string = strings.Split("KTJJT", "")
	sol := compareHands(hand1, hand2)
	expected := 1
	if sol != expected {
		t.Errorf("Expected %d, got %d", expected, sol)
	}
}

func TestTypeHandWWithJoker(t *testing.T) {
	var hand string = "QJJQ2"
	sol := typeHandWithJoker(hand)
	expected := 5
	if sol != expected {
		t.Errorf("Expected %d, got %d", expected, sol)
	}
}

func TestCompareHandsWithJoker(t *testing.T) {
	var hand1 []string = strings.Split("QJJQ2", "")
	var hand2 []string = strings.Split("JKKK2", "")
	sol := compareHands(hand1, hand2)
	expected := 1
	if sol != expected {
		t.Errorf("Expected %d, got %d", expected, sol)
	}
}

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 6440
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}

func TestPart2(t *testing.T) {
	var sol2 int = part2(inputtest)
	var expected int = 5905
	if sol2 != expected {
		t.Errorf("Expected %d, got %d", expected, sol2)
	}
}
