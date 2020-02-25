package graph

import "fmt"

const MaxVertexNum = 100 /* 最大顶点数设为100 */
const INFINITY = 65535

type Vertex int
type WeightType int
type DataType bool

type ENode struct {
	Weight WeightType //权重
	V1, V2 Vertex     //有向边
}

type PtrToENode *ENode
type Edge PtrToENode

type GNode struct {
	Nv     int                                    //顶点数
	Ne     int                                    //边数
	Matrix [MaxVertexNum][MaxVertexNum]WeightType //临接矩阵
	Data   []DataType                             //顶点数据
}

type PtrToGNode *GNode
type MGraph PtrToGNode

func CreateGraph(vertexNum int) MGraph {
	// build a graph only with nodes without edges
	graph := &GNode{
		Nv:   vertexNum,
		Ne:   0,
		Data: nil,
	}

	/* 初始化邻接矩阵 */
	/* 注意：这里默认顶点编号从0开始，到(Graph->Nv - 1) */
	for v := 0; v < graph.Nv; v++ {
		for w := 0; w < graph.Nv; w++ {
			graph.Matrix[v][w] = 0
		}
	}

	return graph
}

func InsertEdge(G MGraph, E Edge) {
	// insert edge
	G.Matrix[E.V1][E.V2] = E.Weight
	//无向图
	G.Matrix[E.V2][E.V1] = E.Weight

	G.Ne++
}

func PrintMGraphMatrix(G MGraph) {
	matrix := G.Matrix
	fmt.Println(G.Nv, G.Ne)
	fmt.Print("G : ")
	for i := 0; i < G.Nv; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for i := 0; i < G.Nv; i++ {
		fmt.Printf("%d : ", i)
		for j := 0; j < G.Nv; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
