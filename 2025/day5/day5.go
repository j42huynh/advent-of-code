package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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
	reachedNewline := false
	ranges := [][]int{}
	ingredients := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			reachedNewline = true
			continue
		}

		if !reachedNewline {
			rangeStrings := strings.Split(line, "-")
			lower := getInt(rangeStrings[0])
			upper := getInt(rangeStrings[1])
			ranges = append(ranges, []int{lower, upper})
		} else {
			ingredient := getInt(line)
			ingredients = append(ingredients, ingredient)
		}
	}

	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
	slices.Sort(ingredients)

	//result := part1(ranges, ingredients)
	result := part2(ranges)
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return num
}

func part1(ranges [][]int, ingredients []int) int {
	rangeIndex := 0
	ingredientIndex := 0
	result := 0
	for ingredientIndex < len(ingredients) {
		if rangeIndex >= len(ranges) {
			break
		}

		ingredient := ingredients[ingredientIndex]
		if ingredient < ranges[rangeIndex][0] {
			ingredientIndex += 1
		} else if ingredient > ranges[rangeIndex][1] {
			rangeIndex += 1
		} else {
			result += 1
			ingredientIndex += 1
		}
	}
	return result
}

func part2(ranges [][]int) int {
	currentRange := ranges[0]
	mergedRanges := [][]int{}
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		if r[0] <= currentRange[1] {
			currentRange[1] = max(r[1], currentRange[1])
		} else {
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = r
		}
	}
	mergedRanges = append(mergedRanges, currentRange)

	result := 0
	for _, r := range mergedRanges {
		result += r[1] - r[0] + 1
	}
	return result
}
