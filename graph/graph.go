package graph

import (
	"fmt"

	. "../list"
)

// Vertex - вершина графа
type Vertex struct {
	data interface{}
}

func (vertex *Vertex) String() string {
	return fmt.Sprint(vertex.data)
}

// Edge - ребро графа
type Edge struct {
	weight int
	from   *Vertex
	to     *Vertex
}

func (edge *Edge) String() string {
	return fmt.Sprint(edge.from) + " -> " + fmt.Sprint(edge.to)
}

// Graph - реализация графа в виде списка смежности
type Graph struct {
	adjacencyMap map[*Vertex]*List
}

// InitializeGraph - создание графа (аналог конструктора
func InitializeGraph() *Graph {
	return &Graph{make(map[*Vertex]*List, 0)}
}

// AddVertex - создать вершину из данных и добавить в граф
func (graph *Graph) AddVertex(data interface{}) {
	vertex := Vertex{data}
	graph.adjacencyMap[&vertex] = &List{}
}

// AddEdge - создать ребро и добавить в граф
func (graph *Graph) AddEdge(from, to *Vertex, weight int) {
	edge := Edge{weight, from, to}

	if _, keyExist := graph.adjacencyMap[from]; keyExist {
		graph.adjacencyMap[from].AddData(&edge)
	} else {
		graph.adjacencyMap[from] = &List{}
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

func (graph Graph) String() (out string) {
	for key, value := range graph.adjacencyMap {
		out += fmt.Sprint(key) + ": " + fmt.Sprint(value) + "\n"
	}

	return out
}
