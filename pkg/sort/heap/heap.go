package heap

import "algorithm/pkg/sort"

type heap struct {
	arr sort.Array
	n   int
}

func New(arr []int) *heap {
	return &heap{arr: arr, n: len(arr) + 1}
}

func (h heap) SortDesc() {

}

func (h heap) SortAsc() {
	l := len(h.arr) - 1
	for k := l / 2; k > 0; k-- {
		h.sinkAsc(k, l)
	}
	for l > 1 {
		sort.Exchange(h.arr, 1, l)
		l--
		h.sinkAsc(1, l)
	}
}

func (h heap) GetArr() []int {
	return h.arr
}

//下沉
func (h heap) sinkAsc(i, n int) {
	for 2*i <= n {
		j := 2 * i
		if sort.Less(h.arr, i, j) {
			if j+1 <= n && !sort.Less(h.arr, j+1, j) {
				sort.Exchange(h.arr, i, j+1)
				i = j + 1
			} else {
				sort.Exchange(h.arr, i, j)
				i = j
			}
		} else if j+1 <= n && sort.Less(h.arr, i, j+1) {
			sort.Exchange(h.arr, i, j+1)
			i = j + 1
		} else {
			break
		}
	}
}
