package _03_数组中重复的数字

func findRepeatNumber(nums []int) int {
	dict := make(map[int]bool)
	for _, num := range nums {
		if _, ok := dict[num]; ok {
			return num
		} else {
			dict[num] = true
		}
	}
	return -1
}
