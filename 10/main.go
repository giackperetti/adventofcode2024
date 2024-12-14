package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type InputData = []string
type SolutionOutput = int

type Point struct {
	x, y int
}

func findStarts(m map[Point]int, goal int) []Point {
	st := []Point{}
	for k, v := range m {
		if v == goal {
			st = append(st, k)
		}
	}
	return st
}

func getMoves(m map[Point]int, cur Point, v int) []Point {
	dirs := []Point{{cur.x + 1, cur.y}, {cur.x - 1, cur.y}, {cur.x, cur.y + 1}, {cur.x, cur.y - 1}}
	valid := []Point{}
	for _, dir := range dirs {
		if c, ok := m[dir]; ok && c == v+1 {
			valid = append(valid, dir)
		}
	}
	return valid
}

func countPaths(m map[Point]int, start Point, goal int, distinct bool) int {
	np := 0
	v := make(map[Point]bool)
	queue := []Point{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if _, ok := v[cur]; ok && distinct {
			continue
		}
		v[cur] = true

		if m[cur] == goal {
			np++
			continue
		}

		mvs := getMoves(m, cur, m[cur])

		queue = append(queue, mvs...)
	}

	return np
}

func findAllPaths(m map[Point]int, start, goal int, distinct bool) int {
	sum := 0
	starts := findStarts(m, start)
	for _, pt := range starts {
		score := countPaths(m, pt, goal, distinct)
		sum += score
	}
	return sum
}

func parse(in InputData) map[Point]int {
	res := make(map[Point]int)
	for y, line := range in {
		for x, c := range line {
			v, _ := strconv.Atoi(string(c))
			res[Point{x, y}] = v
		}
	}
	return res
}

func solve(in InputData, distinct bool) SolutionOutput {
	m := parse(in)
	return findAllPaths(m, 0, 9, distinct)
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var inputData []string
	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			inputData = append(inputData, line)
		}
	}

	part1 := solve(inputData, true)
	part2 := solve(inputData, false)

	fmt.Println("Part 1:", part1) // Part 1 Solution: 548
	fmt.Println("Part 2:", part2) // Part 2 Solution: 1252
}
