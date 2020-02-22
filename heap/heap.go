package heap

import (
	"fmt"
)

type ElementType int

type HNode struct {
	Data []ElementType
	Size int
}

type Heap *HNode

type MaxHeap Heap
type MinHeap Heap

// define MAX data 10000
func CreateHeap() MaxHeap {
	heap := &HNode{}
	heap.Data = make([]ElementType, 1001, 1001)
	heap.Data[0] = 10000
	return heap
}

func isFull(H MaxHeap) bool {
	return H.Size == cap(H.Data)
}

func Insert(H MaxHeap, E ElementType) bool {
	if isFull(H) {
		fmt.Println("堆已满")
		return false
	}

	H.Size += 1
	i := H.Size

	for ; H.Data[i/2] < E && i/2 >= 1; i /= 2 {
		H.Data[i] = H.Data[i/2] //向上过滤
	}

	H.Data[i] = E

	return true
}

func isEmpty(H MaxHeap) bool {
	return H.Size == 0
}

func PrintHeap(H Heap) {
	var parent, child int
	for parent = 1; parent*2 <= H.Size; parent++ {
		child = parent * 2
		if child != H.Size {
			fmt.Println("current node is :", H.Data[parent],
				" left node is :", H.Data[child], " right node is : ", H.Data[child+1])
		}
	}
}

func DeleteMax(H MaxHeap) (ElementType, error) {
	var Parent, Child int
	var MaxItem, X ElementType

	if isEmpty(H) {
		return -100000, fmt.Errorf("堆栈已经空了")
	}

	MaxItem = H.Data[1]
	X = H.Data[H.Size]
	H.Size -= 1

	for Parent = 1; Parent*2 <= H.Size; Parent = Child {
		Child = Parent * 2
		if Child != H.Size && H.Data[Child] < H.Data[Child+1] {
			Child += 1
		}
		if X >= H.Data[Child] {
			break //找到了合适的位置
		} else {
			H.Data[Parent] = H.Data[Child]
		}
	}

	H.Data[Parent] = X
	return MaxItem, nil
}

/* 建造最大堆 */
func PercDown(H MaxHeap, p int) {
	/* 下滤：将H中以H->Data[p]为根的子堆调整为最大堆 */
	var Parent, Child int
	var X ElementType
	X = H.Data[p] /* 取出根结点存放的值 */
	for Parent = p; Parent*2 <= H.Size; Parent = Child {
		Child = Parent * 2
		if Child != H.Size && H.Data[Child] < H.Data[Child+1] {
			Child += 1
		}
		if X > H.Data[Child] {
			break //找到了合适的位置
		} else {
			H.Data[Parent] = H.Data[Child]
		}
	}

	H.Data[Parent] = X
}

func BuildHeap(H MaxHeap) {
	/* 从最后一个结点的父节点开始，到根结点1 */
	for i := H.Size / 2; i > 0; i-- {
		PercDown(H, i)
	}
}
