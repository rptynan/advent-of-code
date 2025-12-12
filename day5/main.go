package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func findFreshIngredients(ranges [][2]int, ingredients []int) int {
	res := 0

	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				res++
				break
			}
		}
	}

	return res
}

func countAllFreshIngredients(ranges [][2]int) int {
	// sort ranges by start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	fmt.Println("Sorted ranges:", ranges)

	i, res := -1, 0
	for _, r := range ranges {
		// easy case, range is above our highest point so far, count it all
		if r[0] > i {
			res += r[1] - r[0] + 1
			i = r[1]
		}
		// range overlaps with our highest point so far, count only the new part
		if r[1] > i {
			res += r[1] - i
			i = r[1]
		}
		// otherwise range has been fully counted already, so we do nothing
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ranges := [][2]int{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		r := [2]int{}
		fmt.Sscanf(line, "%d-%d", &r[0], &r[1])
		ranges = append(ranges, r)
	}
	fmt.Println("Read ranges:", ranges)

	ingredients := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		ingredient := 0
		fmt.Sscanf(line, "%d", &ingredient)
		ingredients = append(ingredients, ingredient)
	}
	fmt.Println("Read ingredients:", ingredients)

	fmt.Println("Number of fresh ingredients:", findFreshIngredients(ranges, ingredients))

	fmt.Println("Total number of fresh ingredients in all ranges:", countAllFreshIngredients(ranges))

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
