package graph

import "testing"

func TestCreateGraph(t *testing.T) {

	MGraph := CreateGraph(10)
	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     0,
		V2:     1,
	})
	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     3,
		V2:     4,
	})
	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     3,
		V2:     5,
	})

	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     3,
		V2:     6,
	})

	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     5,
		V2:     7,
	})

	PrintMGraphMatrix(MGraph)

	BFS(MGraph, 3)
}

func TestBFS(t *testing.T) {
	MGraph := CreateGraph(10)
	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     0,
		V2:     1,
	})
	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     3,
		V2:     2,
	})
	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     4,
		V2:     5,
	})

	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     2,
		V2:     5,
	})

	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     1,
		V2:     2,
	})

	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     7,
		V2:     4,
	})

	InsertEdge(MGraph, &ENode{
		Weight: 1,
		V1:     1,
		V2:     8,
	})

	PrintMGraphMatrix(MGraph)

	BFS(MGraph, 2)
}
