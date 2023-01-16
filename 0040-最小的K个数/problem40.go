package _40_最小的K个数

func down(heap []int, u, size int) {
	t := u
	if u*2+1 <= size && heap[t] < heap[2*u+1] {
		t = 2*u + 1
	}

	if u*2+2 <= size && heap[t] < heap[2*u+2] {
		t = 2*u + 2
	}

	if t != u {
		heap[t], heap[u] = heap[u], heap[t]
		down(heap, t, size)
	}
}

func getLeastNumber(arr []int, k int) []int {
	size := 4
	res := make([]int, k)
	for i := 1; i >= 0; i-- {
		down(arr, i, size)
	}

	for i := k; k < len(arr); i++ {
		if arr[0] > arr[i] {
			arr[0] = arr[4]
			arr[4] = arr[i]
			down(arr, 0, 4)
		}
	}

	for i := 0; i < k; i++ {
		res[i] = arr[i]
	}

	return res
}
