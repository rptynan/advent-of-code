package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"adventofcode25/utils"
)

func getIthBit(n, i int) int {
	return (n >> i) & 1
}

func solveCasePart1(goal int, buttonMasks []int) int {
	numPresses := math.MaxInt64
	limit := 1 << len(buttonMasks)
	fmt.Printf("Trying all button combinations up to %b\n", limit)

	// i represents which buttons are pressed, e.g. 101 means button 0 and 2 are pressed
	for i := 0; i < limit; i++ {
		fmt.Printf("Trying combination: %b\n", i)
		currentState := 0

		for b := range buttonMasks {
			if getIthBit(i, b) == 1 {
				// fmt.Printf("Pressing button %d (%b)\n", b, buttonMasks[b])
				currentState ^= buttonMasks[b]
			}
		}

		fmt.Printf("Current state after presses: %b\n", currentState)
		if currentState == goal {
			fmt.Printf("Found solution with button presses: %b\n", i)
			n := utils.CountSetBits(i)
			if n < numPresses {
				numPresses = n
			}
		}
	}
	return numPresses
}

func parseButton(buttons []string) [][]int {
	res := [][]int{}
	for _, button := range buttons {
		strNums := strings.Split(button[1:len(button)-1], ",")

		nums := []int{}
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		res = append(res, nums)
	}
	return res
}

// Represent each button as bitmask, e.g. (0,2,3) -> 1101
func makeButtonsBitmasks(buttons [][]int) []int {
	res := []int{}
	for _, button := range buttons {
		r := 0
		for _, num := range button {
			r |= 1 << num
		}
		res = append(res, r)
		fmt.Printf("%v -> %b\n", button, r)
	}
	return res
}

func makeGoalsBitmask(goal string) int {
	r := 0
	for i, ch := range goal {
		if ch == '#' {
			r |= 1 << i
		}
	}
	fmt.Printf("Goal: %b\n", r)
	return r
}

func parseJoltages(joltagesStr string) []int {
	strNums := strings.Split(joltagesStr[1:len(joltagesStr)-1], ",")
	joltages := []int{}
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}
		joltages = append(joltages, num)
	}
	return joltages
}

func incrementButtonPresses(buttonPresses []int, maxJoltage int) bool {
	buttonPresses[0]++
	for i := range buttonPresses {
		if buttonPresses[i] > maxJoltage {
			if (i + 1) >= len(buttonPresses) {
				return false
			}
			buttonPresses[i] = 0
			buttonPresses[i+1]++
		}
	}
	return true
}

func countTotalPresses(buttonPresses []int) int {
	total := 0
	for _, presses := range buttonPresses {
		total += presses
	}
	return total
}

func solveCasePart2(buttonDefs [][]int, joltageReqs []int) int {
	minPresses := math.MaxInt64
	buttonPresses := make([]int, len(buttonDefs))
	maxJoltageReq := 0
	for _, req := range joltageReqs {
		if req > maxJoltageReq {
			maxJoltageReq = req
		}
	}

	for {
		// fmt.Printf("Trying button presses: %v\n", buttonPresses)

		// Optimisation, skip if we already exceed minPresses
		for countTotalPresses(buttonPresses) >= minPresses {
			if !incrementButtonPresses(buttonPresses, maxJoltageReq) {
				return minPresses
			}
		}

		joltages := make([]int, len(joltageReqs))
		for buttonIndex, presses := range buttonPresses {
			buttonDef := buttonDefs[buttonIndex]
			for _, joltageToIncrease := range buttonDef {
				joltages[joltageToIncrease] += presses

				// Optimisation, skip if we already exceed maxJoltageReq
				if joltages[joltageToIncrease] > maxJoltageReq {
					if !incrementButtonPresses(buttonPresses, maxJoltageReq) {
						return minPresses
					}
					continue
				}
			}
		}
		// fmt.Printf("Resulting joltages: %v\n", joltages)

		// Check if we have a solution
		if slices.Equal(joltages, joltageReqs) {
			totalPresses := 0
			for _, presses := range buttonPresses {
				totalPresses += presses
			}
			if totalPresses < minPresses {
				minPresses = totalPresses
			}
			fmt.Printf("Found solution with button presses: %v (total presses: %d)\n", buttonPresses, totalPresses)
		}

		if !incrementButtonPresses(buttonPresses, maxJoltageReq) {
			break
		}
	}

	return minPresses
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		var goal, joltages string
		var buttonDefs [][]int

		line := scanner.Text()
		fmt.Printf("Processing line: %s\n", line)
		chunks := strings.Split(line, " ")

		goal = chunks[0][1 : len(chunks[0])-1]
		buttonDefs = parseButton(chunks[1 : len(chunks)-1])
		joltages = chunks[len(chunks)-1]
		fmt.Println(goal, buttonDefs, joltages)

		buttonBitmasks := makeButtonsBitmasks(buttonDefs)
		goalBitmask := makeGoalsBitmask(goal)
		joltagesParsed := parseJoltages(joltages)
		fmt.Println(buttonBitmasks, goalBitmask, joltagesParsed)

		// result := solveCasePart1(goalBitmask, buttonBitmasks)
		result := solveCasePart2(buttonDefs, joltagesParsed)
		fmt.Printf("Result: %d\n", result)
		total += result
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total: %d\n", total)
}
