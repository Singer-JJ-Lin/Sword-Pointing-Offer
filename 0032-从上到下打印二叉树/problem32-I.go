package _032_从上到下打印二叉树

func levelOrder(root *TreeNode) []int {
	if root == nil { // 如果根节点为空，则直接返回nil
		return nil
	}

	// 如果二叉树只有根节点，则把根节点的Val属性放入到数组中返回
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	// result: 用于存储结果
	result := make([]int, 0)

	// 定义队列并把根节点加入到队列中去
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	// while循环
	for len(queue) > 0 {
		// 从队头取出第一个元素并出队
		node := queue[0]
		queue = queue[1:]

		// 将当前节点的值加入到result中
		result = append(result, node.Val)

		// 如果左孩子不为空，则把左孩子加入队列
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// 如果右孩子不为空，则把右孩子加入队列
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	// 返回结果
	return result
}
