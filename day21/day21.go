package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input_test.txt
var inputday string

type Position struct {
	I, J int
}

type PositionLevels struct {
	I, J, levelHigh, levelRight int
}

func conversionMatrix(str string) ([][]string, int, int) {
	var nbColonnes int
	var matrix [][]string
	lignes := strings.Split(str, "\r\n")
	var nbLignes int = len(lignes)
	for _, ligne := range lignes {
		caracteres := strings.Split(ligne, "")
		matrix = append(matrix, caracteres)
		nbColonnes = len(caracteres)
	}
	return matrix, nbLignes, nbColonnes
}

func identifyStart(matrix [][]string, nbLignes int, nbColonnes int) Position {
	var i int = 0
	for i < nbLignes {
		var j int = 0
		for j < nbColonnes {
			if matrix[i][j] == "S" {
				return Position{i, j}
			}
			j++
		}
		i++
	}
	return Position{-1, -1}
}

type NeighboursFunc func(string) []string

func neighbours(matrix [][]string, nbRow int, nbColumn int, position Position) []Position {
	i := position.I
	j := position.J
	neighbours := []Position{}
	if i > 0 && matrix[i-1][j] != "#" {
		neighbours = append(neighbours, Position{i - 1, j})
	}
	if j > 0 && matrix[i][j-1] != "#" {
		neighbours = append(neighbours, Position{i, j - 1})
	}

	if i < nbRow-1 && matrix[i+1][j] != "#" {
		neighbours = append(neighbours, Position{i + 1, j})
	}

	if j < nbColumn-1 && matrix[i][j+1] != "#" {
		neighbours = append(neighbours, Position{i, j + 1})
	}
	return neighbours
}

func neighboursLevels(matrix [][]string, nbRow int, nbColumn int, position PositionLevels) []PositionLevels {
	i := position.I
	j := position.J
	levelH := position.levelHigh
	levelR := position.levelRight
	neighbours := []PositionLevels{}
	if i > 0 && matrix[i-1][j] != "#" {
		neighbours = append(neighbours, PositionLevels{i - 1, j, levelH, levelR})
	}
	if i == 0 && matrix[nbRow-1][j] != "#" {
		neighbours = append(neighbours, PositionLevels{nbRow - 1, j, levelH - 1, levelR})
	}
	if j > 0 && matrix[i][j-1] != "#" {
		neighbours = append(neighbours, PositionLevels{i, j - 1, levelH, levelR})
	}
	if j == 0 && matrix[i][nbColumn-1] != "#" {
		neighbours = append(neighbours, PositionLevels{i, nbColumn - 1, levelH, levelR - 1})
	}
	if i < nbRow-1 && matrix[i+1][j] != "#" {
		neighbours = append(neighbours, PositionLevels{i + 1, j, levelH, levelR})
	}
	if i == nbRow-1 && matrix[0][j] != "#" {
		neighbours = append(neighbours, PositionLevels{0, j, levelH + 1, levelR})
	}
	if j < nbColumn-1 && matrix[i][j+1] != "#" {
		neighbours = append(neighbours, PositionLevels{i, j + 1, levelH, levelR})
	}
	if j == nbColumn-1 && matrix[i][0] != "#" {
		neighbours = append(neighbours, PositionLevels{i, 0, levelH, levelR + 1})
	}
	return neighbours
}

func count(matrix [][]string, nbRow int, nbColumn int, start Position, steps int) int {
	currentList := list.New()
	currentList.PushBack(start)

	marked := make(map[Position]int)
	marked[start] = 1

	for i := 1; i < steps+1; i++ {
		max := currentList.Len()
		for j := 0; j < max; j++ {
			element := currentList.Front()
			current := element.Value.(Position)
			currentList.Remove(element)

			if k, _ := marked[current]; k != i {
				marked[current] = 0
			}

			for _, next := range neighbours(matrix, nbRow, nbColumn, current) {
				if k, exists := marked[next]; !exists {
					marked[next] = i
					currentList.PushBack(next)
				} else {
					if k != i {
						marked[next] = i
						currentList.PushBack(next)
					}
				}
			}
		}
	}
	return currentList.Len()
}

func countLevels(matrix [][]string, nbRow int, nbColumn int, start PositionLevels, steps int) int {
	currentList := list.New()
	currentList.PushBack(start)

	marked := make(map[PositionLevels]int)
	marked[start] = 1

	for i := 1; i < steps+1; i++ {
		max := currentList.Len()
		for j := 0; j < max; j++ {
			element := currentList.Front()
			current := element.Value.(PositionLevels)
			currentList.Remove(element)

			if k, _ := marked[current]; k != i {
				marked[current] = 0
			}

			for _, next := range neighboursLevels(matrix, nbRow, nbColumn, current) {
				if k, exists := marked[next]; !exists {
					marked[next] = i
					currentList.PushBack(next)
				} else {
					if k != i {
						marked[next] = i
						currentList.PushBack(next)
					}
				}
			}
		}
	}
	return currentList.Len()
}

func part1(str string, steps int) int {
	matrix, nbRow, nbColumn := conversionMatrix(str)
	start := identifyStart(matrix, nbRow, nbColumn)
	return count(matrix, nbRow, nbColumn, start, steps)
}

func part2(str string, steps int) int {
	matrix, nbRow, nbColumn := conversionMatrix(str)
	start := identifyStart(matrix, nbRow, nbColumn)
	return countLevels(matrix, nbRow, nbColumn, PositionLevels{start.I, start.J, 0, 0}, steps)
}

func main() {
	var sol1 int = part1(inputday, 64)
	fmt.Println(sol1)

	var sol2 int = part2(inputday, 50)
	fmt.Println(sol2)
}
