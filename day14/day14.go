package main

import (
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

func move(matrix [][]string, nbRow int, nbColumn int, i int, j int, direction string) [][]string {
	if direction == "north" {
		if matrix[i][j] != "O" {
			return matrix
		} else {
			var k int = i - 1
			for k >= 0 && matrix[k][j] == "." {
				k--
			}
			matrix[i][j] = "."
			matrix[k+1][j] = "O"
		}
		return matrix
	}
	if direction == "south" {
		if matrix[i][j] != "O" {
			return matrix
		} else {
			var k int = i + 1
			for k < nbRow && matrix[k][j] == "." {
				k++
			}
			matrix[i][j] = "."
			matrix[k-1][j] = "O"
		}
		return matrix
	}
	if direction == "west" {
		if matrix[i][j] != "O" {
			return matrix
		} else {
			var k int = j - 1
			for k >= 0 && matrix[i][k] == "." {
				k--
			}
			matrix[i][j] = "."
			matrix[i][k+1] = "O"
		}
		return matrix
	}
	if direction == "east" {
		if matrix[i][j] != "O" {
			return matrix
		} else {
			var k int = j + 1
			for k < nbColumn && matrix[i][k] == "." {
				k++
			}
			matrix[i][j] = "."
			matrix[i][k-1] = "O"
		}
		return matrix
	}
	return matrix
}

func count(matrix [][]string, i int, nbColumn int) int {
	var c int = 0
	var j int = 0
	for j < nbColumn {
		if matrix[i][j] == "O" {
			c++
		}
		j++
	}
	return c
}
func part1(str string) int {
	var sum int = 0
	matrix, nbRow, nbColumn := conversionMatrix(str)
	var i int = 0
	var j int = 0
	for i < nbRow {
		for j < nbColumn {
			matrix = move(matrix, nbRow, nbColumn, i, j, "north")
			j++
		}
		j = 0
		i++
	}
	i = 0
	for i < nbRow {
		sum += (nbRow - i) * count(matrix, i, nbColumn)
		i++
	}
	return sum
}

func part2(str string, nbCycles int) int {
	var sum int = 0
	matrix, nbRow, nbColumn := conversionMatrix(str)
	var cycles int = 0
	for cycles < nbCycles {
		var i int = 0
		var j int = 0
		for i < nbRow {
			for j < nbColumn {
				matrix = move(matrix, nbRow, nbColumn, i, j, "north")
				j++
			}
			j = 0
			i++
		}
		i = 0
		j = 0
		for j < nbColumn {
			for i < nbRow {
				matrix = move(matrix, nbRow, nbColumn, i, j, "west")
				i++
			}
			i = 0
			j++
		}
		i = nbRow - 1
		j = 0
		for i >= 0 {
			for j < nbColumn {
				matrix = move(matrix, nbRow, nbColumn, i, j, "south")
				j++
			}
			j = 0
			i--
		}
		i = 0
		j = nbColumn - 1
		for j >= 0 {
			for i < nbRow {
				matrix = move(matrix, nbRow, nbColumn, i, j, "east")
				i++
			}
			i = 0
			j--
		}
		cycles++
	}
	var i int = 0
	for i < nbRow {
		sum += (nbRow - i) * count(matrix, i, nbColumn)
		i++
	}
	return sum
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday, 1000000000)
	fmt.Println(sol2)
}
