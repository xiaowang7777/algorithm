package insertion

import "algorithm/pkg/sort"

type insertion struct {
	sort.Array
}

func New(arr []int) *insertion {
	return &insertion{arr}
}

func (i insertion) SortDesc() {
	l := len(i.Array)
	for j := 0; j < l; j++ {
		for k := j; k > 0 && sort.More(i.Array, k, k-1); k-- {
			sort.Exchange(i.Array, k, k-1)
		}
	}
}

func (i insertion) SortAsc() {
	l := len(i.Array)
	for j := 0; j < l; j++ {
		for k := j; k > 0 && sort.Less(i.Array, k, k-1); k-- {
			sort.Exchange(i.Array, k, k-1)
		}
	}
}

func (i insertion) GetArr() []int {
	return i.Array
}
