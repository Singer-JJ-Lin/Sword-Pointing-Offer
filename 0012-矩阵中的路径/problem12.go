package _12_矩阵中的路径

var directions [][]int = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
var visited [][]bool

func dfs(board [][]byte, word string, i, x, y int) bool {
	if i == len(word) {
		return true
	}

	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] || word[i] != board[x][y] {
		return false
	}

	visited[x][y] = true

	for _, direction := range directions {
		new_x, new_y := x+direction[0], y+direction[1]
		if dfs(board, word, i+1, new_x, new_y) {
			return true
		}
	}

	visited[x][y] = false
	return false
}

func exist(board [][]byte, word string) bool {
	visited = make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := range board {
		for j := range board[i] {
			if dfs(board, word, 0, i, j) {
				return true
			}
		}
	}

	return false
}
