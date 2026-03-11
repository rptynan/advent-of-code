package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Assuming input is DAG

// I think this solution is not strictly correct, as it doesn't take into account
// topological ordering, but works for the input for part 1.
func bfs(edges map[string][]string, start, end string) int {
	queue := []string{start}
	// Map of visited nodes to how many paths reach them
	visited := make(map[string]int)
	visited[start] = 1

	for len(queue) > 0 {
		// fmt.Printf("Queue: %v\n", queue)
		current := queue[0]
		queue = queue[1:]

		for _, destination := range edges[current] {
			// fmt.Printf("  Current: %s, Destination: %s\n", current, destination)
			if _, ok := visited[destination]; !ok {
				// First time visiting this node
				// Add to queue to process
				// Set number of paths to how many we can read current from
				queue = append(queue, destination)
				visited[destination] = visited[current]
			} else {
				// Already visited this node
				// Just increment number of paths to reach this node
				visited[destination] += visited[current]
			}
		}
	}

	return visited[end]
}

// An actually correct path counting using DFS with memoization.
func dfsCountPaths(edges map[string][]string, current, end string, memo map[string]int) int {
	if current == end {
		return 1
	}

	if val, ok := memo[current]; ok {
		return val
	}

	totalPaths := 0
	for _, neighbor := range edges[current] {
		pathsFromNeighbor := dfsCountPaths(edges, neighbor, end, memo)
		totalPaths += pathsFromNeighbor
	}

	memo[current] = totalPaths
	return totalPaths
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Key is source, value is list of destinations
	edges := make(map[string][]string)
	// Given node x, distance from start node. All initialized to 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("Processing line: %s\n", line)
		tokens := strings.Split(line, " ")

		// First token with : stripped off
		source := tokens[0][0 : len(tokens[0])-1]
		destinations := tokens[1:]

		// fmt.Printf("Source: %s\n", source)
		// fmt.Printf("Destinations: %v\n", destinations)
		edges[source] = destinations
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Part 1
	// var (
	// 	StartNode = "you"
	// 	EndNode   = "out"
	// )
	// total := bfs(edges, StartNode, EndNode)
	// fmt.Println(total)
	// fmt.Println(dfsCountPaths(edges, StartNode, EndNode, make(map[string]int)))

	// Part 2
	// Assuming a DAG, only one of these will be a valid path
	fmt.Println("---- dac before fft ----")
	fmt.Println(dfsCountPaths(edges, "svr", "dac", make(map[string]int)))
	fmt.Println(dfsCountPaths(edges, "dac", "fft", make(map[string]int)))
	fmt.Println(dfsCountPaths(edges, "fft", "out", make(map[string]int)))
	fmt.Println("---- fft before dac ----")
	fmt.Println(dfsCountPaths(edges, "svr", "fft", make(map[string]int)))
	fmt.Println(dfsCountPaths(edges, "fft", "dac", make(map[string]int)))
	fmt.Println(dfsCountPaths(edges, "dac", "out", make(map[string]int)))
	// Total paths is the sum of the multiplications of the segments in the valid path
}
