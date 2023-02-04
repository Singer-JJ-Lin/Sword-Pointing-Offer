package _04_二维数组中的查找

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix[0]) == 0 {
		return false
	}

	length := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		l, r := 0, length-1
		for l < r {
			mid := (l + r) >> 1
			if matrix[i][mid] >= target {
				r = mid
			} else {
				l = mid + 1
			}
		}

		if matrix[i][l] == target {
			return true
		}
	}

	return false
}
