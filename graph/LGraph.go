package graph

import "fmt"

type AdjVNode struct {
	AdjV   Vertex //邻接节点下标
	Weight WeightType
	Next   PtrToAdjVNode
}

type PtrToAdjVNode *AdjVNode

type VNode struct {
	FistEdge PtrToAdjVNode
	Data     DataType
}

type AdjList [MaxVertexNum]VNode

type GAdjNode struct {
	Nv int //顶点数
	Ne int //边数
	G  AdjList
}

type LGraph *GAdjNode

func CreateLinkedGraph(vertexNum int) LGraph {
	graph := &GAdjNode{
		Nv: vertexNum,
		Ne: 0,
		G:  AdjList{},
	}
	return graph
}

func InsertLinkedEdge(Graph LGraph, E Edge) {
	newNode := AdjVNode{
		AdjV:   E.V2,
		Weight: E.Weight,
		Next:   Graph.G[E.V1].FistEdge,
	}

	Graph.G[E.V1].FistEdge = &newNode

	//如果是无向图
	node := AdjVNode{
		AdjV:   E.V1,
		Weight: E.Weight,
		Next:   Graph.G[E.V2].FistEdge,
	}

	Graph.G[E.V2].FistEdge = &node

	Graph.Ne++
}

func PrintLGraph(graph LGraph) {
	for i := 0; i < graph.Nv; i++ {
		fmt.Printf("%d :", i)
		for w := graph.G[i].FistEdge; w != nil; w = w.Next {
			fmt.Print("-->", w.AdjV)
		}
		fmt.Println()
	}
}
