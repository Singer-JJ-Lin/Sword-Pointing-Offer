# 调整数组顺序使奇数位于偶数前面
使用快速排序的思想来做这道题。
1. 设置两个指针l,r分别指向第一个元素和最后一个元素
2. l不断往后走，找到第一个偶数， r不断往前走，找到第一个奇数，然后交换这两个数
3. 重复第2步直到`l >= r`
