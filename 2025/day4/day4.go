package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	m := len(grid)
	n := len(grid[0])
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	fmt.Print("Part1: ")
	fmt.Println(part1(m, n, grid, directions))
	fmt.Print("Part2: ")
	fmt.Println(part2(m, n, grid, directions))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1(m int, n int, grid []string, directions [][]int) int {
	result := 0
	for i := range m {
		for j := range n {
			if string(grid[i][j]) == "@" && canAccess(i, j, m, n, grid, directions) {
				result += 1
			}
		}
	}
	return result
}

func part2(m int, n int, grid []string, directions [][]int) int {
	result := 0
	for i := range m {
		for j := range n {
			if string(grid[i][j]) == "@" && canAccess(i, j, m, n, grid, directions) {
				result += countRemoved(i, j, m, n, &grid, directions)
			}
		}
	}
	return result
}

func countRemoved(i int, j int, m int, n int, grid *[]string, directions [][]int) int {
	(*grid)[i] = (*grid)[i][:j] + "." + (*grid)[i][j+1:]
	removed := 1
	for _, dir := range directions {
		iChg := dir[0]
		jChg := dir[1]
		nextI := i + iChg
		nextJ := j + jChg
		if 0 <= nextI && nextI < m && 0 <= nextJ && nextJ < n {
			if string((*grid)[nextI][nextJ]) == "@" && canAccess(nextI, nextJ, m, n, *grid, directions) {
				removed += countRemoved(nextI, nextJ, m, n, grid, directions)
			}
		}
	}
	return removed
}

func canAccess(i int, j int, m int, n int, grid []string, directions [][]int) bool {
	numPapers := 0
	for _, dir := range directions {
		iChg := dir[0]
		jChg := dir[1]
		nextI := i + iChg
		nextJ := j + jChg
		if 0 <= nextI && nextI < m && 0 <= nextJ && nextJ < n {
			if string(grid[nextI][nextJ]) == "@" {
				numPapers += 1
			}
		}
	}
	return numPapers < 4
}
