package _68_I_二叉搜索树的最近公共祖先二叉搜索树的最近公共祖先

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}
