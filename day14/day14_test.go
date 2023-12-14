package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 136
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}

func TestPart2(t *testing.T) {
	var sol21 int = part2(inputtest, 1)
	var expected21 int = 87
	if sol21 != expected21 {
		t.Errorf("Expected %d, got %d", expected21, sol21)
	}
	var sol22 int = part2(inputtest, 2)
	var expected22 int = 69
	if sol22 != expected22 {
		t.Errorf("Expected %d, got %d", expected22, sol22)
	}
	var sol23 int = part2(inputtest, 3)
	var expected23 int = 69
	if sol23 != expected23 {
		t.Errorf("Expected %d, got %d", expected23, sol23)
	}
	var sol2 int = part2(inputtest, 1000000000)
	var expected2 int = 64
	if sol2 != expected2 {
		t.Errorf("Expected %d, got %d", expected2, sol2)
	}
}
