package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 114
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}
func TestPart2(t *testing.T) {
	var sol2 int = part2(inputtest)
	var expected int = 2
	if sol2 != expected {
		t.Errorf("Expected %d, got %d", expected, sol2)
	}
}
