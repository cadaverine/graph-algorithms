package graph

import "fmt"

// Vertex - вершина графа
type Vertex struct {
	data interface{}
}

func (vertex *Vertex) String() string {
	return fmt.Sprint(vertex.data)
}
