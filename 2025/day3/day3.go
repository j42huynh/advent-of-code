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
	banks := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		banks = append(banks, line)
	}

	fmt.Print("Part1: ")
	fmt.Println(part1(banks))
	fmt.Print("Part2: ")
	fmt.Println(part2(banks))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1(banks []string) int {
	result := 0
	for _, bank := range banks {
		result += getMaxJoltage(bank, 2)
	}
	return result
}

func part2(banks []string) int {
	result := 0
	for _, bank := range banks {
		result += getMaxJoltage(bank, 12)
	}
	return result
}

func getMaxJoltage(bank string, numBatteriesOn int) int {
	maxJoltage := 0
	for i := range numBatteriesOn {
		partialBank := bank[:len(bank)-(numBatteriesOn-1-i)]

		maxDigit := -1
		maxIndex := -1
		for j, d := range partialBank {
			digit := getInt(string(d))
			if digit > maxDigit {
				maxDigit = digit
				maxIndex = j
			}
		}
		maxJoltage = maxJoltage*10 + maxDigit

		bank = bank[maxIndex+1:]
	}
	return maxJoltage
}

func getInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return num
}
