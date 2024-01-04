package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var inputday string

func convertToTabInt(str string) []int {
	numbersStr := strings.Split(str, ",")
	var numbers []int
	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func sumTab(tab []int) int {
	var len int = len(tab)
	var sum int = 0
	var i int = 0
	for i < len {
		sum += tab[i]
		i++
	}
	return sum
}

func splitIntoGroups(s string) []string {
	// Split the string by '.'
	parts := strings.Split(s, ".")
	// Filter out empty strings
	var groups []string
	for _, part := range parts {
		if part != "" {
			groups = append(groups, part)
		}
	}
	return groups
}

func transformRecord(record string) string {
	strTab := strings.Split(record, " ")
	pattern := strTab[0]
	counts := strTab[1]

	repeatedPatterns := strings.Repeat(pattern+"?", 5)
	repeatedCounts := strings.Repeat(counts+",", 5)

	transformedRecord := strings.TrimSuffix(repeatedPatterns, "?") + " " + strings.TrimSuffix(repeatedCounts, ",")
	return transformedRecord
}

func generatePossibilities(str string, currentIndex int) []string {
	strTab := strings.Split(str, " ")
	pattern := strTab[0]
	strInt := convertToTabInt(strTab[1])
	totalHashes := sumTab(strInt)
	if currentIndex == len(pattern) {
		if strings.Count(pattern, "#") == totalHashes {
			return []string{str}
		}
		return []string{}
	}

	if pattern[currentIndex] != '?' {
		return generatePossibilities(str, currentIndex+1)
	}

	withHash := str[:currentIndex] + "#" + str[currentIndex+1:]
	withDot := str[:currentIndex] + "." + str[currentIndex+1:]

	possibilitiesWithHash := generatePossibilities(withHash, currentIndex+1)
	possibilitiesWithDot := generatePossibilities(withDot, currentIndex+1)

	return append(possibilitiesWithHash, possibilitiesWithDot...)
}

func respectArrangement(str string) bool {
	strTab := strings.Fields(str)
	pattern := strTab[0]
	counts := convertToTabInt(strTab[1])
	hashGroups := splitIntoGroups(pattern)
	if len(hashGroups) != len(counts) {
		return false
	}

	for i, count := range counts {
		if hashGroups[i] != strings.Repeat("#", count) {
			return false
		}
	}
	return true
}

func part1(str string) int {
	var lines []string = strings.Split(str, "\r\n")
	var sum int = 0
	for _, line := range lines {
		possibleArrangements := generatePossibilities(line, 0)
		for _, possibleArrangement := range possibleArrangements {
			if respectArrangement(possibleArrangement) {
				sum += 1
			}
		}
	}
	return sum
}

func part2(str string) int {
	lines := strings.Split(str, "\r\n")
	sum := 0
	var wg sync.WaitGroup

	// Channel to collect results from goroutines
	results := make(chan int, len(lines))

	for _, line := range lines {
		wg.Add(1) // Add to the wait group for each goroutine we are about to launch

		// Start a goroutine
		go func(l string) {
			defer wg.Done() // Signal the wait group once this goroutine is done

			transformedRecord := transformRecord(l)
			possibilities := generatePossibilities(transformedRecord, 0)

			localSum := 0
			for _, possibility := range possibilities {
				if respectArrangement(possibility) {
					localSum += 1
				}
			}

			results <- localSum // Send the local sum to the results channel
		}(line) // Pass the current line to the goroutine
	}

	// Close the results channel when all goroutines are done
	go func() {
		wg.Wait()      // Wait for all goroutines to finish
		close(results) // Close the channel to signal completion
	}()

	// Sum results from the channel
	for result := range results {
		sum += result
	}

	return sum
}

func main() {
	var sol1 int = part1(inputday)
	fmt.Println(sol1)

	var sol2 int = part2(inputday)
	fmt.Println(sol2)
}
