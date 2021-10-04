package merge

import "algorithm/pkg/sort"

var tmp []int

func mergeAsc(arr []int, lo, mid, hi int) {
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		tmp[k] = arr[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = tmp[j]
			j++
		} else if j > hi {
			arr[k] = tmp[i]
			i++
		} else if sort.Less(tmp, i, j) {
			arr[k] = tmp[i]
			i++
		} else {
			arr[k] = tmp[j]
			j++
		}
	}
}

func mergeDesc(arr []int, lo, mid, hi int) {
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		tmp[k] = arr[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = tmp[j]
			j++
		} else if j > hi {
			arr[k] = tmp[i]
			i++
		} else if sort.More(tmp, i, j) {
			arr[k] = tmp[i]
			i++
		} else {
			arr[k] = tmp[j]
			j++
		}
	}
}

type merge struct {
	sort.Array
}

func New(arr []int) *merge {
	return &merge{arr}
}

func (m merge) SortDesc() {
	tmp = make([]int, len(m.Array))
	sortDesc(m.Array, 0, len(m.Array)-1)
}

func (m merge) SortAsc() {
	tmp = make([]int, len(m.Array))
	sortAsc(m.Array, 0, len(m.Array)-1)
}

func (m merge) GetArr() []int {
	return m.Array
}

func sortDesc(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	sortDesc(arr, lo, mid)
	sortDesc(arr, mid+1, hi)
	mergeDesc(arr, lo, mid, hi)
}

func sortAsc(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	sortAsc(arr, lo, mid)
	sortAsc(arr, mid+1, hi)
	mergeAsc(arr, lo, mid, hi)
}
