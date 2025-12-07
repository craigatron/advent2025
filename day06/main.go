package main

import (
	_ "embed"
	"fmt"
	"strconv"
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

	problems := make([][]int, 0)
	var ops []string
	for _, line := range lines {
		parts := strings.Fields(line)

		if strings.Contains(line, "+") {
			ops = parts
		} else {
			for i, part := range parts {
				if i >= len(problems) {
					problems = append(problems, make([]int, 0))
				}
				num, _ := strconv.Atoi(part)
				problems[i] = append(problems[i], num)
			}
		}
	}

	total := 0
	for i, problem := range problems {
		op := ops[i]

		answer := problem[0]
		for i := 1; i < len(problem); i++ {
			if op == "+" {
				answer += problem[i]
			} else {
				answer *= problem[i]
			}
		}
		total += answer
	}

	return total
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	grid := make([][]string, 0)
	var ops []string

	for _, line := range lines {
		if strings.Contains(line, "+") {
			ops = strings.Fields(line)
		} else {
			grid = append(grid, strings.Split(line, ""))
		}
	}

	splits := make([]int, 0)
	for j := 0; j < len(grid[0]); j++ {
		isSplit := true
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != " " {
				isSplit = false
				break
			}
		}

		if isSplit {
			splits = append(splits, j)
		}
	}
	splits = append(splits, len(grid[0]))

	total := 0

	startIndex := 0
	for opIndex, split := range splits {
		answer := 0
		for j := split - 1; j >= startIndex; j-- {
			numStr := ""
			for i := 0; i < len(grid); i++ {
				if grid[i][j] != " " {
					numStr += grid[i][j]
				}
			}
			num, _ := strconv.Atoi(numStr)
			if answer == 0 {
				answer = num
			} else if ops[opIndex] == "+" {
				answer += num
			} else {
				answer *= num
			}
		}
		total += answer
		startIndex = split + 1
	}

	return total
}
