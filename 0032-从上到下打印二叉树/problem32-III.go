package _032_从上到下打印二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return [][]int{
			{
				root.Val,
			},
		}
	}

	queue := make([]*TreeNode, 1010)
	start, end := root, root
	hh, tt := 0, -1
	tt++
	queue[tt] = root
	layer_result := make([]int, 0)
	result := make([][]int, 0)
	result = append(result, []int{root.Val})

	// 用于表示需不需要将当前层的节点逆转
	flag := true

	for tt > hh {
		hh++
		start = queue[hh]

		if start.Left != nil {
			tt++
			queue[tt] = start.Left
			layer_result = append(layer_result, start.Left.Val)
		}

		if start.Right != nil {
			tt++
			queue[tt] = start.Right
			layer_result = append(layer_result, start.Right.Val)
		}

		if start == end && len(layer_result) > 0 {

			// 如果需要逆转，就逆转
			if flag == true {
				i, j := 0, len(layer_result)-1
				for i < j {
					temp := layer_result[i]
					layer_result[i] = layer_result[j]
					layer_result[j] = temp
					i++
					j = j - 1
				}
			}
			result = append(result, layer_result)
			layer_result = nil
			end = queue[tt]

			if flag == true {
				flag = false
			} else {
				flag = true
			}
		}
	}

	return result
}
