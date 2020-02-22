package heap

import "fmt"

type HuffmanTree *TreeNode

type TreeNode struct {
	Weight      int
	Left, Right HuffmanTree
}

type HuffmanNode struct {
	Data []TreeNode
	Size int
}

type HuffmanMinHeap *HuffmanNode

func BuildMinHeap(H HuffmanMinHeap) {
	for i := H.Size / 2; i > 0; i-- {
		PerDownForMinHeap(H, i)
	}
}

func PerDownForMinHeap(H HuffmanMinHeap, p int) {
	var Parent, Child int
	X := H.Data[p] //find the root point
	for Parent = p; Parent*2 <= H.Size; Parent = Child {
		Child = Parent * 2
		if Child != H.Size && H.Data[Child+1].Weight < H.Data[Child].Weight {
			Child += 1
		}

		if X.Weight <= H.Data[Child].Weight {
			break
		} else {
			H.Data[Parent] = H.Data[Child]
		}
	}

	H.Data[Parent] = X
}

func isHuffmanFull(H HuffmanMinHeap) bool {
	return H.Size == cap(H.Data)
}

func InsertHuffmanMinHeap(H HuffmanMinHeap, X TreeNode) bool {
	if isHuffmanFull(H) {
		fmt.Println("堆已满")
		return false
	}

	H.Size += 1
	i := H.Size

	for ; H.Data[i/2].Weight > X.Weight && i != 1; i /= 2 {
		H.Data[i] = H.Data[i/2]
	}

	H.Data[i] = X
	return true
}

func isHuffmanEmpty(H HuffmanMinHeap) bool {
	return H.Size == 0
}

func DeleteMin(H HuffmanMinHeap) (TreeNode, error) {
	var Parent, Child int
	var MinItem, X TreeNode

	if isHuffmanEmpty(H) {
		return TreeNode{}, fmt.Errorf("堆栈已经空了")
	}

	MinItem = H.Data[1]
	X = H.Data[H.Size]
	H.Size -= 1
	for Parent = 1; Parent*2 <= H.Size; Parent = Child {
		Child = Parent * 2
		if Child != H.Size && H.Data[Child+1].Weight < H.Data[Child].Weight {
			Child += 1
		}

		if X.Weight <= H.Data[Child].Weight {
			break
		} else {
			H.Data[Parent] = H.Data[Child]
		}
	}

	H.Data[Parent] = X
	fmt.Println(H.Data[1])
	return MinItem, nil
}

func Huffman(H HuffmanMinHeap) HuffmanTree {
	BuildMinHeap(H)
	total := H.Size
	for i := 1; i < total; i++ {
		Left, _ := DeleteMin(H)
		//fmt.Printf("left %p", &Left)
		Right, _ := DeleteMin(H)
		//fmt.Printf("right %p", &Right)
		T := TreeNode{
			Weight: Left.Weight + Right.Weight,
			Left:   &Left,
			Right:  &Right,
		}
		InsertHuffmanMinHeap(H, T)
		fmt.Println(H, H.Size)
	}
	Min, _ := DeleteMin(H)
	return &Min
}

func PrintHuffmanTree(H HuffmanTree) {
	fmt.Println(H.Weight)
	if H.Left != nil {
		fmt.Println(H.Weight, ": left node is", H.Left.Weight)
		PrintHuffmanTree(H.Left)
	}

	if H.Right != nil {
		fmt.Println(H.Weight, ": right node is", H.Right.Weight)
		PrintHuffmanTree(H.Right)
	}
}
