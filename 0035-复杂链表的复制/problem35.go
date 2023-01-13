package _35_复杂链表的复制

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// 特殊判断
	if head == nil {
		return nil
	}

	// 建立哑节点
	dummy := new(Node)

	// 设立ptr指向旧链表，new_ptr指向新链表
	ptr, new_ptr := head, dummy
	mapper := make(map[*Node]*Node)

	// 构建新链表同时建立映射关系
	for ptr != nil {
		node := new(Node)
		node.Val = ptr.Val
		new_ptr.Next = node
		mapper[ptr] = node
		ptr, new_ptr = ptr.Next, new_ptr.Next
	}

	// reset指针
	ptr, new_ptr = head, dummy.Next

	// 构建新链表的random指针
	for ptr != nil {
		new_ptr.Random = mapper[ptr.Random]
		ptr, new_ptr = ptr.Next, new_ptr.Next
	}

	return dummy.Next
}
