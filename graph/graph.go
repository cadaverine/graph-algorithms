package graph

import (
	"fmt"
	"math"

	"../heap"
	"../list"
	"../queue"
	"../stack"
)

// Graph - реализация графа в виде списка смежности
type Graph struct {
	adjacencyMap map[*Vertex]*list.List
	verteces     map[int]*Vertex
}

// InitializeGraph - создание графа (аналог конструктора)
func InitializeGraph() *Graph {
	return &Graph{
		adjacencyMap: make(map[*Vertex]*list.List, 0),
		verteces:     make(map[int]*Vertex, 0),
	}
}

// AddVertex - создать вершину из данных и добавить в граф
func (graph *Graph) AddVertex(id int, data interface{}) *Vertex {
	vertex := &Vertex{data, id}

	graph.verteces[id] = vertex
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

// Dijkstra - алгоритм Дейкстры поиска кратчайшего пути в графе
func (graph *Graph) Dijkstra(from *Vertex, to *Vertex) int {
	const infinity int = math.MaxInt32

	distances := make(map[*Vertex]int, len(graph.verteces))
	paths := make(map[*Vertex]*list.List, len(graph.verteces))

	for _, vertex := range graph.verteces {
		distances[vertex] = infinity
	}
	distances[from] = 0

	priorityQueue := heap.Init(heap.Minimum)
	priorityQueue.Enqueue(from, 0)

	for priorityQueue.Size() != 0 {
		element, _ := priorityQueue.Dequeue()
		vertex := element.(*Vertex)

		edgesList := graph.adjacencyMap[vertex]
		edgesListItem := edgesList.Head

		for edgesListItem != nil {
			edge := edgesListItem.Data.(*Edge)
			currentDistance := distances[edge.from] + edge.weight
			paths[edge.to] = &list.List{}

			if distances[edge.to] > currentDistance {
				distances[edge.to] = currentDistance
				priorityQueue.Enqueue(edge.to, currentDistance)
				paths[edge.to].AddData(edge)
			}

			edgesListItem = edgesListItem.Next
		}

	}

	fmt.Println()
	for vertex, path := range paths {
		fmt.Println(vertex, ": ", path)
	}

	return distances[to]
}
