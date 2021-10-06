package quick

import "algorithm/pkg/sort"

type quick struct {
	sort.Array
}

func New(arr []int) *quick {
	return &quick{arr}
}

func (q quick) SortAsc() {
	sortAsc(q.Array, 0, len(q.Array)-1)
}

func (q quick) SortDesc() {
	sortDesc(q.Array, 0, len(q.Array)-1)
}

func (q quick) GetArr() []int {
	return q.Array
}

func sortAsc(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partitionAsc(arr, lo, hi)
	sortAsc(arr, lo, j-1)
	sortAsc(arr, j+1, hi)
}

func sortDesc(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partitionDesc(arr, lo, hi)
	sortDesc(arr, lo, j-1)
	sortDesc(arr, j+1, hi)
}

func partitionAsc(arr []int, lo, hi int) int {
	i, j := lo, hi+1
	//取指定的第一个元素作为切分标准
	v := arr[lo]
	for true {

		for i++; v >= arr[i]; i++ {
			if i == hi {
				break
			}
		}

		for j--; v <= arr[j]; j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		sort.Exchange(arr, i, j)
	}
	sort.Exchange(arr, lo, j)
	return j
}

func partitionDesc(arr []int, lo, hi int) int {
	i, j := lo, hi+1
	v := arr[lo]

	for true {
		for i++; v <= arr[i]; i++ {
			if i == hi {
				break
			}
		}

		for j--; v >= arr[j]; j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		sort.Exchange(arr, i, j)
	}
	sort.Exchange(arr, lo, j)
	return j
}
