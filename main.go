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

	gr.AddEdgeData(1, 3, 5)
	gr.AddEdgeData(1, 4, 2)
	gr.AddEdgeData(5, 7, 1)
	gr.AddEdgeData(3, 1, 3)

	fmt.Println(gr)
}
