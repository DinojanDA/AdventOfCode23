package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputday string

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

type Neighbour struct {
	I, J        int
	Orientation string
}

type Position struct {
	I, J int
}
type NeighboursFunc func(string) []string

func neighbours(matrix [][]string, nbRow int, nbColumn int, i int, j int, orientation string) []Neighbour {
	neighbours := []Neighbour{}
	if matrix[i][j] == "." {
		switch orientation {
		case "N":
			if i > 0 {
				neighbours = append(neighbours, Neighbour{i - 1, j, "N"})
			}
		case "W":
			if j > 0 {
				neighbours = append(neighbours, Neighbour{i, j - 1, "W"})
			}
		case "S":
			if i < nbRow-1 {
				neighbours = append(neighbours, Neighbour{i + 1, j, "S"})
			}
		case "E":
			if j < nbColumn-1 {
				neighbours = append(neighbours, Neighbour{i, j + 1, "E"})
			}
		}
	} else {
		switch orientation {
		case "N":
			if matrix[i][j] == "|" {
				if i > 0 {
					neighbours = append(neighbours, Neighbour{i - 1, j, "N"})
				}
			}
			if matrix[i][j] == "-" {
				if j > 0 {
					neighbours = append(neighbours, Neighbour{i, j - 1, "W"})
				}
				if j < nbColumn-1 {
					neighbours = append(neighbours, Neighbour{i, j + 1, "E"})
				}
			}
			if matrix[i][j] == "/" {
				if j < nbColumn-1 {
					neighbours = append(neighbours, Neighbour{i, j + 1, "E"})
				}
			}
			if matrix[i][j] == "\\" {
				if j > 0 {
					neighbours = append(neighbours, Neighbour{i, j - 1, "W"})
				}

			}
			// East direction
		case "E":
			if matrix[i][j] == "|" {
				if i > 0 {
					neighbours = append(neighbours, Neighbour{i - 1, j, "N"})
				}
				if i < nbRow-1 {
					neighbours = append(neighbours, Neighbour{i + 1, j, "S"})
				}
			}
			if matrix[i][j] == "-" {
				if j < nbColumn-1 {
					neighbours = append(neighbours, Neighbour{i, j + 1, "E"})
				}
			}
			if matrix[i][j] == "/" {
				if i > 0 {
					neighbours = append(neighbours, Neighbour{i - 1, j, "N"})
				}
			}
			if matrix[i][j] == "\\" {
				if i < nbRow-1 {
					neighbours = append(neighbours, Neighbour{i + 1, j, "S"})
				}
			}
		case "S":
			if matrix[i][j] == "|" {
				if i < nbRow-1 {
					neighbours = append(neighbours, Neighbour{i + 1, j, "S"})
				}
			}
			if matrix[i][j] == "-" {
				if j > 0 {
					neighbours = append(neighbours, Neighbour{i, j - 1, "W"})
				}
				if j < nbColumn-1 {
					neighbours = append(neighbours, Neighbour{i, j + 1, "E"})
				}
			}
			if matrix[i][j] == "/" {
				if j > 0 {
					neighbours = append(neighbours, Neighbour{i, j - 1, "W"})
				}
			}
			if matrix[i][j] == "\\" {
				if j < nbColumn-1 {
					neighbours = append(neighbours, Neighbour{i, j + 1, "E"})
				}
			}

			// West direction
		case "W":
			if matrix[i][j] == "|" {
				if i > 0 {
					neighbours = append(neighbours, Neighbour{i - 1, j, "N"})
				}
				if i < nbRow-1 {
					neighbours = append(neighbours, Neighbour{i + 1, j, "S"})
				}
			}
			if matrix[i][j] == "-" {
				if j > 0 {
					neighbours = append(neighbours, Neighbour{i, j - 1, "W"})
				}
			}
			if matrix[i][j] == "/" {
				if i < nbRow-1 {
					neighbours = append(neighbours, Neighbour{i + 1, j, "S"})
				}
			}
			if matrix[i][j] == "\\" {
				if i > 0 {
					neighbours = append(neighbours, Neighbour{i - 1, j, "N"})
				}
			}

		}
	}
	return neighbours
}

func bfs(matrix [][]string, nbRow int, nbColumn int, startI int, startJ int, startOrientation string) int {
	frontier := list.New()
	frontier.PushBack(Neighbour{startI, startJ, startOrientation})

	visited := make(map[Neighbour]bool)
	energized := make(map[Position]bool)
	visited[Neighbour{startI, startJ, startOrientation}] = true
	energized[Position{startI, startJ}] = true

	for frontier.Len() > 0 {
		element := frontier.Front()
		current := element.Value.(Neighbour)
		frontier.Remove(element)

		if _, exists := energized[Position{current.I, current.J}]; !exists {
			energized[Position{current.I, current.J}] = true
		}

		for _, next := range neighbours(matrix, nbRow, nbColumn, current.I, current.J, current.Orientation) {
			if _, exists := visited[next]; !exists {
				visited[next] = true
				frontier.PushBack(next)
			}
		}
	}
	return len(energized)
}

func part1(str string) int {
	matrix, nbRow, nbColumn := conversionMatrix(str)
	return bfs(matrix, nbRow, nbColumn, 0, 0, "E")
}

func part2(str string) int {
	matrix, nbRow, nbColumn := conversionMatrix(str)
	var maximum = 0
	var i int
	var j int
	var current int

	i = 0
	j = 0
	for j < nbColumn {
		if current = bfs(matrix, nbRow, nbColumn, i, j, "S"); current > maximum {
			maximum = current
		}
		j++
	}

	i = nbRow - 1
	j = 0
	for j < nbColumn {
		if current = bfs(matrix, nbRow, nbColumn, i, j, "N"); current > maximum {
			maximum = current
		}
		j++
	}

	i = 0
	j = 0
	for i < nbRow-1 {
		if current = bfs(matrix, nbRow, nbColumn, i, j, "W"); current > maximum {
			maximum = current
		}
		i++
	}

	i = 0
	j = nbColumn - 1
	for i < nbRow-1 {
		if current = bfs(matrix, nbRow, nbColumn, i, j, "W"); current > maximum {
			maximum = current
		}
		i++
	}

	return maximum
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
