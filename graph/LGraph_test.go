package graph

import "testing"

func TestCreateLinkedGraph(t *testing.T) {

	LGraph := CreateLinkedGraph(10)

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     1,
		V2:     2,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     1,
		V2:     6,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     2,
		V2:     3,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     2,
		V2:     5,
	})
	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     2,
		V2:     7,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     8,
		V2:     9,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     6,
		V2:     9,
	})

	PrintLGraph(LGraph)

	DFS(LGraph, 2)
}

func TestDFS(t *testing.T) {
	LGraph := CreateLinkedGraph(10)

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     2,
		V2:     5,
	})
	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     2,
		V2:     7,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     8,
		V2:     9,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     6,
		V2:     9,
	})

	PrintLGraph(LGraph)

	DFS(LGraph, 2)
}