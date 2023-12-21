package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 19114
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}

func TestPart2(t *testing.T) {
	var sol2 int = part2(inputtest)
	var expected2 int = 167409079868000
	if sol2 != expected2 {
		t.Errorf("Expected %d, got %d", expected2, sol2)
	}
}
