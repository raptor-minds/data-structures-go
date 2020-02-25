package graph

import "fmt"

func visit(v Vertex) {
	fmt.Println("visited ", v)
}

func DFS(graph LGraph, v Vertex) {
	visit(v)
	graph.G[v].Data = true // data store the visited message
	for w := graph.G[v].FistEdge; w != nil; w = w.Next {
		if !graph.G[w.AdjV].Data {
			DFS(graph, w.AdjV)
		}
	}
}
