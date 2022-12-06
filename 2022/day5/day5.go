package main

import (
	"aoc2022/util"
	"fmt"
	"strings"
)

func main() {
	util.Run(2022, 5, parseInput, part1, part2)
}

type Input struct {
	stacks [][]string
	moves  []Move
}

type Move struct {
	count    int
	srcStack int
	dstStack int
}

func parseInput(input string) Input {
	parts := strings.Split(input, "\n\n")
	cratesStr := strings.Split(parts[0], "\n")
	cratesStr = cratesStr[:len(cratesStr)-1]
	var stacks [][]string
	stacks = append(stacks, nil)
	for i := 0; i < 9; i++ {
		col := 1 + 4*i
		var stack []string
		for row := len(cratesStr) - 1; row >= 0; row-- {
			crate := string(cratesStr[row][col])
			if crate == " " {
				break
			}
			stack = append(stack, crate)
		}
		stacks = append(stacks, stack)
	}
	var moves []Move
	for _, moveStr := range strings.Split(parts[1], "\n") {
		var move Move
		fmt.Sscanf(moveStr, "move %d from %d to %d", &move.count, &move.srcStack, &move.dstStack)
		moves = append(moves, move)
	}
	return Input{stacks, moves}
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func part1(input Input) any {
	stacks := make([][]string, len(input.stacks))
	for i, s := range input.stacks {
		stack := make([]string, len(s))
		copy(stack, s)
		stacks[i] = stack
	}
	for _, move := range input.moves {
		movedCrates := stacks[move.srcStack][len(stacks[move.srcStack])-move.count:]
		reverse(movedCrates)
		stacks[move.dstStack] = append(stacks[move.dstStack], movedCrates...)
		stacks[move.srcStack] = stacks[move.srcStack][:len(stacks[move.srcStack])-move.count]
	}
	tops := ""
	for _, stack := range stacks[1:] {
		if len(stack) > 0 {
			tops += stack[len(stack)-1]
		}
	}
	return tops
}

func part2(input Input) any {
	stacks := make([][]string, len(input.stacks))
	for i, s := range input.stacks {
		stack := make([]string, len(s))
		copy(stack, s)
		stacks[i] = stack
	}
	for _, move := range input.moves {
		stacks[move.dstStack] = append(stacks[move.dstStack], stacks[move.srcStack][len(stacks[move.srcStack])-move.count:]...)
		stacks[move.srcStack] = stacks[move.srcStack][:len(stacks[move.srcStack])-move.count]
	}
	tops := ""
	for _, stack := range stacks[1:] {
		if len(stack) > 0 {
			tops += stack[len(stack)-1]
		}
	}
	return tops
}
