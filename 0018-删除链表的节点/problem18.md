# 删除链表的节点
- 简单题目，设置pre和ptr两个变量，pre指向ptr前一个元素
- 不断遍历，直到ptr的值与要删除的节点相同
- pre->next = ptr.next;
- 结束
