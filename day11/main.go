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

	nodes := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		n1 := parts[0]
		if _, ok := nodes[n1]; !ok {
			nodes[n1] = make([]string, 0)
		}
		for node := range strings.SplitSeq(parts[1], " ") {
			nodes[n1] = append(nodes[n1], node)
		}
	}

	queue := make([][]string, 0)
	queue = append(queue, []string{"you"})

	paths := 0

	for len(queue) != 0 {
		var state []string
		state, queue = queue[0], queue[1:]

		for _, nextNode := range nodes[state[len(state)-1]] {
			if nextNode == "out" {
				paths += 1
			} else {
				newState := slices.Clone(state)
				newState = append(newState, nextNode)
				queue = append(queue, newState)
			}
		}
	}

	return paths
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	pathCounts := make(map[string]map[string]int)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		n1 := parts[0]
		if _, ok := pathCounts[n1]; !ok {
			pathCounts[n1] = make(map[string]int)
		}
		for node := range strings.SplitSeq(parts[1], " ") {
			pathCounts[n1][node] = 1
		}
	}

	for len(pathCounts) > 3 {
		var nodeKey string
		for k := range pathCounts {
			if k != "svr" && k != "fft" && k != "dac" {
				nodeKey = k
				break
			}
		}
		nodePaths := pathCounts[nodeKey]
		delete(pathCounts, nodeKey)

		for _, v := range pathCounts {
			if ct, ok := v[nodeKey]; ok {
				for pk, pv := range nodePaths {
					if _, ok := v[pk]; !ok {
						v[pk] = 0
					}
					v[pk] += ct * pv
				}
				delete(v, nodeKey)
			}
		}
	}

	var firstKey string
	var secondKey string
	if len(pathCounts["dac"]) == 2 {
		firstKey = "dac"
		secondKey = "fft"
	} else {
		firstKey = "fft"
		secondKey = "dac"
	}

	return pathCounts["svr"][firstKey] * pathCounts[firstKey][secondKey] * pathCounts[secondKey]["out"]
}
