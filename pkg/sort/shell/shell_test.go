package shell

import (
	"algorithm/pkg/sort"
	"fmt"
	"testing"
)

func TestShell_SortDesc(t *testing.T) {
	arr := []int{5, 1, 15, 3, 4, 3, 6, 0, 2, 1}
	s := New(arr)
	s.SortDesc()
	fmt.Println(s.GetArr())
	fmt.Println(sort.IsDescSorted(s.Array))
}

func TestShell_SortAsc(t *testing.T) {
	arr := []int{5, 1, 15, 3, 4, 3, 6, 0, 2, 1}
	s := New(arr)
	s.SortAsc()
	fmt.Println(s.GetArr())
	fmt.Println(sort.IsAscSorted(s.Array))
}
