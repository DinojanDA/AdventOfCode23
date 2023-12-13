package main

import (
	_ "embed"
	"fmt"
	"strconv"
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

func isEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findIdentical(matrix [][]string, nbRow int, nbColumn int) [][]string {
	matches := [][]string{}

	for i := 0; i < nbRow-1; i++ {
		if isEqual(matrix[i], matrix[i+1]) {
			matches = append(matches, []string{strconv.Itoa(i), "row"})
		}
	}

	for j := 0; j < nbColumn-1; j++ {
		col1 := []string{}
		col2 := []string{}
		for i := 0; i < nbRow; i++ {
			col1 = append(col1, matrix[i][j])
			col2 = append(col2, matrix[i][j+1])
		}
		if isEqual(col1, col2) {
			matches = append(matches, []string{strconv.Itoa(j), "column"})
		}
	}
	return matches
}

func checkPairs(matrix [][]string, nbRow int, nbColumn int, index int, typ string) bool {
	if typ == "row" {
		i := index
		j := index + 1
		for i >= 0 && j < nbRow {
			if !isEqual(matrix[i], matrix[j]) {
				return false
			}
			i--
			j++
		}
	} else {
		i := index
		j := index + 1
		for i >= 0 && j < nbColumn {
			colI := []string{}
			colJ := []string{}
			var row int = 0
			for row < nbRow {
				colI = append(colI, matrix[row][i])
				colJ = append(colJ, matrix[row][j])
				row++
			}
			if !isEqual(colI, colJ) {
				return false
			}
			i--
			j++
		}
	}
	return true
}

func findAndCorrectSmudge(matrix [][]string, nbRow int, nbColumn int) (correctedMatrix [][]string) {
	// Copie de la matrice pour éviter de modifier l'originale pendant la recherche.
	correctedMatrix = make([][]string, len(matrix))
	for i := range matrix {
		correctedMatrix[i] = make([]string, len(matrix[i]))
		copy(correctedMatrix[i], matrix[i])
	}
	i := 0
	for i < nbRow {
		j := 0
		for j < nbColumn {
			char := matrix[i][j]
			if char == "." {
				correctedMatrix[i][j] = "#"
			} else {
				correctedMatrix[i][j] = "."
			}
			// Vérifier si la matrice a maintenant une ligne de réflexion valide.
			if hasValidReflectionLine(correctedMatrix, matrix, nbRow, nbColumn) {
				return correctedMatrix
			}
			correctedMatrix[i][j] = char
			j++
		}
		i++
	}
	return matrix
}

func hasValidReflectionLine(correctedMatrix [][]string, matrix [][]string, nbRow int, nbColumn int) bool {
	a, ind1 := reflexionLine(correctedMatrix, nbRow, nbColumn)
	b, ind2 := reflexionLine(matrix, nbRow, nbColumn)
	return a != b || !(isEqual(ind1, ind2))
}

func reflexionLine(matrix [][]string, nbRow int, nbColumn int) (int, []string) {
	var a int
	var ind []string
	indices := findIdentical(matrix, nbRow, nbColumn)
	for _, indice := range indices {
		index, _ := strconv.Atoi(indice[0])
		if checkPairs(matrix, nbRow, nbColumn, index, indice[1]) {
			if indice[1] == "row" {
				a = 100 * (index + 1)
				ind = indice
				break
			} else {
				a = index + 1
				ind = indice
				break
			}
		}
	}
	return a, ind
}

func part1(str string) int {
	var puzzles []string = strings.Split(str, "\r\n\r\n")
	var sum int = 0
	for _, puzzle := range puzzles {
		var a int
		matrix, nbRow, nbColumn := conversionMatrix(puzzle)
		a, _ = reflexionLine(matrix, nbRow, nbColumn)
		sum += a
	}
	return sum
}

func part2(str string) int {
	var puzzles []string = strings.Split(str, "\r\n\r\n")
	var sum int = 0
	for _, puzzle := range puzzles {
		var a int
		matrix, nbRow, nbColumn := conversionMatrix(puzzle)
		correctedMatrix := findAndCorrectSmudge(matrix, nbRow, nbColumn)
		a, _ = reflexionLine(correctedMatrix, nbRow, nbColumn)
		sum += a
	}
	return sum
}
func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
