package _026_树的子结构

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 这个函数用于判断从A,B这两个根节点开始比较，A是否含有B
func isContain(A *TreeNode, B *TreeNode) bool {
	// 如果B等于空，也可以认为A含有B
	if B == nil {
		return true
	}

	// 如果A为空，或者A.Val与B.Val不相等，则说明不满足条件
	if A == nil || A.Val != B.Val {
		return false
	}

	// 满足条件则继续递归判断左子树和右子树
	return isContain(A.Left, B.Left) && isContain(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// A和B只要有一个为空，B就不是A的子结构
	if A == nil || B == nil {
		return false
	}

	// 从A和B两个根结点开始比较， 根节点不同的话递归考虑左右子树
	return isContain(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
