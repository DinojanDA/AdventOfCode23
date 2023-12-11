package main

import (
	_ "embed"
	"fmt"
	"math"
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

func identifyGalaxy(matrix [][]string, nbLignes int, nbColonnes int) [][2]int {
	var index [][2]int
	var i int = 0
	for i < nbLignes {
		var j int = 0
		for j < nbColonnes {
			if matrix[i][j] == "#" {
				var add [2]int = [2]int{i, j}
				index = append(index, add)
			}
			j++
		}
		i++
	}
	return index
}

func calculateDistance(galaxy1 [2]int, galaxy2 [2]int) int {
	return int(math.Abs(float64(galaxy1[0]-galaxy2[0])) + math.Abs(float64(galaxy1[1]-galaxy2[1])))
}

func rowEmpty(matrix [][]string, nbLignes int, nbColonnes int) []int {
	var index []int
	for i := 0; i < nbLignes; i++ {
		var b bool = true
		for j := 0; j < nbColonnes; j++ {
			if matrix[i][j] == "#" {
				b = false
			}
		}
		if b {
			index = append(index, i)
		}
	}
	return index
}
func columnEmpty(matrix [][]string, nbLignes int, nbColonnes int) []int {
	var index []int
	for j := 0; j < nbColonnes; j++ {
		var b bool = true
		for i := 0; i < nbLignes; i++ {
			if matrix[i][j] == "#" {
				b = false
			}
		}
		if b {
			index = append(index, j)
		}
	}
	return index
}

func isBetweenRow(indexEmpty int, galaxy1 [2]int, galaxy2 [2]int) bool {
	var i1 int = galaxy1[0]
	var i2 int = galaxy2[0]
	if i1 < i2 {
		if indexEmpty < i2 && indexEmpty > i1 {
			return true
		}
		return false
	}
	if indexEmpty < i1 && indexEmpty > i2 {
		return true
	}
	return false
}

func isBetweenColumn(indexEmpty int, galaxy1 [2]int, galaxy2 [2]int) bool {
	var i1 int = galaxy1[1]
	var i2 int = galaxy2[1]
	if i1 < i2 {
		if indexEmpty < i2 && indexEmpty > i1 {
			return true
		}
		return false
	}
	if indexEmpty < i1 && indexEmpty > i2 {
		return true
	}
	return false
}

func part(str string, nbExpansion int) int {
	var sum int = 0
	matrix, nbRow, nbColumn := conversionMatrix(str)
	var tabRowEmpty []int = rowEmpty(matrix, nbRow, nbColumn)
	var tabColumnEmpty []int = columnEmpty(matrix, nbRow, nbColumn)
	galaxies := identifyGalaxy(matrix, nbRow, nbColumn)
	var distance int = 0
	for i, galaxy1 := range galaxies {
		j := i + 1
		for j < len(galaxies) {
			galaxy2 := galaxies[j]
			distance = calculateDistance(galaxy1, galaxy2)
			for _, indexRowEmpty := range tabRowEmpty {
				if isBetweenRow(indexRowEmpty, galaxy1, galaxy2) {
					distance += nbExpansion - 1
				}
			}
			for _, indexColumnEmpty := range tabColumnEmpty {
				if isBetweenColumn(indexColumnEmpty, galaxy1, galaxy2) {
					distance += nbExpansion - 1
				}
			}
			sum += distance
			j++
		}
	}
	return sum
}

func main() {
	var sol1 int = part(inputday, 2)
	fmt.Println(sol1)

	var sol2 int = part(inputday, 1000000)
	fmt.Println(sol2)
}
