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
	result := 0
	start := 50
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[:1]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
			return
		}

		next_start, clicks := applyRotation(start, direction, num)
		result += clicks
		start = next_start
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func applyRotation(start int, direction string, num int) (int, int) {
	dial_max := 100
	clicks := num / dial_max
	to_add := num % dial_max
	sign := 0

	if to_add == 0 {
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

	started_at_zero := start == 0
	start += sign * to_add
	if start == 0 && !started_at_zero {
		clicks += 1
	} else if start < 0 {
		if !started_at_zero {
			clicks += 1
		}
		start += dial_max
	} else if start >= dial_max {
		clicks += 1
		start -= dial_max
	}

	return start, clicks
}
