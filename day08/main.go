package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"math"
	"slices"
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

type Points struct {
	x    int
	y    int
	dist float64
}

const numCircuits = 1000

func part1(input string) int {
	lines := strings.Split(input, "\n")

	points := make([][3]float64, 0, numCircuits)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, [3]float64{float64(x), float64(y), float64(z)})
	}
	pairs := make([]Points, 0, numCircuits)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			dist := math.Sqrt(
				math.Pow(points[i][0]-points[j][0], 2) + math.Pow(points[i][1]-points[j][1], 2) + math.Pow(points[i][2]-points[j][2], 2))
			if len(pairs) < numCircuits || dist < pairs[numCircuits-1].dist {
				pairs = append(pairs, Points{i, j, dist})
				slices.SortFunc(pairs, func(a Points, b Points) int {
					return cmp.Compare(a.dist, b.dist)
				})
				if len(pairs) > numCircuits {
					pairs = pairs[0:numCircuits]
				}
			}
		}
	}

	circuits := map[int]int{}
	for i, pair := range pairs {
		xCircuit, xOk := circuits[pair.x]
		yCircuit, yOk := circuits[pair.y]
		if xOk && yOk {
			// both already in map
			for k, v := range circuits {
				if v == yCircuit {
					circuits[k] = xCircuit
				}
			}
		} else if !xOk && yOk {
			// only y in map
			circuits[pair.x] = yCircuit
		} else if xOk && !yOk {
			// only x in map
			circuits[pair.y] = xCircuit
		} else {
			// neither in map
			circuits[pair.x] = i
			circuits[pair.y] = i
		}
	}

	countMap := map[int]int{}
	for _, v := range circuits {
		_, ok := countMap[v]
		if ok {
			countMap[v] += 1
		} else {
			countMap[v] = 1
		}
	}
	counts := make([]int, 0)
	for _, v := range countMap {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	return counts[0] * counts[1] * counts[2]
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	points := make([][3]int, 0, numCircuits)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, [3]int{x, y, z})
	}
	pairs := make([]Points, 0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			dist := math.Sqrt(
				math.Pow(float64(points[i][0]-points[j][0]), 2) + math.Pow(float64(points[i][1]-points[j][1]), 2) + math.Pow(float64(points[i][2]-points[j][2]), 2))
			pairs = append(pairs, Points{i, j, dist})
		}
	}
	slices.SortFunc(pairs, func(a Points, b Points) int {
		return cmp.Compare(a.dist, b.dist)
	})

	circuits := map[int]int{}
	for i, pair := range pairs {
		xCircuit, xOk := circuits[pair.x]
		yCircuit, yOk := circuits[pair.y]
		if xOk && yOk {
			// both already in map
			for k, v := range circuits {
				if v == yCircuit {
					circuits[k] = xCircuit
				}
			}
		} else if !xOk && yOk {
			// only y in map
			circuits[pair.x] = yCircuit
		} else if xOk && !yOk {
			// only x in map
			circuits[pair.y] = xCircuit
		} else {
			// neither in map
			circuits[pair.x] = i
			circuits[pair.y] = i
		}

		if len(circuits) == len(points) {
			curVal := -1
			allSame := true
			for _, v := range circuits {
				if curVal == -1 {
					curVal = v
				} else if curVal != v {
					allSame = false
					break
				}
			}
			if allSame {
				return points[pair.x][0] * points[pair.y][0]
			}
		}
	}

	return -1
}
