package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(line []string) {
	solution := int64(0)
	diskMap := make([]string, 0)

	for i := range line {
		if i%2 == 0 {
			id := i / 2
			fileSpace, _ := strconv.ParseInt(line[i], 10, 32)
			for j := 0; j < int(fileSpace); j++ {
				diskMap = append(diskMap, fmt.Sprintf("%d", id))
			}
		} else {
			freeSpace, _ := strconv.ParseInt(line[i], 10, 32)
			for j := 0; j < int(freeSpace); j++ {
				diskMap = append(diskMap, ".")
			}
		}
	}

	for i, j := 0, len(diskMap)-1; i <= j; {
		if diskMap[i] == "." && diskMap[j] != "." {
			diskMap[i], diskMap[j] = diskMap[j], diskMap[i]
			i++
			j--
		} else {
			if diskMap[i] != "." {
				i++
			} else if diskMap[j] != "." {
				j--
			} else if diskMap[i] == "." && diskMap[j] == "." {
				j--
			}
		}
	}

	for i := range diskMap {
		if diskMap[i] != "." {
			x, _ := strconv.ParseInt(diskMap[i], 10, 32)
			solution += int64(i) * x
		}
	}

	fmt.Printf("Part One: %d\n", solution)
}

func partTwo(line []string) {
	solution := int64(0)
	diskMap := make([]string, 0)
	files := make([][]int64, 0)
	spaces := make([][]int64, 0)

	for i := range line {
		if i%2 == 0 {
			id := i / 2
			fileSpace, _ := strconv.ParseInt(line[i], 10, 32)
			files = append(files, []int64{int64(len(diskMap)), int64(len(diskMap)-1) + fileSpace})
			for j := 0; j < int(fileSpace); j++ {
				diskMap = append(diskMap, fmt.Sprintf("%d", id))
			}
		} else {
			space, _ := strconv.ParseInt(line[i], 10, 32)
			spaces = append(spaces, []int64{int64(len(diskMap)), int64(len(diskMap)-1) + space})
			for j := 0; j < int(space); j++ {
				diskMap = append(diskMap, ".")
			}
		}
	}

	for j := len(files) - 1; j >= 0; j-- {
		for i := 0; i < len(spaces); i++ {
			space := spaces[i]
			if space[1] < files[j][1] && space[1]-space[0] >= files[j][1]-files[j][0] {
				for k := space[0]; k <= space[0]+files[j][1]-files[j][0]; k++ {
					diskMap[k] = fmt.Sprintf("%d", j)
				}
				for k := files[j][0]; k <= files[j][1]; k++ {
					diskMap[k] = "."
				}
				space[0] = space[0] + files[j][1] - files[j][0] + 1
				break
			}
		}

	}

	for i := range diskMap {
		if diskMap[i] != "." {
			x, _ := strconv.ParseInt(diskMap[i], 10, 32)
			solution += int64(i) * x
		}
	}

	fmt.Printf("Part Two: %d\n", solution)
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := (strings.Split(string(file), "\n"))

	partOne(strings.Split(lines[0], "")) // Part 1 Solution: 6301895872542
	partTwo(strings.Split(lines[0], "")) // Part 2 Solution: 6323761685944

}
