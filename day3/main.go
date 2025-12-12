package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Returns index
// low, high inclusive
func maxJoltageInRange(joltages []int, low, high int) int {
	maxIndex := low
	for i := low + 1; i <= high; i++ {
		if joltages[maxIndex] < joltages[i] {
			maxIndex = i
		}
	}
	return maxIndex
}

func maxJoltagePart1(joltages []int) int {
	maxFirstIndex := maxJoltageInRange(joltages, 0, len(joltages)-2)

	maxSecondIndex := maxJoltageInRange(joltages, maxFirstIndex+1, len(joltages)-1)

	return joltages[maxFirstIndex]*10 + joltages[maxSecondIndex]
}

func maxJoltagePart2(joltages []int) int {
	result := 0
	var highestIndexSoFar int = -1

	// i here is number of digits we can't look at from the end
	for i := 12; i > 0; i-- {
		maxIndex := maxJoltageInRange(joltages, highestIndexSoFar+1, len(joltages)-i)

		result = result*10 + joltages[maxIndex]
		highestIndexSoFar = maxIndex
	}

	fmt.Println("Intermediate result:", result)
	return result
}

func parseBank(line string) []int {
	var res []int
	for _, char := range line {
		digit, _ := strconv.Atoi(string(char))
		res = append(res, digit)
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		joltages := parseBank(line)
		joltage := maxJoltagePart2(joltages)
		fmt.Println(joltage)

		res += joltage
	}

	fmt.Printf("Result: %d\n", res)

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
