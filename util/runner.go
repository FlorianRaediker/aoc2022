package util

import (
	"flag"
	"fmt"
)

func Run[T interface{}](year, day int, parseInput func(string) T, part1, part2 func(T) int) {
	var part string
	flag.StringVar(&part, "part", "all", "part(s) to run: 1, 2 or all")

	flag.Parse()

	if part != "1" && part != "2" && part != "all" {
		panic("invalid part")
	}

	input := parseInput(GetInput(year, day))

	if part == "1" || part == "all" {
		fmt.Println("Part 1:", part1(input))
	}
	if part == "2" || part == "all" {
		fmt.Println("Part 2:", part2(input))
	}
}
