package main

import (
	"aoc2022/util"
)

func main() {
	util.Run(2022, 6, parseInput, part1, part2)
}

func parseInput(input string) string {
	return input
}

func part1(input string) any {
	for i := 4; i < len(input); i++ {
		marker := input[i-4 : i]
		different := true
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if j != k && marker[j] == marker[k] {
					different = false
					break
				}
			}
		}
		if different {
			return i
		}
	}
	return -1
}

func part2(input string) any {
	for i := 14; i < len(input); i++ {
		marker := input[i-14 : i]
		different := true
		for j := 0; j < 14; j++ {
			for k := 0; k < 14; k++ {
				if j != k && marker[j] == marker[k] {
					different = false
					break
				}
			}
		}
		if different {
			return i
		}
	}
	return -1
}
