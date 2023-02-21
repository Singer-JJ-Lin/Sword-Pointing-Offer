package _56_II_数组中数字出现的次数

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}
