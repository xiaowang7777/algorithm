package priority_queue

import "algorithm/pkg/sort"

type priorityMaxQueue struct {
	//存储数据的数组
	arr []int
	//存储数据的长度
	n int
}

func NewMaxQueue(len int) *priorityMaxQueue {
	return &priorityMaxQueue{
		arr: make([]int, len+1),
	}
}

func NewMaxQueueWithArr(arr []int) *priorityMaxQueue {
	dstArr := make([]int, len(arr)+1)
	p := &priorityMaxQueue{arr: dstArr}
	for i := 0; i < len(arr); i++ {
		p.Insert(arr[i])
	}
	return p
}

func (p priorityMaxQueue) IsEmpty() bool {
	return p.n == 0
}

func (p *priorityMaxQueue) Insert(k int) {
	p.arr[p.n+1] = k
	p.n++
	p.swim(p.n)
}

func (p *priorityMaxQueue) DelAndGet() int {
	i := p.arr[1]

	i2 := p.arr[p.n]
	p.arr[1] = i2
	p.arr[p.n] = 0
	p.sink(1)
	p.n--
	return i
}

//上浮
func (p priorityMaxQueue) swim(i int) {
	for i > 1 && sort.Less(p.arr, i/2, i) {
		sort.Exchange(p.arr, i/2, i)
		i /= 2
	}
}

//下沉
func (p priorityMaxQueue) sink(i int) {

	for 2*i < p.n {
		j := 2 * i
		if sort.Less(p.arr, i, j) {
			if sort.Less(p.arr, j, j+1) && j+1 <= p.n {
				sort.Exchange(p.arr, j+1, i)
				i = j + 1
			} else {
				sort.Exchange(p.arr, j, i)
				i = j
			}
		} else if j+1 <= p.n && sort.Less(p.arr, i, j+1) {
			sort.Exchange(p.arr, j+1, i)
			i = j + 1
		} else {
			break
		}
	}
}

type priorityMinQueue struct {
	//存储数据的数组
	arr []int
	//存储数据的长度
	n int
}

func NewMinQueue(len int) *priorityMinQueue {
	return &priorityMinQueue{arr: make([]int, len+1)}
}

func NewMinQueueWithArr(arr []int) *priorityMinQueue {
	dstArr := make([]int, len(arr)+1)
	p := &priorityMinQueue{arr: dstArr}
	for i := 0; i < len(arr); i++ {
		p.Insert(arr[i])
	}
	return p
}

func (p *priorityMinQueue) Insert(k int) {
	p.arr[p.n+1] = k
	p.n++
	p.swim(p.n)
}

func (p priorityMinQueue) IsEmpty() bool {
	return p.n == 0
}

func (p *priorityMinQueue) DelAndGet() int {
	if p.IsEmpty() {
		return 0
	}

	i := p.arr[1]

	i2 := p.arr[p.n]
	p.arr[1] = i2
	p.arr[p.n] = 0
	p.n--
	p.sink(1)

	return i
}

//上浮
func (p priorityMinQueue) swim(i int) {
	for sort.More(p.arr, i/2, i) {
		sort.Exchange(p.arr, i/2, i)
		i /= 2
	}
}

//下沉
func (p priorityMinQueue) sink(i int) {
	for 2*i <= p.n {
		j := 2 * i
		if sort.More(p.arr, i, j) {
			if sort.Less(p.arr, j+1, j) && j+1 > p.n {
				sort.Exchange(p.arr, i, j)
				i = j
			} else if sort.Less(p.arr, j, j+1) {
				sort.Exchange(p.arr, i, j)
				i = j
			} else {
				sort.Exchange(p.arr, i, j+1)
				i = j + 1
			}
		} else if sort.More(p.arr, i, j+1) && j+1 <= p.n {
			sort.Exchange(p.arr, i, j+1)
			i = j + 1
		} else {
			break
		}
	}
}
