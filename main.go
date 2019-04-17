package main

import (
	"fmt"

	"./graph"
)

func main() {
	gr := graph.InitializeGraph()

	gr.AddVertex(0, 12)
	gr.AddVertex(1, 32)
	gr.AddVertex(2, 54)
	gr.AddVertex(3, 77)
	gr.AddVertex(4, 34)
	gr.AddVertex(5, 73)
	gr.AddVertex(6, 57)
	gr.AddVertex(7, 76)

	gr.AddEdgeByIDs(0, 2, 5)
	gr.AddEdgeByIDs(0, 3, 2)
	gr.AddEdgeByIDs(4, 6, 1)
	gr.AddEdgeByIDs(2, 0, 3)
	gr.AddEdgeByIDs(0, 1, 1)
	gr.AddEdgeByIDs(1, 6, 4)
	gr.AddEdgeByIDs(6, 3, 3)
	gr.AddEdgeByIDs(6, 5, 1)
	gr.AddEdgeByIDs(7, 3, 2)
	gr.AddEdgeByIDs(7, 5, 3)
	gr.AddEdgeByIDs(5, 7, 2)

	fmt.Println(gr)

	fmt.Println("distance between 4 and 3: ", gr.Dijkstra(gr.GetVertexByID(4), gr.GetVertexByID(3)))
	fmt.Println("distance between 1 and 3: ", gr.Dijkstra(gr.GetVertexByID(1), gr.GetVertexByID(3)))
	fmt.Println("distance between 3 and 2: ", gr.Dijkstra(gr.GetVertexByID(3), gr.GetVertexByID(2)))
}
