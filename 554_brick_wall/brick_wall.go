package brick_wall

import (
// "fmt"
)

func leastBricks(wall [][]int) int {
	edges := countEdges(wall)
	height := len(wall)
	return height - maxEdges(edges)
}

func countEdges(wall [][]int) map[int]int {

	edges := make(map[int]int)

	for _, row := range wall {
		position := 0

		for _, brick := range row {
			edges[position]++
			position += brick
		}
	}
	return edges
}

func maxEdges(edges map[int]int) int {
	max_edges := 0
	for position, edges := range edges {
		if position > 0 {
			max_edges = max(max_edges, edges)
		}
		// fmt.Printf("Position %v has %v edges\n", position, edges)
	}
	return max_edges
}

func max(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}
