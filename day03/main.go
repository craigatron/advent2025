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

	total := 0
	for _, line := range lines {
		max := 0

		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				val, _ := strconv.Atoi(string(line[i]) + string(line[j]))
				if val > max {
					max = val
				}
			}
		}
		total += max
	}

	return total
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		strval := ""
		startI := 0
		endI := len(line) - 12
		for {
			maxI := -1
			maxVal := -1
			// find the furthest left highest number in the line
			for i := endI; i >= startI; i -= 1 {
				candidateVal, _ := strconv.Atoi(string(line[i]))
				if candidateVal >= maxVal {
					maxVal = candidateVal
					maxI = i
				}
			}
			strval += string(line[maxI])
			startI = maxI + 1
			endI = len(line) - 12 + len(strval)

			if len(strval) == 12 {
				break
			}
		}
		val, _ := strconv.Atoi(strval)
		total += val
	}

	return total
}
