package heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCreateHeap2(t *testing.T) {

	H := HuffmanNode{
		Data: make([]TreeNode, 10),
		Size: 10,
	}

	H.Data[0] = TreeNode{
		Weight: 0,
		Left:   nil,
		Right:  nil,
	}

	fmt.Println(len(H.Data))
	fmt.Println(cap(H.Data))
}

func TestCreateHeap3(t *testing.T) {
	type NewHuffmanTree struct {
		Data [10]TreeNode
		Size int
	}

	H := &NewHuffmanTree{
		Data: [10]TreeNode{},
		Size: 0,
	}

	H.Data[1] = TreeNode{
		Weight: 1,
		Left:   nil,
		Right:  nil,
	}

	fmt.Println(len(H.Data))
	fmt.Println(H.Data[0], H.Data[1])
}

func TestCpNode(t *testing.T) {
	X := TreeNode{
		Weight: 8,
		Left:   nil,
		Right:  nil,
	}
	Y := X
	fmt.Println(X, Y)
	X.Weight = 9
	fmt.Println(X, Y)
}

func TestBuildHuffmanTree(t *testing.T) {
	huffmanArray := &HuffmanNode{
		Data: make([]TreeNode, 11),
		Size: 10,
	}

	for i:= 1; i< 11; i++ {
		T := &TreeNode{
			Weight: rand.Intn(100),
			Left:   nil,
			Right:  nil,
		}
		huffmanArray.Data[i] = *T
	}

	fmt.Println(&huffmanArray.Data)

	H := Huffman(huffmanArray)

	PrintHuffmanTree(H)
}
