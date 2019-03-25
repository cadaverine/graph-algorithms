package graph

import "fmt"

// Vertex - вершина графа
type Vertex struct {
	data interface{}
	id   int
}

// GetData - получение данных из вершины
func (vertex *Vertex) GetData() interface{} {
	return vertex.data
}

func (vertex *Vertex) String() string {
	return fmt.Sprint(vertex.id)
}
