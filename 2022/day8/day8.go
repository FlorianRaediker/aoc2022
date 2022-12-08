package main

import (
	"aoc2022/util"
	"fmt"
	"strings"
)

func main() {
	util.Run(2022, 8, parseInput, part1, part2)
}

func parseInput(input string) [][]int {
	grid := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, util.ToInts(strings.Split(line, "")))
	}
	return grid
}

func part1(grid [][]int) any {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			top := make([]int, 0)
			for i := 0; i < row; i++ {
				top = append(top, grid[i][col])
			}
			bottom := make([]int, 0)
			for i := row + 1; i < len(grid); i++ {
				bottom = append(bottom, grid[i][col])
			}
			left := grid[row][:col]
			right := grid[row][col+1:]
			isVisible := func(h int) bool {
				return h < grid[row][col]
			}
			if util.All(isVisible, top) || util.All(isVisible, bottom) || util.All(isVisible, left) || util.All(isVisible, right) {
				fmt.Println(row, col, grid[row][col])
				count++
			}
		}
	}
	return count
}

func part2(grid [][]int) any {
	max := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			top := make([]int, 0)
			for i := row - 1; i >= 0; i-- {
				top = append(top, grid[i][col])
			}
			bottom := make([]int, 0)
			for i := row + 1; i < len(grid); i++ {
				bottom = append(bottom, grid[i][col])
			}
			left := make([]int, 0)
			for i := col - 1; i >= 0; i-- {
				left = append(left, grid[row][i])
			}
			right := grid[row][col+1:]
			countTrees := func(trees []int) int {
				for i, tree := range trees {
					if tree >= grid[row][col] {
						return i + 1
					}
				}
				return len(trees)
			}
			score := countTrees(top) * countTrees(bottom) * countTrees(left) * countTrees(right)
			if score > max {
				max = score
			}
		}
	}
	return max
}
