package heap

import "algorithm/pkg/sort"

type heap struct {
	sort.Array
}

func New(arr []int) *heap {
	return &heap{arr}
}

func (h heap) SortDesc() {
	l := len(h.Array) - 1
	for k := l / 2; k > 0; k-- {
		h.sinkDesc(k, l)
	}
	for l > 1 {
		sort.Exchange(h.Array, 1, l)
		l--
		h.sinkDesc(1, l)
	}
}

func (h *heap) SortAsc() {
	l := len(h.Array) - 1
	//对每个子堆进行下沉操作，最后得到的堆是一个优先队列的堆
	for k := l / 2; k > 0; k-- {
		h.sinkAsc(k, l)
	}
	//然后依次删除头节点（将头节点和最后一个节点进行交换，然后堆大小减一，并对对进行下沉操作）
	for l > 1 {
		sort.Exchange(h.Array, 1, l)
		l--
		h.sinkAsc(1, l)
	}
}

func (h heap) GetArr() []int {
	return h.Array
}

//sinkAsc 下沉
func (h heap) sinkAsc(i, n int) {
	for 2*i <= n {
		j := 2 * i
		//判断节点的值是否小于左子节点的值
		if sort.Less(h.Array, i, j) {
			//当节点的值小于左子节点的值，
			//并且右子节点的值小于左子节点的值时，
			//交换节点和左子节点的值，
			//否则交换节点与右子节点的值
			if j+1 <= n && !sort.Less(h.Array, j+1, j) {
				sort.Exchange(h.Array, i, j+1)
				i = j + 1
			} else {
				sort.Exchange(h.Array, i, j)
				i = j
			}
		} else
		//当节点的值大于左节点的值时，
		//判断是否有右子节点的值，
		//以及节点的值是否小于右子节点的值，
		//如果小于右子节点的值则交换
		if j+1 <= n && sort.Less(h.Array, i, j+1) {
			sort.Exchange(h.Array, i, j+1)
			i = j + 1
		} else {
			break
		}
	}
}

func (h heap) sinkDesc(i, n int) {
	for i*2 <= n {
		j := i * 2
		if sort.More(h.Array, i, j) {
			if j+1 <= n && sort.More(h.Array, j, j+1) {
				sort.Exchange(h.Array, j+1, i)
				i = j + 1
			} else {
				sort.Exchange(h.Array, j, i)
				i = j
			}
		} else if j+1 <= n && sort.More(h.Array, i, j+1) {
			sort.Exchange(h.Array, j+1, i)
			i = j + 1
		} else {
			break
		}
	}
}
