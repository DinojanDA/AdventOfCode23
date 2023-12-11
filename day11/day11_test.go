package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestPart1(t *testing.T) {
	var sol1 int = part(inputtest, 2)
	var expected int = 374
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}
func TestPart2(t *testing.T) {
	var sol2 int = part(inputtest, 10)
	var expected2 int = 1030
	if sol2 != expected2 {
		t.Errorf("Expected %d, got %d", expected2, sol2)
	}
	var sol3 int = part(inputtest, 100)
	var expected3 int = 8410
	if sol3 != expected3 {
		t.Errorf("Expected %d, got %d", expected3, sol3)
	}
}
