package _57_和为s的两个数字

func twoSum(nums []int, target int) []int {
	length := len(nums)

	// 特殊判断
	if length == 0 {
		return nil
	}

	i, j := 0, length-1

	for i < length {
		for j > 0 && nums[i]+nums[j] > target {
			j--
		}

		if j > 0 && nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		}

		i++
	}

	return nil
}
