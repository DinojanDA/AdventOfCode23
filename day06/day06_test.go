package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputtest string

func TestOneRace(t *testing.T) {
	var time int = 7
	var distance int = 9
	var count int = 0
	for i := 0; i < 8; i++ {
		if arrive(i, time, distance) {
			count += 1
		}
	}
	var sol int = count
	var expected int = 4
	if sol != expected {
		t.Errorf("Expected %d, got %d", expected, sol)
	}
}

func TestConvertToTabInt(t *testing.T) {
	var str string = "      7  15   30"
	var expected []int = []int{7, 15, 30}
	tab := convertToTabInt(str)
	for i := 0; i < len(tab); i++ {
		if tab[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], tab[i])
		}
	}
}

func TestPart1(t *testing.T) {
	var sol1 int = part1(inputtest)
	var expected int = 288
	if sol1 != expected {
		t.Errorf("Expected %d, got %d", expected, sol1)
	}
}

func TestPart2(t *testing.T) {
	var sol2 int = part2(inputtest)
	var expected int = 71503
	if sol2 != expected {
		t.Errorf("Expected %d, got %d", expected, sol2)
	}
}
