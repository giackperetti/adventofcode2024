package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumList(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    var col1, col2 []int
    var distances []int
    var totalDistance int

    columnSimilarities := make(map[int]int)
    var similarityScores []int
    var totalSimilarityScore int

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)

        num1, err := strconv.Atoi(parts[0])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }
        num2, err := strconv.Atoi(parts[1])
        if err != nil {
            fmt.Println("Error converting string to int:", err)
            return
        }

        col1 = append(col1, num1)
		col2 = append(col2, num2)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading the file:", err)
        return
    }

    sort.Ints(col1)
    sort.Ints(col2)

    for i := 0; i < len(col1); i++ {
        distances = append(distances, int(math.Abs(float64(col1[i]-col2[i]))))
    }

    totalDistance = sumList(distances)
    fmt.Printf("Total Distance: %d\n", totalDistance)
    // Part 1 Solution: 2285373

    for _, num := range col2 {
        columnSimilarities[num]++
    }

    for i := 0; i < len(col1); i++ {
        similarityScores = append(similarityScores, col1[i] * columnSimilarities[col1[i]])
    }

    totalSimilarityScore = sumList(similarityScores)
    fmt.Printf("Total Similarity Score: %d\n", totalSimilarityScore)
    // Part 2 Solution: 21142653
}