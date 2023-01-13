package _35_复杂链表的复制

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	dummy := new(Node)
	ptr, new_ptr := head, dummy
	mapper := make(map[*Node]*Node)

	for ptr != nil {
		node := new(Node)
		node.Val = ptr.Val
		new_ptr.Next = node
		mapper[ptr] = node
		ptr, new_ptr = ptr.Next, new_ptr.Next
	}

	ptr, new_ptr = head, dummy.Next
	for ptr != nil {
		new_ptr.Random = mapper[ptr.Random]
		ptr, new_ptr = ptr.Next, new_ptr.Next
	}

	return dummy.Next
}
