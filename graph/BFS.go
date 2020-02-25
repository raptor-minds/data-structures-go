package graph

/* 邻接矩阵存储的图 - BFS */

/* IsEdge(Graph, V, W)检查<V, W>是否图Graph中的一条边，即W是否V的邻接点。  */
/* 此函数根据图的不同类型要做不同的实现，关键取决于对不存在的边的表示方法。*/
/* 例如对有权图, 如果不存在的边被初始化为INFINITY, 则函数实现如下:         */

type QueueData struct {
	data []Vertex
	size int
}

type Queue *QueueData

func CreateQueue(size int) Queue {
	return &QueueData{
		data: make([]Vertex, size),
		size: 0,
	}
}

func AddQ(queue Queue, v Vertex) {
	queue.data[queue.size] = v
	queue.size += 1
}

func IsEmpty(queue Queue) bool {
	return queue.size == 0
}

func DeleteQ(queue Queue) Vertex {
	queue.size -= 1
	v := queue.data[queue.size]
	queue.data[queue.size] = 0
	return v
}

func isEdge(graph MGraph, v Vertex, w Vertex) bool {
	if graph.Matrix[v][w] != 0 {
		return true
	}
	return false
}

func BFS(graph MGraph, v Vertex) {

	var X Vertex
	visited := make([]bool, graph.Nv)
	Q := CreateQueue(graph.Nv)
	visit(v)
	visited[v] = true
	AddQ(Q, v)
	for ; !IsEmpty(Q); {
		w := DeleteQ(Q)
		for X = 0; int(X) < graph.Nv; X++ {
			if !visited[X] && isEdge(graph, w, X) {
				visit(X)
				visited[X] = true
				AddQ(Q, X)
			}
		}
	}
}
