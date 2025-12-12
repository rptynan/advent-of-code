package main

import (
	"bufio"
	"fmt"
	"os"
)

func printRuneGrid(grid [][]rune) {
	for i, row := range grid {
		fmt.Printf("Row %d: %s\n", i, string(row))
	}
}

func countAdjacentRolls(grid [][]rune, x, y int, xLen, yLen int) int {
	dirs := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	count := 0
	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]

		if nx < 0 || nx >= xLen || ny < 0 || ny >= yLen {
			continue
		}

		if grid[ny][nx] == '@' {
			count++
		}

	}

	return count
}

func countAccessibleRolls(grid [][]rune) int {
	yLen := len(grid)
	xLen := len(grid[0])

	res := 0
	for y, row := range grid {
		midres := []int{}

		for x, cell := range row {
			adjRolls := countAdjacentRolls(grid, x, y, xLen, yLen)
			midres = append(midres, adjRolls)

			if cell == '@' && adjRolls < 4 {
				grid[y][x] = 'x'
				res++
			}
		}

		fmt.Println(midres)
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)
		grid = append(grid, row)
	}

	printRuneGrid(grid)

	accRes := 0
	for {
		accessibleRolls := countAccessibleRolls(grid)
		fmt.Println("Accessible rolls:", accessibleRolls)
		if accessibleRolls == 0 {
			fmt.Println("No more accessible rolls.")
			break
		}
		accRes += accessibleRolls
	}
	fmt.Println("Total accessible rolls:", accRes)

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
