package _33_二叉搜索树的后序遍历

func verifyPostorder(postorder []int) bool {
	return dfs(postorder, 0, len(postorder)-1)
}

func dfs(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}

	// 寻找左右子树分界点
	p := i
	for ; postorder[j] > postorder[p]; p++ {
	}

	// 寻找第一个不满足大于根节点的数
	m := p
	for ; postorder[p] > postorder[j]; p++ {
	}

	// 递归判断
	return p == j && dfs(postorder, i, m-1) && dfs(postorder, m, j-1)
}
