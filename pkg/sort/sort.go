package sort

type Array []int

type Sort interface {
	SortAsc()
	SortDesc()
	GetArr() []int
}

func IsAscSorted(arr Array) bool {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		if !(Less(arr, i, i+1) || Eq(arr, i, i+1)) {
			return false
		}
	}
	return true
}

func IsDescSorted(arr Array) bool {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		if !(More(arr, i, i+1) || Eq(arr, i, i+1)) {
			return false
		}
	}
	return true
}

func More(arr Array, first, last int) bool {
	return arr[first] > arr[last]
}

func Eq(arr Array, first, last int) bool {
	return arr[first] == arr[last]
}

func Less(arr Array, first, last int) bool {
	return arr[first] < arr[last]
}

func Exchange(arr Array, first, last int) {
	t := arr[first]
	arr[first] = arr[last]
	arr[last] = t
}
