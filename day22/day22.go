package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputday string

type Position struct {
	x, y, z int
}

type Block struct {
	p1, p2 Position
}

func obtainBlocks(str string) []Block {
	var blocks []Block
	lines := strings.Split(str, "\r\n")
	for _, line := range lines {
		positions := strings.Split(line, "~")
		positions1 := strings.Split(positions[0], ",")
		x1, _ := strconv.Atoi(positions1[0])
		y1, _ := strconv.Atoi(positions1[1])
		z1, _ := strconv.Atoi(positions1[2])
		positions2 := strings.Split(positions[1], ",")
		x2, _ := strconv.Atoi(positions2[0])
		y2, _ := strconv.Atoi(positions2[1])
		z2, _ := strconv.Atoi(positions2[2])
		blocks = append(blocks, Block{Position{x1, y1, z1}, Position{x2, y2, z2}})
	}
	return blocks
}

func sortBlocksByMinZ(blocks []Block) {
	sort.Slice(blocks, func(i, j int) bool {
		minZi := min(blocks[i].p1.z, blocks[i].p2.z)
		minZj := min(blocks[j].p1.z, blocks[j].p2.z)
		return minZi < minZj
	})
}

func maxDimensions(blocks []Block) (int, int, int) {
	maxX, maxY, maxZ := 0, 0, 0
	for _, block := range blocks {
		maxX = max(maxX, max(block.p1.x, block.p2.x))
		maxY = max(maxY, max(block.p1.y, block.p2.y))
		maxZ = max(maxZ, max(block.p1.z, block.p2.z))
	}
	return maxX, maxY, maxZ
}

func conversionMatrix(blocks []Block) [][][]int {
	maxX, maxY, maxZ := maxDimensions(blocks)
	matrix := make([][][]int, maxX+1)
	for i := range matrix {
		matrix[i] = make([][]int, maxY+1)
		for j := range matrix[i] {
			matrix[i][j] = make([]int, maxZ+1)
		}
	}
	for i, block := range blocks {
		if block.p2.x != block.p1.x {
			if block.p2.x < block.p1.x {
				for k := block.p2.x; k < block.p1.x+1; k++ {
					matrix[k][block.p2.y][block.p2.z] = i + 1
				}
			} else {
				for k := block.p1.x; k < block.p2.x+1; k++ {
					matrix[k][block.p2.y][block.p2.z] = i + 1
				}
			}
		} else if block.p2.y != block.p1.y {
			if block.p2.y < block.p1.y {
				for k := block.p2.y; k < block.p1.y+1; k++ {
					matrix[block.p2.x][k][block.p2.z] = i + 1
				}
			} else {
				for k := block.p1.y; k < block.p2.y+1; k++ {
					matrix[block.p2.x][k][block.p2.z] = i + 1
				}
			}
		} else if block.p2.z != block.p1.z {
			if block.p2.z < block.p1.z {
				for k := block.p2.z; k < block.p1.z+1; k++ {
					matrix[block.p2.x][block.p2.y][k] = i + 1
				}
			} else {
				for k := block.p1.z; k < block.p2.z+1; k++ {
					matrix[block.p2.x][block.p2.y][k] = i + 1
				}
			}
		} else {
			matrix[block.p2.x][block.p2.y][block.p2.z] = i + 1
		}
	}
	return matrix
}

func canMove(block Block, matrix [][][]int, myMap map[int]bool) bool {
	if block.p2.z == 1 || block.p1.z == 1 {
		return false
	}
	if block.p2.x != block.p1.x {
		if block.p2.x < block.p1.x {
			for k := block.p2.x; k < block.p1.x+1; k++ {
				if _, exists := myMap[matrix[k][block.p2.y][block.p2.z-1]]; !exists {
					return false
				}
			}
			return true
		} else {
			for k := block.p1.x; k < block.p2.x+1; k++ {
				if _, exists := myMap[matrix[k][block.p2.y][block.p2.z-1]]; !exists {
					return false
				}
			}
			return true
		}
	} else if block.p2.y != block.p1.y {
		if block.p2.y < block.p1.y {
			for k := block.p2.y; k < block.p1.y+1; k++ {
				if _, exists := myMap[matrix[block.p2.x][k][block.p2.z-1]]; !exists {
					return false
				}
			}
			return true
		} else {
			for k := block.p1.y; k < block.p2.y+1; k++ {
				if _, exists := myMap[matrix[block.p2.x][k][block.p2.z-1]]; !exists {
					return false
				}
			}
			return true
		}
	} else if block.p2.z != block.p1.z {
		if block.p2.z < block.p1.z {
			if _, exists := myMap[matrix[block.p2.x][block.p2.y][block.p2.z-1]]; !exists {
				return false
			}
			return true
		} else {
			if _, exists := myMap[matrix[block.p1.x][block.p1.y][block.p1.z-1]]; !exists {
				return false
			}
			return true
		}
	} else {
		if _, exists := myMap[matrix[block.p1.x][block.p1.y][block.p1.z-1]]; !exists {
			return false
		}
		return true
	}
}

func move(block *Block, matrix [][][]int) {
	var number int = matrix[block.p2.x][block.p2.y][block.p2.z]
	if block.p2.x != block.p1.x {
		if block.p2.x < block.p1.x {
			for k := block.p2.x; k < block.p1.x+1; k++ {
				matrix[k][block.p2.y][block.p2.z] = 0
				matrix[k][block.p2.y][block.p2.z-1] = number
			}
		} else {
			for k := block.p1.x; k < block.p2.x+1; k++ {
				matrix[k][block.p2.y][block.p2.z] = 0
				matrix[k][block.p2.y][block.p2.z-1] = number
			}
		}
	} else if block.p2.y != block.p1.y {
		if block.p2.y < block.p1.y {
			for k := block.p2.y; k < block.p1.y+1; k++ {
				matrix[block.p2.x][k][block.p2.z] = 0
				matrix[block.p2.x][k][block.p2.z-1] = number
			}
		} else {
			for k := block.p1.y; k < block.p2.y+1; k++ {
				matrix[block.p2.x][k][block.p2.z] = 0
				matrix[block.p2.x][k][block.p2.z-1] = number
			}
		}
	} else if block.p2.z != block.p1.z {
		if block.p2.z < block.p1.z {
			matrix[block.p2.x][block.p2.y][block.p1.z] = 0
			matrix[block.p2.x][block.p2.y][block.p2.z-1] = number
		} else {
			matrix[block.p2.x][block.p2.y][block.p2.z] = 0
			matrix[block.p2.x][block.p2.y][block.p1.z-1] = number
		}
	} else {
		matrix[block.p2.x][block.p2.y][block.p2.z] = 0
		matrix[block.p2.x][block.p2.y][block.p2.z-1] = number
	}
	block.p1.z--
	block.p2.z--
}

func canBeDestroyed(block Block, blocks []Block, matrix [][][]int) bool {
	number := matrix[block.p2.x][block.p2.y][block.p2.z]
	movedBlocks := make(map[int]bool)
	movedBlocks[0] = true
	movedBlocks[number] = true
	for i, blockV := range blocks {
		if number != i+1 {
			if canMove(blockV, matrix, movedBlocks) {
				return false
			}
		}
	}
	return true
}

func part1(str string) int {
	var count int = 0
	blocks := obtainBlocks(str)
	sortBlocksByMinZ(blocks)
	matrix := conversionMatrix(blocks)
	for i := range blocks {
		movedBlocks := make(map[int]bool)
		movedBlocks[0] = true
		for canMove(blocks[i], matrix, movedBlocks) {
			move(&blocks[i], matrix)
		}
	}

	for _, blockDestroyed := range blocks {
		if canBeDestroyed(blockDestroyed, blocks, matrix) {
			count++
		}
	}
	return count
}

func part2(str string) int {
	var count int = 0
	blocks := obtainBlocks(str)
	sortBlocksByMinZ(blocks)
	matrix := conversionMatrix(blocks)
	for i := range blocks {
		movedBlocks := make(map[int]bool)
		movedBlocks[0] = true
		for canMove(blocks[i], matrix, movedBlocks) {
			move(&blocks[i], matrix)
		}
	}
	for i, _ := range blocks {
		movedBlocks := make(map[int]bool)
		movedBlocks[0] = true
		movedBlocks[i+1] = true
		for j, block := range blocks {
			if i != j {
				if canMove(block, matrix, movedBlocks) {
					count += 1
					movedBlocks[j+1] = true
				}
			}
		}
	}
	return count
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
