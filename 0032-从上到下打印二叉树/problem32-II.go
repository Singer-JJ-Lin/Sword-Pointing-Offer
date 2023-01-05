package _032_从上到下打印二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder2(root *TreeNode) [][]int {

	if root == nil { // 如果根节点为空，则直接返回nil
		return nil
	}

	// 如果二叉树只有根节点，则把根节点的Val属性放入到数组中返回
	// 返回[[root.Val]]
	if root.Left == nil && root.Right == nil {
		return [][]int{
			{
				root.Val,
			},
		}
	}

	// 定义队列并把根节点加入到队列中去
	queue := make([]*TreeNode, 1010)
	hh, tt := 0, -1
	tt++
	queue[tt] = root

	// 设置start和end两个指针, end指针表示当前层的最后一个节点
	start, end := root, root

	// 用于存储每一层的节点
	layer_result := make([]int, 0)

	// 返回的结果
	result := make([][]int, 0)

	// 首先把根节点加入到result中
	result = append(result, []int{root.Val})

	// while循环，队列不为空
	for tt > hh {

		// 出队
		start = queue[hh]
		hh++

		// 判断左子树
		if start.Left != nil {
			tt++
			queue[tt] = start.Left
			layer_result = append(layer_result, start.Left.Val)
		}

		// 判断右子数
		if start.Right != nil {
			tt++
			queue[tt] = start.Right
			layer_result = append(layer_result, start.Right.Val)
		}

		// 刚好遍历一层，就把layer_result加入到result中
		// 并设置end为下一层最后一个节点，layer_result清空
		if start == end && len(layer_result) > 0 {
			result = append(result, layer_result)
			layer_result = nil
			end = queue[tt]
		}
	}

	// 返回结果
	return result
}
