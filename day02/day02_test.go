package main

import (
	"testing"
)

func TestTransformMap(t *testing.T) {
	testStr := "3 red, 2 blue, 1 green"
	expected := map[string]int{"red": 3, "blue": 2, "green": 1}
	result := TransformMap(testStr)

	for color, count := range expected {
		if result[color] != count {
			t.Errorf("Expected %d for %s, got %d", count, color, result[color])
		}
	}
}

func TestPower(t *testing.T) {
	testStr := "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	expected := 630
	result := Power(testStr)

	if result != expected {
		t.Errorf("Expected power to be %d, got %d", expected, result)
	}
}
