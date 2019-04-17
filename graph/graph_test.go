package graph

import (
	"testing"
)

func TestDFS(t *testing.T) {
	graph := InitializeGraph()

	graph.AddVertex(0, 12)
	graph.AddVertex(1, 32)
	graph.AddVertex(2, 54)
	graph.AddVertex(3, 77)
	graph.AddVertex(4, 34)
	graph.AddVertex(5, 73)
	graph.AddVertex(6, 57)
	graph.AddVertex(7, 76)

	graph.AddEdgeByIDs(0, 2, 5)
	graph.AddEdgeByIDs(0, 3, 2)
	graph.AddEdgeByIDs(4, 6, 1)
	graph.AddEdgeByIDs(2, 0, 3)
	graph.AddEdgeByIDs(0, 1, 1)
	graph.AddEdgeByIDs(1, 6, 4)
	graph.AddEdgeByIDs(6, 3, 3)
	graph.AddEdgeByIDs(6, 5, 1)
	graph.AddEdgeByIDs(7, 3, 2)
	graph.AddEdgeByIDs(7, 5, 3)
	graph.AddEdgeByIDs(5, 7, 2)

	vertex := graph.DFS(57, graph.GetVertexByID(2))
	if vertex != graph.GetVertexByID(6) {
		t.Errorf("DFS test failed.\nResult: %v\nExpected: %v", vertex, 57)
	}
}

func TestBFS(t *testing.T) {
	graph := InitializeGraph()

	graph.AddVertex(0, 12)
	graph.AddVertex(1, 32)
	graph.AddVertex(2, 54)
	graph.AddVertex(3, 77)
	graph.AddVertex(4, 34)
	graph.AddVertex(5, 73)
	graph.AddVertex(6, 57)
	graph.AddVertex(7, 76)

	graph.AddEdgeByIDs(0, 2, 5)
	graph.AddEdgeByIDs(0, 3, 2)
	graph.AddEdgeByIDs(4, 6, 1)
	graph.AddEdgeByIDs(2, 0, 3)
	graph.AddEdgeByIDs(0, 1, 1)
	graph.AddEdgeByIDs(1, 6, 4)
	graph.AddEdgeByIDs(6, 3, 3)
	graph.AddEdgeByIDs(6, 5, 1)
	graph.AddEdgeByIDs(7, 3, 2)
	graph.AddEdgeByIDs(7, 5, 3)
	graph.AddEdgeByIDs(5, 7, 2)

	vertex := graph.BFS(57, graph.GetVertexByID(2))
	if vertex != graph.GetVertexByID(6) {
		t.Errorf("DFS test failed.\nResult: %v\nExpected: %v", vertex, 57)
	}
}
