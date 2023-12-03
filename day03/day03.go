package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

func conversionMatrice(str string) ([][]string, int, int) {
	var nbColonnes int
	var matrix [][]string
	lignes := strings.Split(str, "\n")
	var nbLignes int = len(lignes)
	for _, ligne := range lignes {
		caracteres := strings.Split(ligne, "")
		matrix = append(matrix, caracteres)
		nbColonnes = len(caracteres)
	}
	return matrix, nbLignes, nbColonnes
}

func voisins(matrix [][]string, i int, j int, maxLignes int, maxColonnes int) bool {
	if i != 0 {
		voisin1 := matrix[i-1][j]
		_, err1 := strconv.Atoi(voisin1)
		if voisin1 != "." && err1 != nil {
			return true
		}
		if j != 0 {
			voisin2 := matrix[i-1][j-1]
			_, err2 := strconv.Atoi(voisin2)
			if voisin2 != "." && err2 != nil {
				return true
			}
		}
		if j != maxColonnes-1 {
			voisin3 := matrix[i-1][j+1]
			_, err3 := strconv.Atoi(voisin3)
			if voisin3 != "." && err3 != nil {
				return true
			}
		}
	}
	if i != maxLignes-1 {
		voisin4 := matrix[i+1][j]
		_, err4 := strconv.Atoi(voisin4)
		if voisin4 != "." && err4 != nil {
			return true
		}
		if j != 0 {
			voisin5 := matrix[i+1][j-1]
			_, err5 := strconv.Atoi(voisin5)
			if voisin5 != "." && err5 != nil {
				return true
			}
		}
		if j != maxColonnes-1 {
			voisin6 := matrix[i+1][j+1]
			_, err6 := strconv.Atoi(voisin6)
			if voisin6 != "." && err6 != nil {
				return true
			}
		}
	}
	if j != 0 {
		voisin7 := matrix[i][j-1]
		_, err7 := strconv.Atoi(voisin7)
		if voisin7 != "." && err7 != nil {
			return true
		}
	}
	if j != maxColonnes-1 {
		voisin8 := matrix[i][j+1]
		_, err8 := strconv.Atoi(voisin8)
		if voisin8 != "." && err8 != nil {
			return true
		}
	}
	return false
}

func chiffres(matrix [][]string, i int, nbColonnes int) []int {
	var indexNb []int
	var j int = 0
	for j < nbColonnes {
		_, err := strconv.Atoi(matrix[i][j])
		if err == nil {
			indexNb = append(indexNb, 1)
		} else {
			indexNb = append(indexNb, 0)
		}
		j++
	}
	return indexNb
}

func numbers(chiffres []int) [][]int {
	var indexNumbers [][]int
	j := 0
	for j < len(chiffres) {
		if chiffres[j] == 1 {
			var number []int
			for j < len(chiffres) && chiffres[j] == 1 {
				number = append(number, j)
				j++
			}
			indexNumbers = append(indexNumbers, number)
		} else {
			j++
		}
	}
	return indexNumbers
}

func part1(str string) int {
	var sum int = 0
	matrix, nbLignes, nbColonnes := conversionMatrice(str)
	var i int = 0
	var number string
	for i < nbLignes {
		chiffresList := chiffres(matrix, i, nbColonnes)
		numbersIndex := numbers(chiffresList)
		for _, list := range numbersIndex {
			var b bool = false
			for _, j := range list {
				if voisins(matrix, i, j, nbLignes, nbColonnes) {
					b = true
				}
				number += matrix[i][j]
			}
			if b {
				nb, _ := strconv.Atoi(number)
				sum += nb
			}
			number = ""
		}
		i++
	}
	return sum
}

func main() {
	matrix, nbLignes, nbCol := conversionMatrice(inputday)
	fmt.Println(matrix[0][2])
	fmt.Println(nbLignes, nbCol)
	fmt.Println(voisins(matrix, 4, 5, nbLignes, nbCol))
	c := chiffres(matrix, 0, nbCol)
	fmt.Println(c)
	fmt.Println(numbers(c))

	fmt.Println(part1(inputday))
}
