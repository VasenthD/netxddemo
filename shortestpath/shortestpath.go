package main

import (
	"fmt"
	"math"
)

func dijkstra(graph [][]int, start int) []int {
	numNodes := len(graph)
	dist := make([]int, numNodes)
	visited := make([]bool, numNodes)

	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	for count := 0; count < numNodes-1; count++ {
		u := minDistIndex(dist, visited)
		visited[u] = true

		for v := 0; v < numNodes; v++ {
			if !visited[v] && graph[u][v] != 0 && dist[u] != math.MaxInt32 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return dist
}

func minDistIndex(dist []int, visited []bool) int {
	minDist := math.MaxInt32
	minIndex := -1

	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= minDist {
			minDist = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

func main() {
	graph := [][]int{
		{0, 3, 0, 0, 0},
		{3, 0, 5, 1, 0},
		{0, 5, 0, 2, 4},
		{0, 1, 2, 0, 6},
		{0, 0, 4, 6, 0},
	}

	start := 0
	shortestDistances := dijkstra(graph, start)

	fmt.Println("Shortest distances from node", start, "to all other nodes:")
	for i, distance := range shortestDistances {
		fmt.Printf("Node %d: Distance %d\n", i, distance)
	}
}

