package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("running part 1")
	fmt.Println("part 1 output:", part1(input))
	fmt.Println("running part 2")
	fmt.Println("part 2 output:", part2(input))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	cols := make([]int, 0)
	splits := 0
	for i, line := range lines {
		chars := strings.Split(line, "")
		if i == 0 {
			cols = append(cols, slices.Index(chars, "S"))
			continue
		}
		newCols := make([]int, 0)
		for j, c := range chars {
			if c == "^" && slices.Contains(cols, j) {
				cols = slices.DeleteFunc(cols, func(x int) bool {
					return x == j
				})
				newCols = append(newCols, j-1, j+1)
				splits += 1

				if len(cols) == 0 {
					break
				}
			}
		}
		newCols = append(newCols, cols...)
		slices.Sort(newCols)
		cols = slices.Compact(newCols)
	}

	return splits
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	grid := make([][]string, 0)

	var startCol int
	for i, line := range lines {
		chars := strings.Split(line, "")
		if i == 0 {
			startCol = slices.Index(chars, "S")
		} else {
			grid = append(grid, strings.Split(line, ""))
		}
	}

	paths := map[[2]int]int{}
	for i := 0; i < len(lines[0]); i++ {
		paths[[2]int{len(lines) - 1, i}] = 1
	}
	for i := len(lines) - 2; i >= 0; i-- {
		for j, c := range grid[i] {
			if c == "^" {
				paths[[2]int{i, j}] = paths[[2]int{i + 1, j - 1}] + paths[[2]int{i + 1, j + 1}]
			} else {
				paths[[2]int{i, j}] = paths[[2]int{i + 1, j}]
			}
		}
	}

	return paths[[2]int{0, startCol}]
}
