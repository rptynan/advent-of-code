package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dial, res := 50, 0
	for scanner.Scan() {
		line := scanner.Text()

		letter := line[0]
		numberStr := line[1:]

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic("number not valid")
		}

		switch letter {
		case 'L':
			dial -= number
		case 'R':
			dial += number
		default:
			panic("invalid input")
		}
		fmt.Printf("You entered: %s\n", line)
		fmt.Println(dial)

		for dial < 0 {
			dial += 100
			res++
		}

		for dial >= 100 {
			dial -= 100
			res++
		}
		// fmt.Println(dial)
	}

	fmt.Printf("Result: %v\n", res)

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
