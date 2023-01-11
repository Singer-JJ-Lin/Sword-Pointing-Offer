package _25_合并两个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// 特殊情况判断，两个链表都为空，直接返回nil
	if l1 == nil && l2 == nil {
		return nil
	}

	// l1为空， l2不为空的情况
	if l1 != nil && l2 == nil {
		return l1
	}

	// l2为空， l1不为空的情况
	if l2 != nil && l1 == nil {
		return l2;
	}

	// 设置哑节点，最后直接返回dummy.Next
	dummy := new(ListNode);
	ptr := dummy

	// 归并排序
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			ptr.Next = l2
			l2 = l2.Next
		} else {
			ptr.Next = l1
			l1 = l1.Next
		}

		ptr = ptr.Next
	}

	// l1还有剩余元素，不断加入到链表中
	for l1 != nil {
		ptr.Next = l1
		ptr = ptr.Next
		l1 = l1.Next
	}
	// l2还有剩余元素，不断加入到链表中
	for l2 != nil {
		ptr.Next = l2
		ptr = ptr.Next
		l2 = l2.Next
	}

	// 返回链表头节点，也就是dummy的下一个节点
	return dummy.Next;
}