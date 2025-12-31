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
	result := 0
	for i := range m {
		for j := range n {
			if string(grid[i][j]) == "@" && canAccess(i, j, m, n, grid) {
				result += countRemoved(i, j, m, n, &grid)
			}
		}
	}
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func countRemoved(i int, j int, m int, n int, grid *[]string) int {
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

	(*grid)[i] = (*grid)[i][:j] + "." + (*grid)[i][j+1:]
	removed := 1
	for _, dir := range directions {
		iChg := dir[0]
		jChg := dir[1]
		nextI := i + iChg
		nextJ := j + jChg
		if 0 <= nextI && nextI < m && 0 <= nextJ && nextJ < n {
			if string((*grid)[nextI][nextJ]) == "@" && canAccess(nextI, nextJ, m, n, *grid) {
				removed += countRemoved(nextI, nextJ, m, n, grid)
			}
		}
	}
	return removed
}

func canAccess(i int, j int, m int, n int, grid []string) bool {
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
