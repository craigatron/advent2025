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

	dialVal := 50
	zeroes := 0

	for _, line := range lines {
		direction := string(line[0])
		ticks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == "L" {
			ticks *= -1
		}
		dialVal += ticks
		dialVal = mod(dialVal, 100)

		if dialVal == 0 {
			zeroes += 1
		}
	}

	return zeroes
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	} else {
		return x
	}
}

func mod(a, b int) int {
	x := a % b
	if x < 0 {
		x += b
	}
	return x
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	dialVal := 50
	zeroes := 0

	for _, line := range lines {
		direction := string(line[0])
		ticks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == "L" {
			ticks *= -1
		}
		lastVal := dialVal
		dialVal += ticks

		zeroes += (abs(dialVal) / 100)
		if lastVal > 0 && dialVal <= 0 {
			zeroes += 1
		}

		dialVal = mod(dialVal, 100)
	}

	return zeroes
}
