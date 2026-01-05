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
	// for part 1
	nums := [][]int{}
	operations := []string{}
	// for part 2
	numStrings := []string{}
	operationsString := ""
	for scanner.Scan() {
		line := scanner.Text()

		lineList := strings.Fields(line)
		if string(line[0]) == "*" || string(line[0]) == "+" {
			operationsString = line
			operations = lineList
		} else {
			numStrings = append(numStrings, line)
			numList := []int{}
			for _, numStr := range lineList {
				numList = append(numList, getInt(numStr))
			}
			nums = append(nums, numList)
		}
	}

	fmt.Print("Part1: ")
	fmt.Println(part1(nums, operations))
	fmt.Print("Part2: ")
	fmt.Println(part2(numStrings, operationsString))

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

func part1(nums [][]int, operations []string) int {
	result := 0
	for i, op := range operations {
		tmp := 0
		if op == "*" {
			tmp = 1
		}

		for x := range len(nums) {
			if op == "*" {
				tmp *= nums[x][i]
			} else {
				tmp += nums[x][i]
			}
		}
		result += tmp
	}
	return result
}

func part2(numStrings []string, operations string) int {
	result := 0
	nums := []int{}
	lastOp := string(operations[0])
	for i := range len(operations) {
		if i < len(operations)-1 && string(operations[i+1]) != " " {
			if lastOp == "*" {
				result += multiply(nums)
			} else {
				result += sum(nums)
			}
			nums = []int{}
			lastOp = string(operations[i+1])
			continue
		}

		num := 0
		for x := range len(numStrings) {
			c := string(numStrings[x][i])
			if c == " " {
				continue
			} else {
				num = num*10 + getInt(c)
			}
		}
		nums = append(nums, num)
	}

	if lastOp == "*" {
		result += multiply(nums)
	} else {
		result += sum(nums)
	}
	return result
}

func sum(nums []int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

func multiply(nums []int) int {
	result := 1
	for _, n := range nums {
		result *= n
	}
	return result
}
