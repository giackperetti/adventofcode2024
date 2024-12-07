package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concatenate(a, b int) int {
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return concatenated
}

func solve(nums []int, target int) bool {
	n := len(nums)
	if n == 1 {
		return nums[0] == target
	}

	var dfs func([]int, int, int, int) bool
	dfs = func(current []int, index int, value int, mode int) bool {
		if index == len(current) {
			return value == target
		}

		for operator := 0; operator < 3; operator++ {
			var nextValue int
			switch operator {
			case 0: // Addition
				nextValue = value + current[index]
			case 1: // Multiplication
				nextValue = value * current[index]
			case 2: // Concatenation
				nextValue = concatenate(value, current[index])
			}

			if dfs(current, index+1, nextValue, operator) {
				return true
			}
		}

		return false
	}

	for start := 0; start < n; start++ {
		for mode := 0; mode < 3; mode++ {
			if dfs(nums[start:], 1, nums[start], mode) {
				return true
			}
		}
	}

	return false
}

func isValidEquation(nums []int, target int) bool {
	return solve(nums, target)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	equations := make(map[int][]int)
	correctEquations := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue
		}

		valuesStr := strings.Fields(parts[1])
		var values []int
		for _, val := range valuesStr {
			num, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			values = append(values, num)
		}

		equations[key] = values
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	correctEquationsSum := 0
	for testValue, equation := range equations {
		if isValidEquation(equation, testValue) {
			correctEquations[testValue] = equation
			correctEquationsSum += testValue
		}
	}

	fmt.Printf("Sum of correct equation's test values: %d\n", correctEquationsSum)
	// Part 1 Solution: 465126289353
	// Part 2 Solution: 70597497486371
}
