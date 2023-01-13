package _06_从尾到头打印字符串

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return append(reversePrint(head.Next), head.Val)
}
