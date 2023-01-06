package _027_二叉树的镜像

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	mirrorTree(root.Left)
	mirrorTree(root.Right)
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	return root
}

