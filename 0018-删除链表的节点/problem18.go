package _18_删除链表的节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	dummy := ListNode{Val: 1<<31 - 1, Next: head}
	pre, ptr := head, head.Next

	for ptr != nil {
		if ptr.Val == val {
			temp := ptr
			pre.Next = ptr.Next
			ptr = temp.Next
		} else {

			pre = ptr
			ptr = ptr.Next
		}
	}

	return dummy.Next
}
