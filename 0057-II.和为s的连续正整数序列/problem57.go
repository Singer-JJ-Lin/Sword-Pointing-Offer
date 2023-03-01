package _57_II_和为s的连续正整数序列

func findContinuousSequence(target int) [][]int {
	s := make([]int, target+1)
	for i := 0; i <= target; i++ {
		s[i] = s[i-1] + i
	}

	var res [][]int

	for i, j := 1, target; i <= target; i++ {
		for s[j]-s[i-1] > target && j > i {
			j--
		}

		if s[j]-s[i-1] == target {
			temp := []int{}
			for k := i; k <= j; k++ {
				temp = append(temp, k)
			}

			res = append(res, temp)
		}
	}

	return res

}
