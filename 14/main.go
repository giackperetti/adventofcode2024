package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	x  int
	y  int
	dx int
	dy int
}

func convString2Int(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func simulateMovement(r Robot, t int, mx int, my int) Robot {
	newX := (((r.x + (r.dx * t)) % mx) + mx) % mx
	newY := (((r.y + (r.dy * t)) % my) + my) % my
	nRobot := Robot{x: newX, y: newY, dx: 0, dy: 0}
	return nRobot
}

func findQuadrant(r Robot, maxVals []int) int {
	if r.x == maxVals[0]/2 || r.y == maxVals[1]/2 {
		return 4
	}
	if r.x > maxVals[0]/2 {
		if r.y > maxVals[1]/2 {
			return 3
		} else {
			return 1
		}
	} else {
		if r.y > maxVals[1]/2 {
			return 2
		} else {
			return 0
		}
	}
}

func part1(robots []Robot, maxVals []int) int {
	var updatedRobots []Robot

	quadrantCounts := []int{0, 0, 0, 0, 0}
	for i, r := range robots {
		updatedRobots = append(updatedRobots, simulateMovement(r, 100, maxVals[0], maxVals[1]))
		quadrantCounts[findQuadrant(updatedRobots[i], maxVals)] += 1
	}

	count := 1
	for i := 0; i < 4; i++ {
		count *= quadrantCounts[i]
	}

	return count
}

func part2(robots []Robot, maxVals []int) int {
	minSafetyAndTimestep := []int{215987200, 100}

	for i := 1; i < 10000; i++ {
		var updatedRobots []Robot

		quadrantCounts := []int{0, 0, 0, 0, 0}
		for j, r := range robots {
			updatedRobots = append(updatedRobots, simulateMovement(r, i, maxVals[0], maxVals[1]))
			quadrantCounts[findQuadrant(updatedRobots[j], maxVals)] += 1
		}

		count := 1
		for j := 0; j < 4; j++ {
			count *= quadrantCounts[j]
		}

		if count < minSafetyAndTimestep[0] {
			minSafetyAndTimestep[0] = count
			minSafetyAndTimestep[1] = i
		}
	}

	return minSafetyAndTimestep[1]
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var robots []Robot
	maxGridValues := []int{101, 103}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		valPos := strings.Split(line[0][2:], ",")
		valDelt := strings.Split(line[1][2:], ",")

		nRobot := Robot{x: convString2Int(valPos[0]), y: convString2Int(valPos[1]), dx: convString2Int(valDelt[0]), dy: convString2Int(valDelt[1])}
		robots = append(robots, nRobot)
	}

	part1 := part1(robots, maxGridValues)
	part2 := part2(robots, maxGridValues)

	fmt.Println("Part 1: ", part1) // Part 1 Solution: 225521010
	fmt.Println("Part 2: ", part2) // Part 2 Solution: 7774
}
