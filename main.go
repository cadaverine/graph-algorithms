package main

import (
	"fmt"

	"./graph"
)

// Pair - pair of vertex indeces
type Pair struct {
	FromID int
	ToID int
}

func printGraph(graphStr string, pairs []Pair) {
	gr := graph.CreateFromString(graphStr)

	fmt.Print("\n\nAdjacency table:\n\n")
	fmt.Println(gr)

	for _, pair := range pairs {
		fmt.Printf(
			"distance between %+v and %+v: %+v\n",
			pair.FromID,
			pair.ToID,
			gr.Dijkstra(gr.GetVertexByID(pair.FromID), gr.GetVertexByID(pair.ToID)),
		)
	}
}

func main() {
	for _, graph := range Graphs {
		printGraph(graph, []Pair{
			Pair{4, 3},
			Pair{1, 3},
			Pair{3, 2},
			Pair{0, 6},
		})
	}
}
