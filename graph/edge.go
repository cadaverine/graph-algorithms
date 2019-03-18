package graph

import "fmt"

// Edge - ребро графа
type Edge struct {
	weight int
	from   *Vertex
	to     *Vertex
}

func (edge *Edge) String() string {
	return fmt.Sprint(edge.from) + " -(" + fmt.Sprint(edge.weight) + ")-> " + fmt.Sprint(edge.to)
}
