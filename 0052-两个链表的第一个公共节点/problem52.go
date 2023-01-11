package _52_两个链表的第一个公共节点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 比较长度的方法
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	len_A, len_B := 0, 0
	ptr_A, ptr_B := headA, headB

	//统计链表A和链表B的长度
	for ptr_A != nil {
		len_A++
		ptr_A = ptr_A.Next
	}

	for ptr_B != nil {
		len_B++
		ptr_B = ptr_B.Next
	}

	// 重新设置节点
	ptr_A, ptr_B = headA, headB

	//如果链表A长，那么链表A先走
	if len_A > len_B {
		difference := len_A - len_B
		for difference > 0 {
			ptr_A = ptr_A.Next
			difference = difference - 1
		}
		//如果链表B长，那么链表B先走
	} else if len_B > len_A {
		difference := len_B - len_A
		for difference > 0 {
			ptr_B = ptr_B.Next
			difference = difference - 1
		}
	}

	//然后开始比较，如果他俩不相等就一直往下走
	for ptr_A != nil && ptr_B != nil && ptr_A != ptr_B {
		ptr_A = ptr_A.Next
		ptr_B = ptr_B.Next
	}

	// 跳出循环后无非两种结果
	// 1.相交
	// 2.没有相交的节点，此时ptr_A和ptr_B均为nil
	return ptr_A

}

// 双指针的方法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 设置指针
	ptr_A, ptr_B := headA, headB
	for ptr_A != ptr_B {
		if ptr_A != nil {
			ptr_A = ptr_A.Next
		} else {
			ptr_A = headB
		}

		if ptr_B != nil {
			ptr_B = ptr_B.Next
		} else {
			ptr_B = headA
		}
	}

	return ptr_A
}
