package 堆栈

import "math"

/*
剑指 Offer 30. 包含min函数的栈
简单
514
相关企业
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。



示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.


提示：

各函数的调用总次数不超过 20000 次
*/

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{
		stack: []int{},
		min:   math.MaxInt32,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x-this.min)

	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	if this.stack[len(this.stack)-1] <= 0 {
		this.min = this.min - this.stack[len(this.stack)-1]
	}

	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if this.stack[len(this.stack)-1] < 0 {
		return this.min
	}

	return this.min + this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
