package main

import (
	"bufio"
	"fmt"
	"os"
)

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func setGridCell(grid [][]rune, row, col int, value rune) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return
	}
	grid[row][col] = value
}

func part1(grid [][]rune) {
	splits := 0
	for row := 1; row < len(grid); row++ {
		// First pass, put | either side of ^
		for col := 0; col < len(grid[row]); col++ {
			if grid[row-1][col] == 'S' {
				setGridCell(grid, row, col, '|')
			}
			if grid[row][col] == '^' && grid[row-1][col] == '|' {
				splits++
				setGridCell(grid, row, col-1, '|')
				setGridCell(grid, row, col+1, '|')
			}
		}
		// Second pass, extend | down if there's a | above and not a ^ in current cell
		for col := 0; col < len(grid[row]); col++ {
			if row > 0 && grid[row-1][col] == '|' && grid[row][col] == '.' {
				setGridCell(grid, row, col, '|')
			}
		}
	}

	printGrid(grid)
	fmt.Printf("Total splits: %d\n", splits)
}

func findStart(grid [][]rune) (int, int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'S' {
				return row, col
			}
		}
	}
	panic("Start position not found")
}

type Coord struct {
	row int
	col int
}

func part2(grid [][]rune, memo map[Coord]int, curRow, curCol int) int {
	if curRow < 0 || curRow >= len(grid) || curCol < 0 || curCol >= len(grid[curRow]) {
		return 1
	}
	if memoized, ok := memo[Coord{curRow, curCol}]; ok {
		return memoized
	}
	fmt.Printf("At (%d, %d) = %c\n", curRow, curCol, grid[curRow][curCol])

	var res int
	switch grid[curRow][curCol] {
	case '^':
		res = part2(grid, memo, curRow, curCol-1) + part2(grid, memo, curRow, curCol+1)
	case '.', 'S':
		res = part2(grid, memo, curRow+1, curCol)
	default:
		panic("Unexpected cell value " + string(grid[curRow][curCol]))
	}

	memo[Coord{curRow, curCol}] = res
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	startRow, startCol := findStart(grid)
	memo := make(map[Coord]int)
	result := part2(grid, memo, startRow, startCol)
	fmt.Printf("Total paths: %d\n", result)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
