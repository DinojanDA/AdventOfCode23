package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

//go:embed input_test2.txt
var inputtest2 string

func TestGiveKey(t *testing.T) {
	var str string = "AAA = (BBB, CCC)"
	sol := giveKey(str)
	expected := "AAA"
	if sol != expected {
		t.Errorf("Expected %s, got %s", expected, sol)
	}
}

func TestGiveLeft(t *testing.T) {
	var str string = "AAA = (BBB, CCC)"
	sol := giveLeft(str)
	expected := "BBB"
	if sol != expected {
		t.Errorf("Expected %s, got %s", expected, sol)
	}
}

func TestGiveRight(t *testing.T) {
	var str string = "AAA = (BBB, CCC)"
	sol := giveRight(str)
	expected := "CCC"
	if sol != expected {
		t.Errorf("Expected %s, got %s", expected, sol)
	}
}

func TestPart1_1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 2
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}
func TestPart1_2(t *testing.T) {
	var sol1 int = part1(inputtest2)
	var expected int = 6
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}

/*func TestPart2(t *testing.T) {
	var sol2 int = part2(inputtest)
	var expected int = 5905
	if sol2 != expected {
		t.Errorf("Expected %d, got %d", expected, sol2)
	}
}*/
