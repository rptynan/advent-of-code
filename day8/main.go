package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Coord struct {
	x, y, z float64
}

type Distance struct {
	coord1, coord2 int
	dist           float64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	coords := []Coord{}
	for scanner.Scan() {
		line := scanner.Text()
		coord := Coord{}
		fmt.Sscanf(line, "%f,%f,%f", &coord.x, &coord.y, &coord.z)
		coords = append(coords, coord)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(coords)

	// Get distances between all pairs of coordinates and sort them ascending
	dists := []Distance{}
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			dist := math.Abs(coords[i].x-coords[j].x)*math.Abs(coords[i].x-coords[j].x) +
				math.Abs(coords[i].y-coords[j].y)*math.Abs(coords[i].y-coords[j].y) +
				math.Abs(coords[i].z-coords[j].z)*math.Abs(coords[i].z-coords[j].z)
			dists = append(dists, Distance{coord1: i, coord2: j, dist: dist})
		}
	}

	sort.Slice(dists, func(i, j int) bool {
		return dists[i].dist < dists[j].dist
	})
	fmt.Println(dists)

	// Connect the closest pairs first, keeping track of sets
	sets := make([]int, len(coords))
	for i := range sets {
		sets[i] = i
	}

	NUM_CONNECTIONS := 10000000
	setsConnected := 0
	for i := 0; i < NUM_CONNECTIONS; i++ {
		d := dists[i]
		fmt.Println("Connecting", d.coord1, "and", d.coord2, "with distance", d.dist)
		fmt.Println(coords[d.coord1], coords[d.coord2])

		if sets[d.coord1] != sets[d.coord2] {
			oldSet := sets[d.coord2]
			newSet := sets[d.coord1]
			for i := range sets {
				if sets[i] == oldSet {
					sets[i] = newSet
				}
			}
			setsConnected++
			if setsConnected == len(coords)-1 {
				fmt.Println("All coordinates connected")
				panic("")
			}
		}
		// else {
		// 	NUM_CONNECTIONS++
		// }
	}

	// Print all set sizes
	sizes := []int{}
	for i := range sets {
		size := 0
		for _, s := range sets {
			if s == i {
				size++
			}
		}
		if size > 0 {
			sizes = append(sizes, size)
		}
		fmt.Printf("Set %d size: %d\n", i, size)
	}
	sort.Ints(sizes)
	fmt.Println("Set sizes:", sizes)
}
