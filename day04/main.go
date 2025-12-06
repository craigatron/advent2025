package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var adjDxDy = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func main() {
	fmt.Println("running part 1")
	fmt.Println("part 1 output:", part1(input))
	fmt.Println("running part 2")
	fmt.Println("part 2 output:", part2(input))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	size := len(lines)

	// hope it's always a square!
	grid := make([][]bool, size)
	for i := range size {
		grid[i] = make([]bool, size)
		line := lines[i]
		for j := range size {
			grid[i][j] = string(line[j]) == "@"
		}
	}

	total := 0
	for i := range size {
		for j := range size {
			if !grid[i][j] {
				continue
			}
			adjPaper := 0
			for _, dxdy := range adjDxDy {
				adjI := i + dxdy[0]
				adjJ := j + dxdy[1]
				if adjI < 0 || adjI >= size || adjJ < 0 || adjJ >= size {
					continue
				}
				if grid[adjI][adjJ] {
					adjPaper += 1
					if adjPaper >= 4 {
						break
					}
				}
			}
			if adjPaper < 4 {
				total += 1
			}
		}
	}

	return total
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	size := len(lines)

	// hope it's always a square!
	grid := make([][]bool, size)
	for i := range size {
		grid[i] = make([]bool, size)
		line := lines[i]
		for j := range size {
			grid[i][j] = string(line[j]) == "@"
		}
	}

	rollsRemoved := 0
	for {
		newRollsRemoved := 0
		for i := range size {
			for j := range size {
				if !grid[i][j] {
					continue
				}
				adjPaper := 0
				for _, dxdy := range adjDxDy {
					adjI := i + dxdy[0]
					adjJ := j + dxdy[1]
					if adjI < 0 || adjI >= size || adjJ < 0 || adjJ >= size {
						continue
					}
					if grid[adjI][adjJ] {
						adjPaper += 1
						if adjPaper >= 4 {
							break
						}
					}
				}
				if adjPaper < 4 {
					newRollsRemoved += 1
					grid[i][j] = false
				}
			}
		}
		if newRollsRemoved == 0 {
			break
		}
		rollsRemoved += newRollsRemoved
	}
	return rollsRemoved
}
