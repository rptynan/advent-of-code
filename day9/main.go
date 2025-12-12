package main

import (
	"bufio"
	"fmt"
	"os"

	"adventofcode25/utils"
)

type Coord struct {
	row, col int
}

func calculateArea(c1, c2 Coord) int {
	return (utils.Abs(c1.row-c2.row) + 1) * (utils.Abs(c1.col-c2.col) + 1)
}

func part1(coords []Coord) {
	m := 0
	for c1 := range coords {
		for c2 := range coords {
			area := calculateArea(coords[c1], coords[c2])
			if m < area {
				fmt.Println(coords[c1], coords[c2], area)
				m = area
			}
		}
	}
	fmt.Println(m)
}

type Rectangle struct {
	maxRow, minRow, maxCol, minCol int
}

func getCornersOfRectangle(c1, c2 Coord) Rectangle {
	minRow := utils.Min(c1.row, c2.row)
	maxRow := utils.Max(c1.row, c2.row)
	minCol := utils.Min(c1.col, c2.col)
	maxCol := utils.Max(c1.col, c2.col)

	return Rectangle{
		maxRow: maxRow,
		minRow: minRow,
		maxCol: maxCol,
		minCol: minCol,
	}
}

func doesLineSegmentIntersectRectangle(r Rectangle, l1, l2 Coord) bool {
	// Vertical line segment
	if l1.col == l2.col {
		// First check the vertical line is within the rectangles vertical bounds
		if l1.col > r.minCol && l1.col < r.maxCol {
			// We know the line intersects if maxRow of line is above minRow of rectangle and minRow of line is below maxRow
			// of rectangle.
			maxLineRow := utils.Max(l1.row, l2.row)
			minLineRow := utils.Min(l1.row, l2.row)
			return maxLineRow > r.minRow && minLineRow < r.maxRow
		}
	}
	// Horizontal line segment
	if l1.row == l2.row { // should be true anyway if above isn't
		if l1.row > r.minRow && l1.row < r.maxRow {
			maxLineCol := utils.Max(l1.col, l2.col)
			minLineCol := utils.Min(l1.col, l2.col)
			return maxLineCol > r.minCol && minLineCol < r.maxCol
		}
	}
	return false
}

func doesAnyLineIntersectRectangle(r Rectangle, coords []Coord) bool {
	for i := range coords {
		l1 := coords[i]
		l2 := coords[(i+1)%len(coords)]
		if doesLineSegmentIntersectRectangle(r, l1, l2) {
			return true
		}
	}
	return false
}

func part2(coords []Coord) {
	maxArea := 0
	for c1 := range coords {
		for c2 := c1 + 1; c2 < len(coords); c2++ {
			area := calculateArea(coords[c1], coords[c2])
			rect := getCornersOfRectangle(coords[c1], coords[c2])
			if area > maxArea && !doesAnyLineIntersectRectangle(rect, coords) {
				fmt.Printf("New max area %d from %v to %v\n", area, coords[c1], coords[c2])
				maxArea = area
			}
		}
	}
	fmt.Println(maxArea)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxRow, maxCol := 0, 0
	coords := []Coord{}
	for scanner.Scan() {
		line := scanner.Text()
		coord := Coord{}
		fmt.Sscanf(line, "%d,%d", &coord.col, &coord.row)
		coords = append(coords, coord)

		if maxRow < coord.row {
			maxRow = coord.row
		}
		if maxCol < coord.col {
			maxCol = coord.col
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(coords)

	part2(coords)
}
