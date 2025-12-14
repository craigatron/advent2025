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

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	points := make([][2]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, [2]int{x, y})
	}

	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			area := (abs(a[0]-b[0]) + 1) * (abs(a[1]-b[1]) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

type Point struct {
	x int
	y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	points := make([]Point, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x, y})
	}

	edges := make([][2]Point, 0)
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		j := i + 1
		if j == len(points) {
			j = 0
		}
		p2 := points[j]
		edges = append(edges, [2]Point{p1, p2})
	}

	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]

			x1 := min(a.x, b.x)
			x2 := max(a.x, b.x)
			y1 := min(a.y, b.y)
			y2 := max(a.y, b.y)

			// assuming the internal tiles are at least 2 across in each dimension
			if y2-y1 < 3 || x2-x1 < 3 {
				continue
			}

			area := (x2 - x1 + 1) * (y2 - y1 + 1)
			if area <= maxArea {
				continue
			}

			innerX1 := x1 + 1
			innerX2 := x2 - 1
			innerY1 := y1 + 1
			innerY2 := y2 - 1

			crosses := false
			for _, edge := range edges {
				horizontal := edge[0].y == edge[1].y
				if horizontal {
					minX := min(edge[0].x, edge[1].x)
					maxX := max(edge[0].x, edge[1].x)
					if edge[0].y >= innerY1 && edge[0].y <= innerY2 && ((minX <= innerX1 && maxX >= innerX1) || (minX <= innerX2 && maxX >= innerX2)) {
						crosses = true
						break
					}
				} else {
					minY := min(edge[0].y, edge[1].y)
					maxY := max(edge[0].y, edge[1].y)
					if edge[0].x >= innerX1 && edge[0].x <= innerX2 && ((minY <= innerY1 && maxY >= innerY1) || (minY <= innerY2 && maxY >= innerY2)) {
						crosses = true
						break
					}
				}
			}

			if !crosses {
				maxArea = area
				fmt.Println("new max:", maxArea, "points:", a, b)
			}
		}
	}

	return maxArea
}
