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
	for scanner.Scan() {
		line := scanner.Text()

		joltage := getMaxJoltage(line)
		fmt.Println(joltage)
		result += joltage
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getMaxJoltage(bank string) int {
	maxJoltage := 0
	for i := range 12 {
		partialBank := bank[:len(bank)-(11-i)]

		maxDigit := -1
		maxIndex := -1
		for i, d := range partialBank {
			digit := getInt(string(d))
			if digit > maxDigit {
				maxDigit = digit
				maxIndex = i
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
