package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countOccurrences(line, word string) int {
	count := 0
	wordLen := len(word)
	for i := 0; i <= len(line)-wordLen; i++ {
		if line[i:i+wordLen] == word {
			count++
		}
	}
	return count
}

func countHorizontal(grid []string, word string) int {
	count := 0
	for _, line := range grid {
		count += countOccurrences(line, word)
		count += countOccurrences(line, reverse(word))
	}
	return count
}

func countVertical(grid []string, word string) int {
	count := 0
	rows, cols := len(grid), len(grid[0])
	for col := 0; col < cols; col++ {
		columnStr := ""
		for row := 0; row < rows; row++ {
			columnStr += string(grid[row][col])
		}
		count += countOccurrences(columnStr, word)
		count += countOccurrences(columnStr, reverse(word))
	}
	return count
}

func countDiagonals(grid []string, word string) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	for start := 0; start < rows+cols-1; start++ {
		diagonal := ""
		for row := 0; row < rows; row++ {
			col := start - row
			if col >= 0 && col < cols {
				diagonal += string(grid[row][col])
			}
		}
		count += countOccurrences(diagonal, word)
		count += countOccurrences(diagonal, reverse(word))
	}

	for start := 0; start < rows+cols-1; start++ {
		diagonal := ""
		for row := 0; row < rows; row++ {
			col := cols - 1 - start + row
			if col >= 0 && col < cols {
				diagonal += string(grid[row][col])
			}
		}
		count += countOccurrences(diagonal, word)
		count += countOccurrences(diagonal, reverse(word))
	}

	return count
}

func findXShapedWords(grid []string, pattern string) int {
	gridHeight := len(grid)
	gridWidth := len(grid[0])
	count := 0

	for row := 1; row < gridHeight-1; row++ {
		for col := 1; col < gridWidth-1; col++ {
			topLeftToBottomRightDiagonal := string(grid[row-1][col-1]) + string(grid[row][col]) + string(grid[row+1][col+1])
			topRightToBottomLeftDiagonal := string(grid[row-1][col+1]) + string(grid[row][col]) + string(grid[row+1][col-1])

			isCenterA := string(grid[row][col]) == string(pattern[len(pattern)/2])

			isPatternMatch := topLeftToBottomRightDiagonal == pattern && topRightToBottomLeftDiagonal == pattern
			isReversedPatternMatch := topLeftToBottomRightDiagonal == pattern && reverse(topRightToBottomLeftDiagonal) == pattern
			isReverseTlBrMatch := reverse(topLeftToBottomRightDiagonal) == pattern && topRightToBottomLeftDiagonal == pattern
			isFullReverseMatch := reverse(topLeftToBottomRightDiagonal) == pattern && reverse(topRightToBottomLeftDiagonal) == pattern

			if isCenterA && (isPatternMatch || isReversedPatternMatch || isReverseTlBrMatch || isFullReverseMatch) {
				count++
			}
		}
	}

	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	word := "XMAS"
	xShapeWord := "MAS"

	occurencesCount := countHorizontal(grid, word)
	occurencesCount += countVertical(grid, word)
	occurencesCount += countDiagonals(grid, word)

	xShapesCount := findXShapedWords(grid, xShapeWord)

	fmt.Printf("Count of 'XMAS' in all directions: %d\n", occurencesCount)        // Part 1 Solution: 2549
	fmt.Printf("Count of X-shaped 'MAS'es in all directions: %d\n", xShapesCount) // Part 2 Solution: 2003
}
