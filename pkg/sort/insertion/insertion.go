package insertion

import "algorithm/pkg/sort"

type insertion struct {
	sort.Array
}

func (i insertion) SortDesc() {
	l := len(i.Array)
	for j := 0; j < l; j++ {
		for k := j; k >= 0 && sort.Less(i.Array, j, k); k-- {
			sort.Exchange(i.Array, j, k)
		}
	}
}

func (i insertion) SortAsc() {
	l := len(i.Array)
	for j := 0; j < l; j++ {
		for k := j; k >= 0 && sort.More(i.Array, j, k); k-- {
			sort.Exchange(i.Array, j, k)
		}
	}
}

func (i insertion) GetArr() []int {
	return i.Array
}
