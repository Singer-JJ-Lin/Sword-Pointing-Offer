package _30_包含min函数的栈

import "math"

type MinStack struct {
	stack, min_stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:     []int{},
		min_stack: []int{math.MaxInt64},
	}
}

func min(x int, y int) int {
	if x > y {
		return y
	}

	return x
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.min_stack[len(this.min_stack)-1]
	this.min_stack = append(this.min_stack, min(x, top))

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min_stack = this.min_stack[:len(this.min_stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min_stack[len(this.min_stack)-1]
}
