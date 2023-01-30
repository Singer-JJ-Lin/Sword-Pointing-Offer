package _47_礼物的最大价值

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func maxValue(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[row-1][col-1]
}
