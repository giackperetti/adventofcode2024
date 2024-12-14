package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitStone(stone int) (int, int) {
	stoneString := strconv.Itoa(stone)
	stone1, stone2 := stoneString[:len(stoneString)/2], stoneString[len(stoneString)/2:]

	num1, _ := strconv.Atoi(stone1)
	num2, _ := strconv.Atoi(stone2)

	return num1, num2
}

func getStonesAfterBlink(stone int) []int {
	resultStones := []int{}

	switch {
	case stone == 0:
		resultStones = append(resultStones, 1)
	case len(strconv.Itoa(stone))%2 == 0:
		s1, s2 := splitStone(stone)
		resultStones = append(resultStones, s1, s2)
	default:
		resultStones = append(resultStones, stone*2024)
	}

	return resultStones
}

func getCountAfterBlinks(stone int, cache map[int][]int, blinkCount int) int {
	if _, ok := cache[stone]; ok {
		if cache[stone][blinkCount-1] != 0 {
			return cache[stone][blinkCount-1]
		}
	} else {
		cache[stone] = make([]int, 75)
	}

	if blinkCount == 1 {
		cache[stone][blinkCount-1] = len(getStonesAfterBlink(stone))
		return len(getStonesAfterBlink(stone))
	}

	sum := 0

	for _, stone := range getStonesAfterBlink(stone) {
		sum += getCountAfterBlinks(stone, cache, blinkCount-1)
	}

	cache[stone][blinkCount-1] = sum
	return sum
}

func getStoneCountAfterBlinking(input []int, timesBlink int) int {
	sum := 0
	cache := make(map[int][]int)
	for _, stone := range input {
		sum += getCountAfterBlinks(stone, cache, timesBlink)
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		} else {
			fmt.Println("Empty file")
		}
		return
	}

	var inputData []int
	for _, numStr := range strings.Fields(scanner.Text()) {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			inputData = append(inputData, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(inputData) == 0 {
		fmt.Println("No valid integers found in input")
		return
	}

	fmt.Printf("Part 1: %d\n", getStoneCountAfterBlinking(inputData, 25))
	fmt.Printf("Part 2: %d\n", getStoneCountAfterBlinking(inputData, 75))
}
