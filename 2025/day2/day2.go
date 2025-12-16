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
	for scanner.Scan() {
		line := scanner.Text()

		ids := strings.Split(line, ",")
		result := 0
		for _, id := range ids {
			id_parts := strings.Split(id, "-")

			start, err := strconv.Atoi(id_parts[0])
			if err != nil {
				log.Fatal(err)
				return
			}

			end, err := strconv.Atoi(id_parts[1])
			if err != nil {
				log.Fatal(err)
				return
			}

			result += validateRange(start, end)
		}

		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func validateRange(start int, end int) int {
	total := 0

	for num := start; num <= end; num++ {
		id := strconv.Itoa(num)
		if validateId(id) {
			total += num
		}
	}

	return total
}

func validateId(id string) bool {
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
