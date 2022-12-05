package main

import (
	"aoc2022/util"
	"strings"
)

type Rucksack struct {
	content string
	first   string
	second  string
}

func main() {
	util.Run(2022, 3, parseInput, part1, part2)
}

func parseInput(input string) []Rucksack {
	return util.Map(func(line string) Rucksack {
		return Rucksack{content: line, first: line[:len(line)/2], second: line[len(line)/2:]}
	}, strings.Split(input, "\n"))
}

func getPriority(item rune) int {
	if int('a') <= int(item) && int(item) <= int('z') {
		return int(item) - int('a') + 1
	} else if int('A') <= int(item) && int(item) <= int('Z') {
		return int(item) - int('A') + 27
	} else {
		panic("invalid item " + string(item))
	}
}

func part1(input []Rucksack) any {
	sum := 0
	for _, rucksack := range input {
		var item rune = -1
		for _, r := range rucksack.first {
			if strings.Contains(rucksack.second, string(r)) {
				if item != -1 && item != r {
					panic("rucksack contains more than two double letters")
				}
				item = r
			}
		}
		sum += getPriority(item)
	}
	return sum
}

func part2(input []Rucksack) any {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		var badge rune = -1
		for _, r := range input[i].content {
			if strings.Contains(input[i+1].content, string(r)) && strings.Contains(input[i+2].content, string(r)) {
				if badge != -1 && badge != r {
					panic("rucksacks contain more than one badge")
				}
				badge = r
			}
		}
		sum += getPriority(badge)
	}
	return sum
}
