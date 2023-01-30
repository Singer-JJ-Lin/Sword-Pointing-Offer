package _42_连续子数组的最大和

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += max(nums[i-1], 0)
		result = max(result, nums[i])
	}

	return result
}
