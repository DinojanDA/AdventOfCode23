package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input_test.txt
var inputtest string

//go:embed input_test2.txt
var inputtest2 string

//go:embed input_test3.txt
var inputtest3 string

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

func TestFindFinalA(t *testing.T) {
	var strTab []string = strings.Split(inputtest3, "\r\n")
	var sol []string = findFinalA(strTab)
	var expected []string = []string{"11A", "22A"}
	for i := 0; i < len(expected); i++ {
		if sol[i] != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], sol[i])
		}
	}
}

func TestFinalZ(t *testing.T) {
	var strTab []string = []string{"11Z", "22A"}
	sol := finalZ(strTab)
	if sol {
		t.Errorf("Expected false, got true")
	}
	var strTab1 []string = []string{"11Z", "22Z"}
	sol1 := finalZ(strTab1)
	if !sol1 {
		t.Errorf("Expected true, got false")
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

func TestPart2(t *testing.T) {
	var sol2 int = part2(inputtest3)
	var expected int = 6
	if sol2 != expected {
		t.Errorf("Expected %d, got %d", expected, sol2)
	}
}
