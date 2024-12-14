package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MiddleElement(arr []int) int {
	if len(arr) > 0 {
		middleIndex := len(arr) / 2
		return arr[middleIndex]
	}

	return 0
}

func isValidOrder(sequence []int, orderRules map[int][]int) bool {
	pagePosition := make(map[int]int)
	for index, page := range sequence {
		pagePosition[page] = index
	}

	for precedingPage, followingPages := range orderRules {
		for _, followingPage := range followingPages {
			precedingPosition, precedingExists := pagePosition[precedingPage]
			followingPosition, followingExists := pagePosition[followingPage]

			if precedingExists && followingExists && (precedingPosition > followingPosition) {
				return false
			}
		}
	}
	return true
}

func correctOrder(sequence []int, orderRules map[int][]int) []int {
	correctedSequence := make([]int, len(sequence))
	copy(correctedSequence, sequence)

	for {
		pagePosition := make(map[int]int)
		for index, page := range correctedSequence {
			pagePosition[page] = index
		}

		corrected := false

		for precedingPage, followingPages := range orderRules {
			for _, followingPage := range followingPages {
				precedingPosition, precedingExists := pagePosition[precedingPage]
				followingPosition, followingExists := pagePosition[followingPage]

				if precedingExists && followingExists && (precedingPosition > followingPosition) {
					newSequence := make([]int, 0, len(correctedSequence))

					newSequence = append(newSequence, correctedSequence[:followingPosition]...)
					newSequence = append(newSequence, correctedSequence[followingPosition+1:]...)

					insertPosition := precedingPosition + 1
					if insertPosition > len(newSequence) {
						insertPosition = len(newSequence)
					}

					newSequence = append(newSequence[:insertPosition], append([]int{followingPage}, newSequence[insertPosition:]...)...)

					correctedSequence = newSequence
					corrected = true

					break
				}
			}

			if corrected {
				break
			}
		}

		if !corrected {
			break
		}
	}

	return correctedSequence
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	printerUpdatesOrder := make(map[int][]int)
	var printerUpdates [][]int
	var correctlyOrderedPrinterUpdates [][]int
	var incorrectlyOrderedPrinterUpdates [][]int
	var correctedIncorrectlyOrderedPrinterUpdates [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			a, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			b, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			printerUpdatesOrder[a] = append(printerUpdatesOrder[a], b)
		} else if strings.Contains(line, ",") {
			nums := strings.Split(line, ",")
			var sequence []int
			for _, num := range nums {
				val, _ := strconv.Atoi(strings.TrimSpace(num))
				sequence = append(sequence, val)
			}
			printerUpdates = append(printerUpdates, sequence)
		}
	}

	for _, update := range printerUpdates {
		if isValidOrder(update, printerUpdatesOrder) {
			correctlyOrderedPrinterUpdates = append(correctlyOrderedPrinterUpdates, update)
		} else {
			incorrectlyOrderedPrinterUpdates = append(incorrectlyOrderedPrinterUpdates, update)
		}
	}

	for _, incorrectUpdate := range incorrectlyOrderedPrinterUpdates {
		correctedIncorrectlyOrderedPrinterUpdates = append(correctedIncorrectlyOrderedPrinterUpdates, correctOrder(incorrectUpdate, printerUpdatesOrder))
	}

	correctMiddlesCount := 0
	for _, correctUpdate := range correctlyOrderedPrinterUpdates {
		correctMiddlesCount += MiddleElement(correctUpdate)
	}

	correctedIncorrectMiddlesCount := 0
	for _, correctedIncorrectUpdate := range correctedIncorrectlyOrderedPrinterUpdates {
		correctedIncorrectMiddlesCount += MiddleElement(correctedIncorrectUpdate)
	}

	fmt.Printf("Sum of middle numbers of already valid updates: %d\n", correctMiddlesCount)                // Part 1 Solution: 4814
	fmt.Printf("Sum of middle numbers of corrected invalid updates: %d\n", correctedIncorrectMiddlesCount) // Part 2 Solution: 5448
}
