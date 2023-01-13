package _13_机器人的运动范围

var visited [][]bool
var directions [][]int = [][]int{
	{-1, 0}, {1, 0}, {0, 1}, {0, -1},
}

func movingCount(m int, n int, k int) int {
	visited = make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return dfs(0, 0, m, n, k)
}

func digitSum(x int) int {
	result := 0
	for x > 0 {
		result += x % 10
		x /= 10
	}
	return result
}

func dfs(x, y, m, n, k int) int {
	if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] || digitSum(x)+digitSum(y) > k {
		return 0
	}

	visited[x][y] = true
	result := 1
	for _, direction := range directions {
		result += dfs(x+direction[0], y+direction[1], m, n, k)
	}

	return result
}
