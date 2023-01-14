package _36_二叉搜索树与双向链表

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var (
	head, pre *Node
)

func dfs(root *Node) {
	// 特殊判断
	if root == nil {
		return
	}

	// 递归遍历左子树
	dfs(root.Left)

	// 构建前驱后继关系
	if pre != nil {
		pre.Right = root
	} else {
		head = root
	}
	root.Left = pre
	pre = root

	// 递归遍历右子树
	dfs(root.Right)
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	dfs(root)

	// 设置首尾节点
	head.Left = pre
	pre.Right = head
	return head
}