package _028_对称的二叉树

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

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p *TreeNode, q *TreeNode) bool {

	// p和q均为空，也是对称
	if p == nil && q == nil {
		return true
	}

	// p和q一个为空，一个不为空，说明不对称
	if p == nil || q == nil {
		return false
	}

	// 递归比较
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
