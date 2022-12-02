package main

import (
	"sort"
	"strings"

	"aoc2022/util"
)

func main() {
	util.Run(2022, 1, parseInput, part1, part2)
}

func parseInput(input string) (elves []int) {
	return util.Map(func(elf string) int {
		return util.Sum(util.ToInts(strings.Split(elf, "\n")))
	}, strings.Split(input, "\n\n"))
}

func part1(elves []int) int {
	return util.Max(elves...)
}

func part2(elves []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return elves[0] + elves[1] + elves[2]
}
