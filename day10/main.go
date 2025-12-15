package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	// cheating a little oh well
	"gonum.org/v1/gonum/stat/combin"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("running part 1")
	fmt.Println("part 1 output:", part1(input))
	fmt.Println("running part 2")
	fmt.Println("part 2 output:", part2(input))
}

type ProblemRow struct {
	targetLights []bool
	buttons      [][]int
	joltages     []int
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	problems := make([]ProblemRow, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		targetLights := make([]bool, 0)
		for i := 1; i < len(parts[0])-1; i++ {
			targetLights = append(targetLights, string(parts[0][i]) == "#")
		}
		buttons := make([][]int, 0)
		for i := 1; i < len(parts)-1; i++ {
			buttonLights := strings.Split(parts[i][1:len(parts[i])-1], ",")
			buttonConfig := make([]int, 0)
			for _, light := range buttonLights {
				lightNum, _ := strconv.Atoi(light)
				buttonConfig = append(buttonConfig, lightNum)
			}

			buttons = append(buttons, buttonConfig)
		}

		problems = append(problems, ProblemRow{targetLights, buttons, nil})
	}

	total := 0
	for _, problem := range problems {
		queue := make([][]int, 0)
		for i := range problem.buttons {
			queue = append(queue, []int{i})
		}

		for {
			candidate := queue[0]
			queue = queue[1:]

			lights := make([]bool, 0)
			for range problem.targetLights {
				lights = append(lights, false)
			}

			for _, button := range candidate {
				for _, lightIndex := range problem.buttons[button] {
					lights[lightIndex] = !lights[lightIndex]
				}
			}
			if slices.Equal(lights, problem.targetLights) {
				total += len(candidate)
				break
			}
			for i := range problem.buttons {
				if i == candidate[len(candidate)-1] {
					continue
				}
				newCandidates := slices.Clone(candidate)
				newCandidates = append(newCandidates, i)
				queue = append(queue, newCandidates)
			}
		}
	}

	return total
}

func patterns(buttons [][]int) map[string]int {
	costs := make(map[string]int)

	for pattern_len := 0; pattern_len <= len(buttons); pattern_len++ {
		for _, combo := range combin.Combinations(len(buttons), pattern_len) {
			pattern := make([]int, 0)
			for range buttons[0] {
				pattern = append(pattern, 0)
			}
			for _, buttonIndex := range combo {
				button := buttons[buttonIndex]
				for i := range button {
					pattern[i] += button[i]
				}
			}

			key := patternKey(pattern)

			_, ok := costs[key]
			if !ok {
				costs[key] = pattern_len
			}
		}
	}

	return costs
}

func patternKey(pattern []int) string {
	strs := make([]string, 0)
	for _, val := range pattern {
		strs = append(strs, strconv.Itoa(val))
	}
	return strings.Join(strs, ",")
}

func reverseKey(pattern string) []int {
	ints := make([]int, 0)
	parts := strings.Split(pattern, ",")
	for _, p := range parts {
		val, _ := strconv.Atoi(p)
		ints = append(ints, val)
	}
	return ints
}

func solve(costs map[string]int, buttons [][]int, joltages []int) int {
	allZero := true
	for _, j := range joltages {
		if j != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}
	result := 100000000
	for k, v := range costs {
		pattern := reverseKey(k)
		allEven := true
		for i := range pattern {
			if pattern[i] > joltages[i] || pattern[i]%2 != joltages[i]%2 {
				allEven = false
				break
			}
		}
		if allEven {
			newJoltages := make([]int, 0)
			for i := range pattern {
				newJoltages = append(newJoltages, (joltages[i]-pattern[i])/2)
			}
			result = min(result, v+2*solve(costs, buttons, newJoltages))
		}
	}
	return result
}

// shamefully stolen from https://old.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
// after a full day of trying to write a linear algebra solver
func part2(input string) int {
	lines := strings.Split(input, "\n")

	problems := make([]ProblemRow, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")

		joltagesStr := parts[len(parts)-1]
		joltagesStr = joltagesStr[1 : len(joltagesStr)-1]
		joltageParts := strings.Split(joltagesStr, ",")
		joltages := make([]int, 0)
		for _, joltage := range joltageParts {
			joltageNum, _ := strconv.Atoi(joltage)
			joltages = append(joltages, joltageNum)
		}

		buttons := make([][]int, 0)
		for i := 1; i < len(parts)-1; i++ {
			buttonConfig := make([]int, 0)
			for range joltages {
				buttonConfig = append(buttonConfig, 0)
			}
			for light := range strings.SplitSeq(parts[i][1:len(parts[i])-1], ",") {
				lightNum, _ := strconv.Atoi(light)
				buttonConfig[lightNum] = 1
			}

			buttons = append(buttons, buttonConfig)
		}

		problems = append(problems, ProblemRow{nil, buttons, joltages})
	}

	total := 0

	for _, problem := range problems {
		costs := patterns(problem.buttons)
		result := solve(costs, problem.buttons, problem.joltages)
		total += result
	}

	return total
}
