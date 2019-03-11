package graph

type Vertex struct {
	data interface{}
}

type Edge struct {
	weight int
	from   *Vertex
	to     *Vertex
	next   *Edge
}

func (edge *Edge) Next() interface{} {
	return edge.next
}

type EdgeList struct {
	head *Edge
}

// Graph - реализация графа в виде списка смежности
type Graph struct {
	edgeList []EdgeList
}
