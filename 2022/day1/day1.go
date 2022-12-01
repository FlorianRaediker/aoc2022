package main

import (
	"strings"

	"aoc2022/util"
)

func main() {
	util.Run(2022, 1, parseInput, part1, part2)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n\n")
}

func part1(input []string) int {
	max := 0
	for _, elf := range input {
		sum := 0
		for _, calories := range util.StringSliceToInt(strings.Split(elf, "\n")) {
			sum += calories
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func part2(input []string) int {
	max1 := 0
	max2 := 0
	max3 := 0
	for _, elf := range input {
		sum := 0
		for _, calories := range util.StringSliceToInt(strings.Split(elf, "\n")) {
			sum += calories
		}
		if sum > max3 {
			max3 = sum
			if max3 > max2 {
				max3, max2 = max2, max3
				if max2 > max1 {
					max2, max1 = max1, max2
				}
			}
		}
	}
	return max1 + max2 + max3
}
