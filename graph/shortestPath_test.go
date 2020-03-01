package graph

import (
	"fmt"
	"testing"
)

func TestUnWeighted(t *testing.T) {

	LGraph := CreateLinkedGraph(7)

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 2,
		V1:     1,
		V2:     2,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 3,
		V1:     2,
		V2:     4,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 10,
		V1:     2,
		V2:     5,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 2,
		V1:     4,
		V2:     5,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 4,
		V1:     4,
		V2:     7,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 2,
		V1:     4,
		V2:     2,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 8,
		V1:     4,
		V2:     6,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 5,
		V1:     3,
		V2:     6,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 6,
		V1:     5,
		V2:     7,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     7,
		V2:     6,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     1,
		V2:     4,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 4,
		V1:     3,
		V2:     1,
	})

	PrintLGraph(LGraph)

	dist := make([]int, LGraph.Nv+1)

	for i := 1; i < len(dist); i++ {
		dist[i] = -1
	}

	path := make([]int, LGraph.Nv+1)

	for i := 1; i < len(path); i++ {
		path[i] = -1
	}

	index := make([]int, LGraph.Nv+1)

	for i := 0; i < len(dist); i++ {
		index[i] = i
	}

	fmt.Print(index, " ")
	fmt.Println("before")
	fmt.Print(dist)
	fmt.Println()
	fmt.Print(path)

	UnWeighted(LGraph, dist, path, 2)

	fmt.Print(index, " ")
	fmt.Println("after")
	fmt.Print(dist)
	fmt.Println()
	fmt.Print(path)

	PrintPath(path, 2, 6)
}

func TestDijkstra(t *testing.T) {
	LGraph := CreateLinkedGraph(7)

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 2,
		V1:     1,
		V2:     2,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 3,
		V1:     2,
		V2:     4,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 10,
		V1:     2,
		V2:     5,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 2,
		V1:     4,
		V2:     5,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 4,
		V1:     4,
		V2:     7,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 8,
		V1:     4,
		V2:     6,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 5,
		V1:     3,
		V2:     6,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 6,
		V1:     5,
		V2:     7,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     7,
		V2:     6,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 1,
		V1:     1,
		V2:     4,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 4,
		V1:     3,
		V2:     1,
	})

	InsertLinkedEdge(LGraph, &ENode{
		Weight: 2,
		V1:     4,
		V2:     3,
	})

	PrintLGraph(LGraph)

	index := make([]int, LGraph.Nv+1)

	for i := 0; i < len(index); i++ {
		index[i] = i
	}

	dist := make([]int, LGraph.Nv+1)

	path := make([]Vertex, LGraph.Nv+1)

	Dijkstra(LGraph, dist, path, 1)

	fmt.Print(path)
	fmt.Print(dist)
}

func TestFloyd(t *testing.T) {
	mGraph := CreateGraph(10)
	InsertEdge(mGraph, &ENode{
		Weight: 10,
		V1:     0,
		V2:     6,
	})

	InsertEdge(mGraph, &ENode{
		Weight: 2,
		V1:     0,
		V2:     4,
	})

	InsertEdge(mGraph, &ENode{
		Weight: 1,
		V1:     4,
		V2:     5,
	})

	InsertEdge(mGraph, &ENode{
		Weight: 3,
		V1:     4,
		V2:     6,
	})

	InsertEdge(mGraph, &ENode{
		Weight: 1,
		V1:     5,
		V2:     6,
	})

	PrintMGraphMatrix(mGraph)

	D := [MaxVertexNum][MaxVertexNum]WeightType{}

	path := [MaxVertexNum][MaxVertexNum]Vertex{}
	Floyd(mGraph, &D, &path)

	for i := 0; i < MaxVertexNum; i++ {
		for j := 0; j < MaxVertexNum; j++ {
			fmt.Print(path[i][j], " ")
		}
		fmt.Println()
	}

	for i := 0; i < MaxVertexNum; i++ {
		for j := 0; j < MaxVertexNum; j++ {
			fmt.Print(D[i][j], " ")
		}
		fmt.Println()
	}

	PrintPathForFloyd(&path, 0, 6)
}
