package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) InBounds(maxRows, maxCols int) bool {
	return p.x >= 0 && p.x < maxRows && p.y >= 0 && p.y < maxCols
}

func parseAntennaLocations(gridLines []string) map[rune][]Point {
	antennaLocations := make(map[rune][]Point)
	for rowIndex, line := range gridLines {
		for colIndex, char := range line {
			if char != '.' {
				antennaLocations[char] = append(antennaLocations[char], Point{x: rowIndex, y: colIndex})
			}
		}
	}
	return antennaLocations
}

func calculateAntinodePositions(antennaLocations map[rune][]Point, maxRows int, maxCols int, limitToFirstAntinode bool) []Point {
	antinodePositions := make([]Point, 0)
	for _, antennas := range antennaLocations {
		for i, antenna1 := range antennas {
			if !limitToFirstAntinode {
				antinodePositions = append(antinodePositions, antenna1)
			}
			for j, antenna2 := range antennas {
				if i <= j {
					continue
				}
				directionVector := Point{x: antenna2.x - antenna1.x, y: antenna2.y - antenna1.y}

				antinodeBeforeAntenna1 := Point{x: antenna1.x - directionVector.x, y: antenna1.y - directionVector.y}
				for antinodeBeforeAntenna1.InBounds(maxRows, maxCols) {
					antinodePositions = append(antinodePositions, antinodeBeforeAntenna1)
					if limitToFirstAntinode {
						break
					}
					antinodeBeforeAntenna1 = Point{x: antinodeBeforeAntenna1.x - directionVector.x, y: antinodeBeforeAntenna1.y - directionVector.y}
				}

				antinodeBeyondAntenna2 := Point{x: antenna2.x + directionVector.x, y: antenna2.y + directionVector.y}
				for antinodeBeyondAntenna2.InBounds(maxRows, maxCols) {
					antinodePositions = append(antinodePositions, antinodeBeyondAntenna2)
					if limitToFirstAntinode {
						break
					}
					antinodeBeyondAntenna2 = Point{x: antinodeBeyondAntenna2.x + directionVector.x, y: antinodeBeyondAntenna2.y + directionVector.y}
				}
			}
		}
	}
	return antinodePositions
}

func removeDuplicatePoints(points []Point) []Point {
	uniquePoints := make(map[Point]bool)
	distinctPoints := []Point{}
	for _, point := range points {
		if !uniquePoints[point] {
			uniquePoints[point] = true
			distinctPoints = append(distinctPoints, point)
		}
	}
	return distinctPoints
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	gridLines := strings.Split(strings.TrimSpace(string(file)), "\n")
	rowCount := len(gridLines)
	colCount := len(gridLines[0])
	antennaLocations := parseAntennaLocations(gridLines)

	antinodesWithCutoff := calculateAntinodePositions(antennaLocations, rowCount, colCount, true)
	totalUniqueAntinodesWithCutoff := len(removeDuplicatePoints(antinodesWithCutoff))
	fmt.Printf("Total unique antinode locations with cutoff condition: %d\n", totalUniqueAntinodesWithCutoff) // Part 1 Solution: 323

	antinodesWithoutCutoff := calculateAntinodePositions(antennaLocations, rowCount, colCount, false)
	totalUniqueAntinodesWithoutCutoff := len(removeDuplicatePoints(antinodesWithoutCutoff))
	fmt.Printf("Total unique antinode locations without cutoff condition: %d\n", totalUniqueAntinodesWithoutCutoff) // Part 2 Solution: 1077
}
