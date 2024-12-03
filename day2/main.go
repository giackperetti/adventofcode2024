package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {
	increasing := make([]int, len(report))
	copy(increasing, report)
	sort.Ints(increasing)

	decreasing := make([]int, len(report))
	copy(decreasing, report)
	sort.Sort(sort.Reverse(sort.IntSlice(decreasing)))

	increasingOrDecreasing := slices.Equal(report, increasing) || slices.Equal(report, decreasing)

	safe := true
	for i := 0; i < len(report)-1; i++ {
		diff := math.Abs(float64(report[i] - report[i+1]))
		if !(diff >= 1 && diff <= 3) {
			safe = false
			break
		}
	}
	return increasingOrDecreasing && safe
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		strNumbers := strings.Fields(line)
		var row []int
		for _, str := range strNumbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}
			row = append(row, num)
		}
		reports = append(reports, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	safeReportsCount := 0
	dampenedSafeReportsCount := 0

	for _, report := range reports {
		if isSafe(report) {
			safeReportsCount++
			continue
		}

		canBecomeGood := false
		for j := 0; j < len(report); j++ {
			dampenedReport := make([]int, 0, len(report)-1)
			dampenedReport = append(dampenedReport, report[:j]...)
			dampenedReport = append(dampenedReport, report[j+1:]...)

			if isSafe(dampenedReport) {
				canBecomeGood = true
			}
		}

		if canBecomeGood {
			dampenedSafeReportsCount++
		}
	}

	fmt.Printf("Number of safe reports(Part 1): %d\n", safeReportsCount) // Part 1 Solution: 220
	fmt.Printf("Number of safe reports with dampening(Part 2): %d\n", safeReportsCount+dampenedSafeReportsCount) // Part 2 Solution: 296
}
