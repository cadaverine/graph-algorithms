package main

import (
	"fmt"

	"./graph"
)

func main() {
	gr := graph.CreateFromString(`
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

	fmt.Println(gr)

	fmt.Println("distance between 4 and 3: ", gr.Dijkstra(gr.GetVertexByID(4), gr.GetVertexByID(3)))
	fmt.Println("distance between 1 and 3: ", gr.Dijkstra(gr.GetVertexByID(1), gr.GetVertexByID(3)))
	fmt.Println("distance between 3 and 2: ", gr.Dijkstra(gr.GetVertexByID(3), gr.GetVertexByID(2)))
}
