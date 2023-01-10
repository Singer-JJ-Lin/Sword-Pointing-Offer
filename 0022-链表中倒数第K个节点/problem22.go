package _22_链表中倒数第K个节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for k > 0 && fast != nil {
		fast = fast.Next
		k = k - 1
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
