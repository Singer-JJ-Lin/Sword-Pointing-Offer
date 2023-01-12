package _21_调整数组顺序使奇数位于偶数前面

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	i, j := -1, len(nums)

	// do-while循环，先i++，再进行判断
	for i < j {
		for {
			i++
			if i >= j || nums[i]&1 == 0 {
				break
			}
		}

		// do-while循环，先j++，再进行判断
		for {
			j--
			if i >= j || nums[j]&1 == 1 {
				break
			}
		}

		// 交换两个数
		if i < j {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
		}
	}

	return nums
}
