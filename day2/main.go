package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func stringIsPalindromeyBased(s string, base int) bool {
	if len(s)%base != 0 {
		return false
	}

	strLen := len(s)
	chunkLen := strLen / base
	for i := 0; i < chunkLen; i++ {
		for j := chunkLen + i; j < strLen; j += chunkLen {
			if s[i] != s[j] {
				return false
			}
		}
	}
	return true
}

func stringIsPalindromey(s string) bool {
	for base := 2; base <= len(s); base++ {
		if stringIsPalindromeyBased(s, base) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			fmt.Printf("Checking range: %v\n", r)

			bounds := strings.Split(r, "-")
			low, high := num(bounds[0]), num(bounds[1])

			for currentNumber := low; currentNumber <= high; currentNumber++ {
				strNum := strconv.Itoa(currentNumber)

				if stringIsPalindromey(strNum) {
					fmt.Println(strNum)
					res += currentNumber
				}
			}
		}

	}

	fmt.Printf("Result: %v\n", res)

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
