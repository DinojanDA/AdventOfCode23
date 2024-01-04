package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

type Instruction struct {
	Orientation string
	Length      int
	Color       string
}

type Position struct {
	I, J int
}

func convertInstructions(str string) []Instruction {
	lines := strings.Split(str, "\r\n")
	var instructions []Instruction
	for _, line := range lines {
		parts := strings.Split(line, " ")
		orientation := parts[0]
		length, _ := strconv.Atoi(parts[1])
		color := strings.ReplaceAll(parts[2], "(", "")
		color = strings.ReplaceAll(color, ")", "")
		instructions = append(instructions, Instruction{orientation, length, color})
	}
	return instructions
}

func move(space map[Position]bool, pos *Position, instruction Instruction) {
	switch instruction.Orientation {
	case "U":
		for k := 1; k <= instruction.Length; k++ {
			space[Position{pos.I + k, pos.J}] = true
		}
		pos.I += instruction.Length
	case "D":
		for k := 1; k <= instruction.Length; k++ {
			space[Position{pos.I - k, pos.J}] = true
		}
		pos.I -= instruction.Length
	case "R":
		for k := 1; k <= instruction.Length; k++ {
			space[Position{pos.I, pos.J + k}] = true
		}
		pos.J += instruction.Length
	case "L":
		for k := 1; k <= instruction.Length; k++ {
			space[Position{pos.I, pos.J - k}] = true
		}
		pos.J -= instruction.Length
	}
}

func moveList(space []Position, pos *Position, instruction Instruction) []Position {
	switch instruction.Orientation {
	case "U":
		space = append(space, Position{pos.I + instruction.Length, pos.J})
		pos.I += instruction.Length
	case "D":
		space = append(space, Position{pos.I - instruction.Length, pos.J})
		pos.I -= instruction.Length
	case "R":
		space = append(space, Position{pos.I, pos.J + instruction.Length})
		pos.J += instruction.Length
	case "L":
		space = append(space, Position{pos.I, pos.J - instruction.Length})
		pos.J -= instruction.Length
	}
	return space
}

func neighbours(space map[Position]bool, i int, j int) []Position {
	var neighboursList []Position
	if _, exists := space[Position{i - 1, j}]; !exists {
		neighboursList = append(neighboursList, Position{i - 1, j})
	}
	if _, exists := space[Position{i + 1, j}]; !exists {
		neighboursList = append(neighboursList, Position{i + 1, j})
	}
	if _, exists := space[Position{i, j - 1}]; !exists {
		neighboursList = append(neighboursList, Position{i, j - 1})
	}
	if _, exists := space[Position{i, j + 1}]; !exists {
		neighboursList = append(neighboursList, Position{i, j + 1})
	}
	return neighboursList
}

func bfs(space map[Position]bool, pos Position) int {
	frontier := list.New()
	frontier.PushBack(pos)

	for frontier.Len() > 0 {
		element := frontier.Front()
		current := element.Value.(Position)
		frontier.Remove(element)

		for _, next := range neighbours(space, current.I, current.J) {
			space[next] = true
			frontier.PushBack(next)

		}
	}
	return len(space)
}

func hexToDirectionDistance(hexCode string) (string, int) {
	hexCode = hexCode[1:]
	distanceHex := hexCode[:5]
	directionHex := hexCode[5:]

	distance, _ := strconv.ParseInt(distanceHex, 16, 64)
	directions := map[string]string{"0": "R", "1": "D", "2": "L", "3": "U"}
	direction, _ := directions[directionHex]

	return direction, int(distance)
}

func convertHexaInstructions(str string) []Instruction {
	lines := strings.Split(str, "\r\n")
	var instructions []Instruction
	for _, line := range lines {
		parts := strings.Split(line, " ")
		color := strings.ReplaceAll(parts[2], "(", "")
		color = strings.ReplaceAll(color, ")", "")
		orientation, length := hexToDirectionDistance(color)
		instructions = append(instructions, Instruction{orientation, length, color})
	}
	return instructions
}

func shoelaceArea(points []Position) int {
	var area int
	for i := 0; i < len(points); i++ {
		j := (i + 1) % len(points)
		area += points[i].I * points[j].J
		area -= points[j].I * points[i].J
	}
	return abs(area) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(str string) int {
	space := make(map[Position]bool)
	var instructions []Instruction = convertInstructions(str)
	start := Position{0, 0}
	for _, instruction := range instructions {
		move(space, &start, instruction)
	}
	return bfs(space, Position{-1, 1})
}

func part2(str string) int {
	var space []Position
	var frontier int
	start := Position{0, 0}
	var instructions []Instruction = convertHexaInstructions(str)
	for _, instruction := range instructions {
		space = moveList(space, &start, instruction)
		frontier += instruction.Length
	}
	var area = shoelaceArea(space) + frontier/2 + 1
	return area
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
