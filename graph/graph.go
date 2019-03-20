package graph

import (
	"fmt"

	"../list"
	"../queue"
	"../stack"
)

// Graph - реализация графа в виде списка смежности
type Graph struct {
	adjacencyMap map[*Vertex]*list.List
	verteces     []*Vertex
}

// InitializeGraph - создание графа (аналог конструктора)
func InitializeGraph() *Graph {
	return &Graph{
		adjacencyMap: make(map[*Vertex]*list.List, 0),
		verteces:     make([]*Vertex, 0),
	}
}

// AddVertex - создать вершину из данных и добавить в граф
func (graph *Graph) AddVertex(data interface{}) *Vertex {
	id := len(graph.verteces)

	vertex := &Vertex{data, id}

	graph.verteces = append(graph.verteces, vertex)
	graph.adjacencyMap[vertex] = &list.List{}

	return vertex
}

// GetVertexByID - получить вершину графа по его ID
func (graph *Graph) GetVertexByID(id int) *Vertex {
	return graph.verteces[id]
}

// AddEdge - создать ребро и добавить в граф
func (graph *Graph) AddEdge(from, to *Vertex, weight int) {
	edge := &Edge{weight, from, to}

	if _, keyExist := graph.adjacencyMap[from]; keyExist {
		graph.adjacencyMap[from].AddData(edge)
	} else {
		graph.adjacencyMap[from] = &list.List{}
		graph.adjacencyMap[from].AddData(edge)
	}
}

// AddEdgeByData - создать ребро из данных и добавить в граф
func (graph *Graph) AddEdgeByData(fromData, toData interface{}, weight int) {
	to := graph.AddVertex(toData)
	from := graph.AddVertex(fromData)

	graph.AddEdge(from, to, weight)
}

// AddEdgeByIDs - создать ребро из между существующими вершинами
func (graph *Graph) AddEdgeByIDs(fromID int, toID int, weight int) {
	to := graph.GetVertexByID(toID)
	from := graph.GetVertexByID(fromID)

	graph.AddEdge(from, to, weight)
}

// String - реализация интерфейса Stringer
func (graph *Graph) String() (out string) {
	for key, value := range graph.adjacencyMap {
		out += fmt.Sprint(key) + ": " + fmt.Sprint(value) + "\n"
	}

	return out
}

// BFS (breadth-first search) - поиск в ширину с использованием очереди
func (graph *Graph) BFS(data interface{}, from *Vertex) *Vertex {
	visited := make(map[*Vertex]bool)

	toVisit := queue.Queue{}
	toVisit.Enqueue(from)

	for toVisit.Size() != 0 {
		value, _ := toVisit.Dequeue()
		currentVertex := value.(*Vertex)

		if visited[currentVertex] {
			continue
		}

		visited[currentVertex] = true

		if currentVertex.data == data {
			return currentVertex
		}

		edges := graph.adjacencyMap[currentVertex]
		currentEdge := edges.Head

		for currentEdge != nil {
			toVisit.Enqueue(currentEdge.Data.(*Edge).to)
			currentEdge = currentEdge.Next
		}
	}

	return nil
}

// DFS (depth-first search) - поиск в глубину с использованием стека
func (graph *Graph) DFS(data interface{}, from *Vertex) *Vertex {
	visited := make(map[*Vertex]bool)

	toVisit := stack.Stack{}
	toVisit.Push(from)

	for toVisit.Size() != 0 {
		value, _ := toVisit.Pop()
		currentVertex := value.(*Vertex)

		if visited[currentVertex] {
			continue
		}

		visited[currentVertex] = true

		if currentVertex.data == data {
			return currentVertex
		}

		edges := graph.adjacencyMap[currentVertex]
		currentEdge := edges.Head

		for currentEdge != nil {
			toVisit.Push(currentEdge.Data.(*Edge).to)
			currentEdge = currentEdge.Next
		}
	}

	return nil
}
