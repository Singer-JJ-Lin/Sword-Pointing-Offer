# 从上到下打印二叉树II
这道题在上一道题的基础上加多了一个tricky part, 上一题是把所有元素放到一个数组中返回，这一题不同之处在于返回的是一个二维数组，是按层分好的。
  
我们需要开一个辅助数组存储每一层的节点，并设置两个指针start和end， end表示当前层最后的节点。当`start==end`时，说明我们已经把一层的结果装入到数组中了，把这个数组加入到结果数组中，并把数组置nil,继续遍历，直到队列为空。



## queue模板
我们用一个数组来模拟队列，hh, tt 分别为队头指针，对尾指针，初值分别为0和1
```go
// 创建队列，假设长度为1000
var queue = make([]int, 1000)
var hh = 0
var tt = -1

// 入队
func push(x int) {
	tt++
	queue[tt] = x
}

// 出队
func pop() int {
	hh++
	return q[hh-1]
}

// 判断队列是否为空
func empty() bool {
	return tt > hh
}
```


## 具体思路看代码文件注释