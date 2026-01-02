package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	rotations := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		rotations = append(rotations, line)
	}

	fmt.Print("Part1: ")
	fmt.Println(part1(rotations))
	fmt.Print("Part2: ")
	fmt.Println(part2(rotations))

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

func part1(rotations []string) int {
	dialMax := 100
	result := 0
	start := 50
	for _, rotation := range rotations {
		direction := rotation[:1]
		num := getInt(rotation[1:])

		switch direction {
		case "L":
			start -= num
		case "R":
			start += num
		default:
			log.Fatal("invalid direction")
			return 0
		}

		if start%dialMax == 0 {
			result += 1
		}
	}
	return result
}

func part2(rotations []string) int {
	result := 0
	start := 50
	for _, rotation := range rotations {
		direction := rotation[:1]
		num := getInt(rotation[1:])

		nextStart, clicks := applyRotation(start, direction, num)
		result += clicks
		start = nextStart
	}
	return result
}

func applyRotation(start int, direction string, num int) (int, int) {
	dialMax := 100
	clicks := num / dialMax
	toAdd := num % dialMax
	sign := 0

	if toAdd == 0 {
		return start, clicks
	}

	switch direction {
	case "L":
		sign = -1
	case "R":
		sign = 1
	default:
		log.Fatal("invalid direction")
		return 0, 0
	}

	startedAtZero := start == 0
	start += sign * toAdd
	if start == 0 && !startedAtZero {
		clicks += 1
	} else if start < 0 {
		if !startedAtZero {
			clicks += 1
		}
		start += dialMax
	} else if start >= dialMax {
		clicks += 1
		start -= dialMax
	}

	return start, clicks
}
