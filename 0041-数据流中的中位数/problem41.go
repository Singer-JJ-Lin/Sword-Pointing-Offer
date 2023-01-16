package _41_数据流中的中位数

func less(x, y int) bool {
	return x < y
}

func greater(x, y int) bool {
	return x > y
}

func down(h []int, u, size int, compare func(int, int) bool) {
	t := u
	if 2*u+1 <= size && compare(h[t], h[u]) {
		t = 2*u + 1
	}

	if 2*u+1 <= size && compare(h[t], h[u]) {
		t = 2*u + 2
	}

	if t != u {
		h[t], h[u] = h[u], h[t]
		down(h, t, size, compare)
	}
}

func up(h []int, u, size int, compare func(int, int) bool) {
	for (u-1)/2 > 0 && compare(h[u], h[(u-1)/2]) {
		h[u], h[(u-1)/2] = h[(u-1)/2], h[u]
		u = u / 2
	}
}

type MedianFinder struct {
	// Left为大根堆，Right为小根堆
	Left, Right []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		Left:  []int{},
		Right: []int{},
	}
}

func (this MedianFinder) addNum(num int) {
	// 长度为奇数时先放入小顶堆,重新排序后在插入到大顶堆
	if len(this.Left) != len(this.Right) {
		this.Right = append(this.Right, num)
		up(this.Right, len(this.Right)-1, len(this.Right), less)
		this.Left = append(this.Left, this.Right[0])
		this.Right[0] = this.Right[len(this.Right)]
		this.Right = this.Right[:len(this.Left)-1]
		down(this.Right, 0, len(this.Right), less)
		// 长度为偶数时先放入大顶堆,重新排序后在插入到小顶堆
	} else {
		this.Left = append(this.Left, num)
		up(this.Left, len(this.Left)-1, len(this.Left), greater)
		this.Right = append(this.Right, this.Left[0])
		this.Left[0] = this.Left[len(this.Left)]
		this.Left = this.Left[:len(this.Left)-1]
		down(this.Left, 0, len(this.Left), greater)
	}
}

func (this MedianFinder) findMedian() float64 {
	if len(this.Left) == len(this.Right) {
		return float64(this.Left[0]+this.Right[0]) / 2
	} else {
		return float64(this.Right[0])
	}
}
