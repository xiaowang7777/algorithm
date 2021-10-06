package priority_queue

import (
	"fmt"
	"testing"
)

func TestPriorityMaxQueue_DelAndGet(t *testing.T) {
	arr := []int{5, 1, 15, 3, 4, 3, 6, 0, 2, 1}
	i := NewMaxQueueWithArr(arr)
	for !i.IsEmpty() {
		fmt.Printf("%d\t", i.DelAndGet())
	}
}

func TestPriorityMinQueue_DelAndGet(t *testing.T) {
	arr := []int{5, 1, 15, 3, 4, 3, 6, 0, 2, 1, 0}
	i := NewMinQueueWithArr(arr)
	for !i.IsEmpty() {
		fmt.Printf("%d\t", i.DelAndGet())
	}
}
