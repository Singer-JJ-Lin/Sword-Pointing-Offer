package _11_旋转数组的最小数字

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		mid := (l + r) >> 1
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}

	return numbers[l]
}
