package main

import (
	"aoc2022/util"
	"strings"
)

func main() {
	util.Run(2022, 2, parseInput, part1, part2)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

// A/X defeats C/Z
// B/Y defeats A/X
// C/Z defeats B/Y

var GAMES = map[string]struct {
	part1 int
	part2 int
}{
	"A X": {1 + 3, 0 + 3},
	"A Y": {2 + 6, 3 + 1},
	"A Z": {3 + 0, 6 + 2},

	"B X": {1 + 0, 0 + 1},
	"B Y": {2 + 3, 3 + 2},
	"B Z": {3 + 6, 6 + 3},

	"C X": {1 + 6, 0 + 2},
	"C Y": {2 + 0, 3 + 3},
	"C Z": {3 + 3, 6 + 1},
}

func part1(input []string) any {
	score := 0
	for _, round := range input {
		score += GAMES[round].part1
	}
	return score
}

func part2(input []string) any {
	score := 0
	for _, round := range input {
		score += GAMES[round].part2
	}
	return score
}
