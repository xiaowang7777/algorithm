package selection

import "algorithm/pkg/sort"

type selection struct {
	sort.Array
}

func New(arr []int) *selection {
	return &selection{arr}
}

func (s selection) SortDesc() {
	l := len(s.Array)

	for i := 0; i < l; i++ {
		max := i
		for j := i + 1; j < l; j++ {
			if sort.More(s.Array, j, max) {
				max = j
			}
		}
		sort.Exchange(s.Array, i, max)
	}
}

func (s selection) SortAsc() {
	l := len(s.Array)

	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if sort.Less(s.Array, j, min) {
				min = j
			}
		}
		sort.Exchange(s.Array, i, min)
	}
}

func (s selection) GetArr() []int {
	return s.Array
}
