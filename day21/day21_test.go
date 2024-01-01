package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest, 6)
	var expected int = 16
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}

func TestPart2(t *testing.T) {
	var sol21 int = part2(inputtest, 6)
	var expected21 int = 16
	if sol21 != expected21 {
		t.Errorf("Expected %d, got %d", expected21, sol21)
	}
	var sol22 int = part2(inputtest, 10)
	var expected22 int = 50
	if sol22 != expected22 {
		t.Errorf("Expected %d, got %d", expected22, sol22)
	}
	var sol23 int = part2(inputtest, 50)
	var expected23 int = 1594
	if sol23 != expected23 {
		t.Errorf("Expected %d, got %d", expected23, sol23)
	}
	var sol24 int = part2(inputtest, 100)
	var expected24 int = 6536
	if sol24 != expected24 {
		t.Errorf("Expected %d, got %d", expected24, sol24)
	}
	var sol25 int = part2(inputtest, 500)
	var expected25 int = 167004
	if sol25 != expected25 {
		t.Errorf("Expected %d, got %d", expected25, sol25)
	}
}
