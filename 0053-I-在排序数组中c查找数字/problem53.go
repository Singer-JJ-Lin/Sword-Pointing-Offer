package _53_I_在排序数组中查找数字

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if nums[l] != target {
		return 0
	}

	old_l := l
	l = 0
	r = len(nums) - 1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l - old_l + 1

}
