package main

import (
	_ "embed"
	"fmt"
	"sort"
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
	allRanges := make([][]int, 0)
	ingredients := make([]int, 0)
	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			allRanges = append(allRanges, []int{start, end})
		} else if len(line) > 0 {
			ingredient, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredient)
		}
	}

	freshTotal := 0

	for _, ingredient := range ingredients {
		fresh := false
		for _, r := range allRanges {
			if ingredient >= r[0] && ingredient <= r[1] {
				fresh = true
				break
			}
		}
		if fresh {
			freshTotal += 1
		}
	}

	return freshTotal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	allRanges := make([][]int, 0)
	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			allRanges = append(allRanges, []int{start, end})
		}
	}
	sort.Slice(allRanges, func(i, j int) bool {
		return allRanges[i][0] < allRanges[j][0]
	})

	combinedRanges := make([][]int, 0)
	var curRange []int
	for _, r := range allRanges {
		if curRange == nil {
			curRange = r
		} else if r[0] <= curRange[1] {
			curRange[1] = max(curRange[1], r[1])
		} else {
			combinedRanges = append(combinedRanges, curRange)
			curRange = r
		}
	}
	combinedRanges = append(combinedRanges, curRange)

	total := 0

	for _, r := range combinedRanges {
		total += (r[1] - r[0] + 1)
	}

	return total
}
