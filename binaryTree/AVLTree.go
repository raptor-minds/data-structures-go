package binaryTree

import "fmt"

type ElementType int

type Position *AVLNode
type AVLTree Position

type AVLNode struct {
	Data   ElementType
	Left   AVLTree
	Right  AVLTree
	Height int
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func SingleLeftRotation(A AVLTree) AVLTree {
	// A must have a left child node
	// update A and B's Height
	B := A.Left
	A.Left = B.Right
	B.Right = A
	A.Height = Max(GetHeight(A.Left), GetHeight(A.Right)) + 1
	B.Height = Max(GetHeight(B.Left), A.Height) + 1
	return B
}

func DoubleLeftRightRotation(A AVLTree) AVLTree {
	// 先将B与C做右边单旋
	A.Left = SingleRightRotation(A.Left)
	// 再将A与C做左单旋
	return SingleLeftRotation(A)
}

func DoubleRightLeftRotation(A AVLTree) AVLTree {
	// 先将B与C做左单旋
	A.Right = SingleLeftRotation(A.Right)
	// 再将A与C做右单旋
	return SingleRightRotation(A)
}

func SingleRightRotation(A AVLTree) AVLTree {
	// A must have a right child node
	B := A.Right
	A.Right = B.Left
	B.Left = A
	A.Height = Max(GetHeight(A.Left), GetHeight(A.Right)) + 1
	B.Height = Max(GetHeight(B.Right), A.Height) + 1
	return B
}

func GetHeight(t AVLTree) int {
	if t == nil {
		return 0
	}
	if t.Left == nil && t.Right == nil {
		return 0
	} else if t.Left != nil {
		return GetHeight(t.Left) + 1
	} else {
		return GetHeight(t.Right) + 1
	}
}

func Insert(T AVLTree, E ElementType) AVLTree {
	if T == nil {
		T = &AVLNode{Data: E}
	}

	if E < T.Data {
		T.Left = Insert(T.Left, E)
		if GetHeight(T.Left) - GetHeight(T.Right) == 2 {
			if E < T.Left.Data {
				T = SingleLeftRotation(T)
			} else {
				T = DoubleLeftRightRotation(T)
			}
		}
	}

	if E > T.Data {
		T.Right = Insert(T.Right, E)
		if GetHeight(T.Left) - GetHeight(T.Right) == -2 {
			if E > T.Right.Data {
				T = SingleRightRotation(T)
			} else {
				T = DoubleRightLeftRotation(T)
			}
		}
	}

	T.Height = Max(GetHeight(T.Left), GetHeight(T.Right)) + 1

	return T
}

func PrintAVLTree(A AVLTree) {
	fmt.Println(A.Data, ", Height is :", A.Height)

	if A.Left != nil {
		fmt.Print(A.Data, ": left is : ")
		PrintAVLTree(A.Left)
	}

	if A.Right != nil {
		fmt.Print(A.Data, ": right is: ")
		PrintAVLTree(A.Right)
	}
}
