package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	controlExpression := regexp.MustCompile(`mul\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)|do\(\)|don't\(\)`)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var allMultiplications []struct{ x, y int }
	var enabledMultiplications []struct{ x, y int }

	shouldBeMultiplied := true
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := controlExpression.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fullMatch := match[0]

			if fullMatch == "do()" {
				shouldBeMultiplied = true
			} else if fullMatch == "don't()" {
				shouldBeMultiplied = false
			} else {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])

				allMultiplications = append(allMultiplications, struct{ x, y int }{x, y})

				if shouldBeMultiplied {
					enabledMultiplications = append(enabledMultiplications, struct{ x, y int }{x, y})
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	multiplicationsSumAll := 0
	for _, pair := range allMultiplications {
		multiplicationsSumAll += pair.x * pair.y
	}
	fmt.Printf("Sum of all multiplications (Part 1): %d\n", multiplicationsSumAll) // Part 1 Solution: 174561379

	multiplicationsSumEnabled := 0
	for _, pair := range enabledMultiplications {
		multiplicationsSumEnabled += pair.x * pair.y
	}
	fmt.Printf("Sum of enabled multiplications (Part 2): %d\n", multiplicationsSumEnabled) // Part 2 Solution: 106921067
}
