package _064_求1_2_____n

func sumNums(n int) int {
	// 递归出口
	if n == 1 {
		return 1
	}

	return n + sumNums(n-1)
}
