package heap

import (
	"algorithm/pkg/sort"
	"fmt"
	"testing"
)

func TestHeap_SortAsc(t *testing.T) {
	arr := []int{5, 1, 15, 3, 4, 3, 6, 0, 2, 1}
	i := New(arr)
	i.SortAsc()
	fmt.Println(i.GetArr())
	fmt.Println(sort.IsAscSorted(i.Array))
}

func TestHeap_SortDesc(t *testing.T) {
	arr := []int{5, 1, 15, 3, 4, 3, 6, 0, 2, 1}
	i := New(arr)
	i.SortDesc()
	fmt.Println(i.GetArr())
	fmt.Println(sort.IsDescSorted(i.Array))
}
