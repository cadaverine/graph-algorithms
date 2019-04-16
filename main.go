package main

import (
	"fmt"

	"./graph"
	"./heap"
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

	vertexBFS := gr.BFS(57, gr.GetVertexByID(2))
	vertexDFS := gr.DFS(57, gr.GetVertexByID(2))
	distance := gr.Dijkstra(gr.GetVertexByID(4), gr.GetVertexByID(3))

	resultBFS := fmt.Sprint(vertexBFS) + ", data: " + fmt.Sprint(vertexBFS.GetData()) + ")"
	resultDFS := fmt.Sprint(vertexDFS) + ", data: " + fmt.Sprint(vertexDFS.GetData()) + ")"

	fmt.Println("BFS result: ", resultBFS)
	fmt.Println("DFS result: ", resultDFS)
	fmt.Println("distance: ", distance)

	fmt.Println()

	hp := heap.Init(heap.Maximum)

	hp.Enqueue(0, 5)
	hp.Enqueue(0, 6)
	hp.Enqueue(0, 2)
	hp.Enqueue(0, 3)
	hp.Enqueue(0, 0)

	fmt.Println(hp)
	fmt.Println(hp.Dequeue())
	fmt.Println(hp.Dequeue())
	fmt.Println(hp.Dequeue())
	fmt.Println(hp.Dequeue())
	fmt.Println(hp.Dequeue())
}
