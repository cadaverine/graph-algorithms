package graph

import (
	"testing"
)

func TestDFS(t *testing.T) {
	graph := CreateFromString(`
		{
			(0, 12),
			(1, 32),
			(2, 54),
			(3, 77),
			(4, 34),
			(5, 73),
			(6, 57),
			(7, 76)
		},
		{
			(0, 2, 5),
			(0, 3, 2),
			(4, 6, 1),
			(2, 0, 3),
			(0, 1, 1),
			(1, 6, 4),
			(6, 3, 3),
			(6, 5, 1),
			(7, 3, 2),
			(7, 5, 3),
			(5, 7, 2)
		}
	`)

	vertex := graph.DFS(57, graph.GetVertexByID(2))
	if vertex != graph.GetVertexByID(6) {
		t.Errorf("DFS test failed.\nResult: %v\nExpected: %v", vertex, 57)
	}

	vertex = graph.DFS(34, graph.GetVertexByID(2))
	if vertex != nil {
		t.Errorf("DFS test failed.\nResult: %v\nExpected: %v", vertex, nil)
	}
}

func TestBFS(t *testing.T) {
	graph := CreateFromString(`
		{
			(0, 12),
			(1, 32),
			(2, 54),
			(3, 77),
			(4, 34),
			(5, 73),
			(6, 57),
			(7, 76)
		},
		{
			(0, 2, 5),
			(0, 3, 2),
			(4, 6, 1),
			(2, 0, 3),
			(0, 1, 1),
			(1, 6, 4),
			(6, 3, 3),
			(6, 5, 1),
			(7, 3, 2),
			(7, 5, 3),
			(5, 7, 2)
		}
	`)

	vertex := graph.BFS(57, graph.GetVertexByID(2))
	if vertex != graph.GetVertexByID(6) {
		t.Errorf("DFS test failed.\nResult: %v\nExpected: %v", vertex, 57)
	}
}

func TestDijkstra(t *testing.T) {
	graph := CreateFromString(`
		{
			(0, 12),
			(1, 32),
			(2, 54),
			(3, 77),
			(4, 34),
			(5, 73)
		},
		{
			(0, 1, 3),
			(1, 5, 4),
			(0, 2, 2),
			(2, 3, 1),
			(3, 4, 1),
			(4, 5, 1)
		}
	`)

	minDistance := graph.Dijkstra(graph.GetVertexByID(0), graph.GetVertexByID(5))
	if minDistance != 5 {
		t.Errorf("DFS test failed.\nResult: %v\nExpected: %v", minDistance, 5)
	}
}
