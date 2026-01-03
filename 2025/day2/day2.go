package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	idRanges := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, ",")

		for _, id := range ids {
			id_parts := strings.Split(id, "-")

			start := getInt(id_parts[0])
			end := getInt(id_parts[1])
			idRanges = append(idRanges, []int{start, end})
		}
	}

	fmt.Print("Part1: ")
	fmt.Println(part1(idRanges))
	fmt.Print("Part2: ")
	fmt.Println(part2(idRanges))

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

func part1(idRanges [][]int) int {
	result := 0
	for _, idRange := range idRanges {
		start := idRange[0]
		end := idRange[1]
		for num := start; num <= end; num++ {
			id := strconv.Itoa(num)
			if validateId1(id) {
				result += num
			}
		}
	}
	return result
}

func part2(idRanges [][]int) int {
	result := 0
	for _, idRange := range idRanges {
		start := idRange[0]
		end := idRange[1]
		for num := start; num <= end; num++ {
			id := strconv.Itoa(num)
			if validateId2(id) {
				result += num
			}
		}
	}
	return result
}

func validateId1(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	mid := len(id) / 2
	return id[:mid] == id[mid:]
}

func validateId2(id string) bool {
	partMax := len(id) / 2
	for partSize := 1; partSize <= partMax; partSize++ {
		if len(id)%partSize != 0 {
			continue
		}

		part := id[0:partSize]
		partsEqual := true
		for index := partSize; index <= len(id)-partSize; index += partSize {
			nextPart := id[index:min(index+partSize, len(id))]
			if part != nextPart {
				partsEqual = false
				break
			}
		}

		if partsEqual {
			return true
		}
	}
	return false
}
