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

func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}
	res := x
	for i := 2; i <= y; i++ {
		res *= x
	}
	return res
}

func part1(input string) int {
	ranges := strings.Split(input, ",")

	total := 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		startLen := len(parts[0])
		endLen := len(parts[1])
		if startLen%2 != 0 && startLen == endLen {
			// not possible to have any invalid IDs in this range
			continue
		}

		var firstHalf int
		if startLen%2 == 0 {
			firstHalf, _ = strconv.Atoi(parts[0][0 : startLen/2])
		} else {
			firstHalf = pow(10, (len(parts[0]) / 2))
		}

		for {
			halfStr := strconv.Itoa(firstHalf)
			candidate, _ := strconv.Atoi(halfStr + halfStr)
			if candidate > end {
				break
			}
			if candidate >= start {
				total += candidate
			}
			firstHalf += 1
		}
	}

	return total
}

func part2(input string) int {
	ranges := strings.Split(input, ",")

	total := 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		endLen := len(parts[1])

		seenInvalid := map[int]bool{}

		repeatI := 1
		var candidateI int
		for {
			repeatStr := strconv.Itoa(repeatI)
			if len(repeatStr) > endLen/2 {
				break
			}

			candidateStr := repeatStr + repeatStr
			for {
				candidateI, _ = strconv.Atoi(candidateStr)
				if candidateI > end {
					break
				}
				if candidateI >= start {
					_, seen := seenInvalid[candidateI]
					if !seen {
						seenInvalid[candidateI] = true
						total += candidateI
					}
				}
				candidateStr += repeatStr
			}
			repeatI += 1
		}
	}

	return total
}
