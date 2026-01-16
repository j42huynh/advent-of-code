package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatal("missing input file argument")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	fmt.Print("Part1: ")
	fmt.Println(part1(grid))
	fmt.Print("Part2: ")
	fmt.Println(part2(grid))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1(grid []string) int {
	start := strings.Index(grid[0], "S")
	indexes := map[int]int{start: 0}
	result := 0
	for i := 1; i < len(grid); i++ {
		next_indexes := map[int]int{}
		for ind := range indexes {
			if string(grid[i][ind]) == "^" {
				result += 1
				if ind-1 >= 0 {
					next_indexes[ind-1] = 0
				}
				if ind+1 < len(grid[i]) {
					next_indexes[ind+1] = 0
				}
			} else {
				next_indexes[ind] = 0
			}
		}
		indexes = next_indexes
	}

	return result
}

func part2(grid []string) int {
	start := strings.Index(grid[0], "S")
	memo := [][]int{}
	for range len(grid) {
		memoRow := []int{}
		for range len(grid[0]) {
			memoRow = append(memoRow, -1)
		}
		memo = append(memo, memoRow)
	}
	return dfs(0, start, grid, &memo)
}

func dfs(row int, col int, grid []string, memo *[][]int) int {
	if row == len(grid) {
		return 1
	}

	if (*memo)[row][col] != -1 {
		return (*memo)[row][col]
	}

	if string(grid[row][col]) == "^" {
		left := 0
		if col-1 >= 0 {
			left = dfs(row+1, col-1, grid, memo)
		}

		right := 0
		if col+1 < len(grid[0]) {
			right = dfs(row+1, col+1, grid, memo)
		}

		res := left + right
		(*memo)[row][col] = res
		return res
	} else {
		res := dfs(row+1, col, grid, memo)
		(*memo)[row][col] = res
		return res
	}
}
