package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCreateHeap(t *testing.T) {
	H := CreateHeap()
	assert.Equal(t,0, H.Size)
}

func TestInsert(t *testing.T) {
	H := CreateHeap()
	for i:= 1; i <= 20; i++ {
		Insert(H, ElementType(rand.Intn(1000)))
	}
	fmt.Println(H.Size)
	PrintHeap(Heap(H))
	for i:=1 ; i<= 20; i++ {
		X, err := DeleteMax(H)
		if err != nil {
			fmt.Errorf("err %s", err)
		}
		fmt.Print(X, " ")
	}
	assert.Equal(t, 0, H.Size)
}

func testArray(array *[]int)  {
	fmt.Printf("%p", array)
}

func TestArray(t *testing.T) {
	a := make([]int, 5)
	for i:=0; i< 5; i++ {
		a[i] = i
	}
	fmt.Printf("%p", &a)
	testArray(&a)
}