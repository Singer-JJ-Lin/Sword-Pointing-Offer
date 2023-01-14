package _54_二叉搜索树的第k大节点

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var res int
	dfs(root, &k, &res)
	return res
}

func dfs(root *TreeNode, k *int, res *int) {
	// 特殊判断
	if root == nil {
		return
	}

	// 递归右子树
	dfs(root.Right, k, res)

	// 如果k为0，则进行减枝
	if *k == 0 {
		return
	}

	// 如果当前节点为第K个节点
	*k--
	if *k == 0 {
		*res = root.Val
		return
	}

	// 递归遍历左子树
	dfs(root.Left, k, res)
}
