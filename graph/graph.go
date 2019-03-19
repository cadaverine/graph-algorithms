package graph

import (
	"fmt"

	"../list"
	"../queue"
)

// Graph - реализация графа в виде списка смежности
type Graph struct {
	adjacencyMap map[*Vertex]*list.List
}

// InitializeGraph - создание графа (аналог конструктора)
func InitializeGraph() *Graph {
	return &Graph{make(map[*Vertex]*list.List, 0)}
}

// AddVertex - создать вершину из данных и добавить в граф
func (graph *Graph) AddVertex(data interface{}) {
	vertex := Vertex{data}
	graph.adjacencyMap[&vertex] = &list.List{}
}

// AddEdge - создать ребро и добавить в граф
func (graph *Graph) AddEdge(from, to *Vertex, weight int) {
	edge := Edge{weight, from, to}

	if _, keyExist := graph.adjacencyMap[from]; keyExist {
		graph.adjacencyMap[from].AddData(&edge)
	} else {
		graph.adjacencyMap[from] = &list.List{}
		graph.adjacencyMap[from].AddData(&edge)
	}
}

// AddEdgeData - создать ребро из данных и добавить в граф
func (graph *Graph) AddEdgeData(fromData, toData interface{}, weight int) {
	for from := range graph.adjacencyMap {
		if from.data == fromData {
			to := Vertex{toData}

			graph.AddEdge(from, &to, weight)
			return
		}
	}

	to := Vertex{toData}
	from := Vertex{fromData}

	graph.AddEdge(&from, &to, weight)
}

func (graph *Graph) String() (out string) {
	for key, value := range graph.adjacencyMap {
		out += fmt.Sprint(key) + ": " + fmt.Sprint(value) + "\n"
	}

	return out
}

// DFS (depth-first search) - поиск в глубину
// func (graph *Graph) DFS(data interface{}, from *Vertex) *Vertex {
// 	list := graph.adjacencyMap[from]
// 	stack := stack.Stack{}
// }

// BFS (breadth-first search) - поиск в ширину
func (graph *Graph) BFS(data interface{}, from *Vertex) *Vertex {
	edges := graph.adjacencyMap[from]
	visited := make(map[*Vertex]bool)

	toVisit := queue.Queue{}
	toVisit.Enqueue(from)

	for toVisit.Size() != 0 {
		value, _ := toVisit.Dequeue()
		currentVertex := value.(Vertex)

		visited[&currentVertex] = true

		fmt.Println(currentVertex)

		if currentVertex.data == data {
			return &currentVertex
		}

		currentEdge := edges.Head
		for currentEdge != nil {
			toVisit.Enqueue(currentEdge.Data.(Edge).to)
			currentEdge = currentEdge.Next
		}
	}

	return nil
}
