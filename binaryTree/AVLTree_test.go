package binaryTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVLTreeGetHeight(t *testing.T) {
	A := AVLNode{Data: 1}
	B := AVLNode{Data: 2}
	C := AVLNode{Data: 3}
	D := AVLNode{Data: 4}
	E := AVLNode{Data: 5, Left: &D, Right: &C}
	F := AVLNode{Data: 6, Left: &B, Right: &A}
	G := AVLNode{Data: 7, Left: &E, Right: &F}
	assert.Equal(t, 2, GetHeight(&G))
	assert.Equal(t, 0, GetHeight(&A))
	assert.Equal(t, 1, GetHeight(&F))
}

func TestSingleRightRotation(t *testing.T) {

}

func TestSingleLeftRotation(t *testing.T) {
	H := AVLNode{Data: 8}
	A := AVLNode{Data: 1, Left: &H}
	B := AVLNode{Data: 2, Left: &A}
	C := AVLNode{Data: 3}
	D := AVLNode{Data: 4}
	E := AVLNode{Data: 5, Left: &D, Right: &C}
	F := AVLNode{Data: 6, Left: &B}
	G := AVLNode{Data: 7, Left: &E, Right: &F}

	A.Height = GetHeight(&A)
	B.Height = GetHeight(&B)
	C.Height = GetHeight(&C)
	D.Height = GetHeight(&D)
	E.Height = GetHeight(&E)
	F.Height = GetHeight(&F)
	G.Height = GetHeight(&G)

	PrintAVLTree(&G)

	PrintAVLTree(SingleLeftRotation(&B))
}

func TestPrintAVLTree(t *testing.T) {
	A := AVLNode{Data: 1}
	B := AVLNode{Data: 2}
	C := AVLNode{Data: 3}
	D := AVLNode{Data: 4}
	E := AVLNode{Data: 5, Left: &D, Right: &C}
	F := AVLNode{Data: 6, Left: &B, Right: &A}
	G := AVLNode{Data: 7, Left: &E, Right: &F}
	PrintAVLTree(&G)
}

func TestInsert(t *testing.T) {
	A := &AVLNode{Data: 10}
	A = Insert(A, 8)
	A = Insert(A, 6)
	A = Insert(A, 5)
	A = Insert(A, 4)
	A = Insert(A, 3)
	A = Insert(A, 2)
	A = Insert(A, 1)
	//A = Insert(A, -1)

	PrintAVLTree(A)

	PrintAVLTree(Insert(A, 7))
}
