package shell

import "algorithm/pkg/sort"

type shell struct {
	sort.Array
}

func New(arr []int) *shell {
	return &shell{Array: arr}
}

func (s shell) SortAsc() {
	l := len(s.Array)
	h := 1
	for h < l/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h && sort.Less(s.Array, j, j-h); j -= h {
				sort.Exchange(s.Array, j, j-h)
			}
		}
		h /= 3
	}
}

func (s shell) SortDesc() {
	l := len(s.Array)
	h := 1
	for h < l/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h && !sort.Less(s.Array, j, j-h); j -= h {
				sort.Exchange(s.Array, j, j-h)
			}
		}
		h /= 3
	}
}

func (s shell) GetArr() []int {
	return s.Array
}
