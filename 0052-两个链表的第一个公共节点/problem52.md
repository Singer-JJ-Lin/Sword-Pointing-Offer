# 两个链表的第一个公共节点
## 比较链表长度的方法
1. 统计两个链表A，B的长度len_A, len_B   
2. 如果len_A > len_B, 则A先走`len_A - len_B`步  
如果len_B > len_A, 则B先走`len_B - len_A`步
3. 然后逐个节点比较

## 双指针的方法

我们设置两个指针ptr_A指向headA, ptr_B指向headB, 然后ptr_A和ptr_B同时走。当ptr_A到达链表尾时，ptr_A指向headB，当ptr_B到达链表尾时，ptr_B指向headA。
当两个指针都经历过重新指向headA和headB的过程后，两个指针距离公共节点的距离是相同的，只需不断的往后走到第一个相同的节点即可