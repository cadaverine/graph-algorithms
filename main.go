package main

import (
	"fmt"

	"./graph"
)

func main() {
	gr := graph.InitializeGraph()

	gr.AddVertex(1)
	gr.AddVertex(2)
	gr.AddVertex(3)
	gr.AddVertex(4)
	gr.AddVertex(5)
	gr.AddVertex(6)
	gr.AddVertex(7)

	gr.AddEdgeByIDs(0, 2, 5)
	gr.AddEdgeByIDs(0, 3, 2)
	gr.AddEdgeByIDs(4, 6, 1)
	gr.AddEdgeByIDs(2, 0, 3)
	gr.AddEdgeByIDs(0, 1, 1)
	gr.AddEdgeByIDs(1, 6, 4)
	gr.AddEdgeByIDs(6, 3, 3)
	gr.AddEdgeByIDs(6, 5, 1)

	fmt.Println(gr)

	// gr.BFS(4)
}
