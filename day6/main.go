package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	DOWN Direction = iota
	UP
	LEFT
	RIGHT
)

type Position struct {
	x int
	y int
}

type Visit struct {
	position Position
	direction Direction
}

type LaboratoryMap [][]byte

const (
	OBSTACLE       byte = '#'
	START_POSITION byte = '^'
	EMPTY          byte = '.'
)

func (visit Visit) moveForward() Visit {
	switch visit.direction {
	case DOWN:
		return Visit{Position{visit.position.x, visit.position.y - 1}, visit.direction}
	case UP:
		return Visit{Position{visit.position.x, visit.position.y + 1}, visit.direction}
	case LEFT:
		return Visit{Position{visit.position.x - 1, visit.position.y}, visit.direction}
	case RIGHT:
		return Visit{Position{visit.position.x + 1, visit.position.y}, visit.direction}
	default:
		return visit
	}
}

func (visit Visit) rotateClockwise() Visit {
	switch visit.direction {
	case DOWN:
		return Visit{visit.position, RIGHT}
	case UP:
		return Visit{visit.position, LEFT}
	case LEFT:
		return Visit{visit.position, DOWN}
	case RIGHT:
		return Visit{visit.position, UP}
	default:
		return visit
	}
}

func (laboratoryMap LaboratoryMap) isPositionOutside(position Position) bool {
	return (position.x < 0 || position.x >= len(laboratoryMap[0])) || (position.y < 0 || position.y >= len(laboratoryMap))
}

func (laboratoryMap LaboratoryMap) takeStep(visit Visit) (Visit, bool) {
	nextVisit := visit.moveForward()
	nextPosition := nextVisit.position
	if laboratoryMap.isPositionOutside(nextPosition) {
		return Visit{}, true
	}
	for laboratoryMap[nextPosition.y][nextPosition.x] == OBSTACLE {
		nextVisit = visit.rotateClockwise()
		nextPosition = nextVisit.position
	}
	return nextVisit, false
}

func (laboratoryMap LaboratoryMap) findGuardStart() (Visit, error) {
	for y, row := range laboratoryMap {
		for x, cell := range row {
			if cell == START_POSITION {
				return Visit{Position{x, y}, DOWN}, nil
			}
		}
	}
	return Visit{}, fmt.Errorf("no start position found")
}

func (laboratoryMap LaboratoryMap) predictPatrolPath() (map[Position]bool, bool) {
	visitedPositions := map[Position]bool{}
	visits := map[Visit]bool{}
	visit, err := laboratoryMap.findGuardStart()
	if err != nil {
		fmt.Println(err)
		return nil, true
	}

	var exited bool
	for {
		if visits[visit] {
			break
		}
		visitedPositions[visit.position] = true
		visits[visit] = true
		visit, exited = laboratoryMap.takeStep(visit)
		if visit == (Visit{}) {
			break
		}
	}
	return visitedPositions, exited
}

func countPotentialLoops(laboratoryMap LaboratoryMap, visitedPositions map[Position]bool) int {
	loopCount := 0
	start, _ := laboratoryMap.findGuardStart()
	delete(visitedPositions, start.position)

	for position := range visitedPositions {
		laboratoryMapCopy := make(LaboratoryMap, len(laboratoryMap))
		for i := range laboratoryMap {
			laboratoryMapCopy[i] = make([]byte, len(laboratoryMap[i]))
			copy(laboratoryMapCopy[i], laboratoryMap[i])
		}

		laboratoryMapCopy[position.y][position.x] = OBSTACLE
		_, exited := laboratoryMapCopy.predictPatrolPath()
		if !exited {
			loopCount++
		}
	}
	return loopCount
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	laboratoryMap := LaboratoryMap{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) > 0 {
			laboratoryMap = append(laboratoryMap, append([]byte{}, line...))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	visitedPositions, exited := laboratoryMap.predictPatrolPath()
	if exited {
		fmt.Println("The guard left the laboratory map.")
	}
	fmt.Printf("Number of distinct positions visited by the guard: %d\n", len(visitedPositions))

	loopCount := countPotentialLoops(laboratoryMap, visitedPositions)
	fmt.Printf("Number of positions where loops can form: %d\n", loopCount)
}