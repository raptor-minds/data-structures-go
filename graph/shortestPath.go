package graph

import "fmt"

const ERROR = -10

func UnWeighted(graph LGraph, dist []int, path []int, s Vertex) {

	var queue Queue
	var v Vertex
	var w PtrToAdjVNode

	queue = CreateQueue(graph.Nv)

	AddQ(queue, s)

	dist[s] = 0

	for ; !IsEmpty(queue); {
		v = DeleteQ(queue)
		visit(v)
		for w = graph.G[v].FistEdge; w != nil; w = w.Next {
			if dist[w.AdjV] == -1 { //not yet visited currently
				dist[w.AdjV] = dist[v] + 1
				path[w.AdjV] = int(v)
				AddQ(queue, w.AdjV)
			}
		}
	}
}

func Dijkstra(graph LGraph, dis []int, path []Vertex, v Vertex) {
	collection := make([]bool, graph.Nv+1)

	for i := 0; i < len(dis); i++ {
		dis[i] = INFINITY
		path[i] = -1
	}

	// initialize dis
	dis[v] = INFINITY
	for w := graph.G[v].FistEdge; w != nil; w = w.Next {
		dis[w.AdjV] = int(w.Weight)
		path[w.AdjV] = v
	}

	collection[v] = true
	collection[0] = true

	for i := 1; ; i++ {
		s := findMinDis(dis, collection) // 每次寻找最短的那个节点V
		if s == ERROR {
			break
		}

		collection[s] = true

		// 更新临接节点的到 v 的距离
		for w := graph.G[s].FistEdge; w != nil; w = w.Next {
			if !collection[w.AdjV] && dis[w.AdjV] > int(w.Weight)+dis[s] {
				dis[w.AdjV] = dis[s] + int(w.Weight)
				path[w.AdjV] = s
			}
		}
	}

	fmt.Print(path)
	fmt.Print(dis)

}

func Floyd(graph MGraph, D *[MaxVertexNum][MaxVertexNum]WeightType, path *[MaxVertexNum][MaxVertexNum]Vertex) {
	var i, j int
	for i = 0; i < graph.Nv; i++ {
		for j = 0; j < graph.Nv; j++ {
			D[i][j] = graph.Matrix[i][j]
			fmt.Print(D[i][j], " ")
			if i == j || D[i][j] != INFINITY {
				path[i][j] = 0
			} else {
				path[i][j] = -1
			}
		}
		fmt.Println()
	}

	for k := 0; k < graph.Nv; k++ {
		for i = 0; i < graph.Nv; i++ {
			for j = 0; j < graph.Nv; j++ {
				if D[i][j] > D[i][k]+D[k][j] {
					D[i][j] = D[i][k] + D[k][j]
					path[i][j] = Vertex(k)
				}
			}
		}
	}
}

func PrintPathForFloyd(path *[MaxVertexNum][MaxVertexNum]Vertex, src Vertex, des Vertex) {
	if path[src][des] == -1 {
		fmt.Println("there is no path between ", src, " and ", des, path[src][des])
		return
	}

	if path[src][des] == 0 {
		fmt.Print("--> next --->", des)
		return
	} else {
		k := path[src][des]
		fmt.Println("from ", src, " to " , des, " go through ", k)
		PrintPathForFloyd(path, src, k)
		PrintPathForFloyd(path, k, des)
	}
}

func findMinDis(dis []int, collection []bool) Vertex {
	minDis := INFINITY
	var minV Vertex = ERROR
	for i := 1; i < len(dis); i++ {
		if !collection[i] && dis[i] < minDis {
			minDis = dis[i]
			minV = Vertex(i)
		}
	}

	return minV
}

func PrintPath(path []int, src int, des int) {

	fmt.Print(des, "  is from -->")

	if path[des] == src {
		return
	}

	PrintPath(path, src, path[des])

}
