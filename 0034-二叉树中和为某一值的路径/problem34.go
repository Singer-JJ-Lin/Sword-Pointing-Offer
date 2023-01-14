package _34_二叉树中和为莫一值的路径

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	// 特殊判断
	if root == nil {
		return nil
	}
	var res [][]int
	dfs(root, target, []int{}, &res)
	return res
}
func dfs(root *TreeNode, target int, path []int, res *[][]int) {
	// 特殊判断
	if root == nil {
		return
	}

	// 添加当前节点到当前路径
	path = append(path, root.Val)

	// 判断是否为符合条件的路径
	if root.Val == target && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	// 递归遍历左右子树
	dfs(root.Left, target-root.Val, path, res)
	dfs(root.Right, target-root.Val, path, res)

	// 状态恢复
	path = path[:len(path)-1]
}
