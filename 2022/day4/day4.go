package main

import (
	"aoc2022/util"
	"fmt"
	"strings"
)

func main() {
	util.Run(2022, 4, parseInput, part1, part2)
}

type Assignment struct {
	start1 int
	end1   int
	start2 int
	end2   int
}

func parseInput(input string) []Assignment {
	return util.Map(func(line string) Assignment {
		var assignment Assignment
		fmt.Sscanf(line, "%d-%d,%d-%d", &assignment.start1, &assignment.end1, &assignment.start2, &assignment.end2)
		return assignment
	}, strings.Split(input, "\n"))
}

func part1(input []Assignment) (count int) {
	for _, assignment := range input {
		if (assignment.start1 <= assignment.start2 && assignment.end2 <= assignment.end1) ||
			(assignment.start2 <= assignment.start1 && assignment.end1 <= assignment.end2) {
			count += 1
		}
	}
	return
}

func part2(input []Assignment) (count int) {
	for _, a := range input {
		if (a.start1 <= a.start2 && a.start2 <= a.end1) ||
			(a.start1 <= a.end2 && a.end2 <= a.end1) ||
			(a.start2 <= a.start1 && a.start1 <= a.end2) ||
			(a.start2 <= a.end1 && a.end1 <= a.end2) {
			count += 1
		}
	}
	return
}
