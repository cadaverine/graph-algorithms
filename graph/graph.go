package graph

import . "../list"

type Vertex struct {
	data interface{}
}

type Edge struct {
	weight int
	from   *Vertex
	to     *Vertex
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
