package 堆栈

/*
剑指 Offer 09. 用两个栈实现队列
简单
763
相关企业
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
[[],[3],[],[],[]]
输出：[null,null,3,-1,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/

type CQueue struct {
	stack []int
}

func Constructor1() CQueue {
	return CQueue{stack: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.stack = append(this.stack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack) == 0 {
		return -1
	}

	var temp []int
	var newStack []int
	headNum := this.stack[0]

	for i := len(this.stack) - 1; i >= 1; i-- {
		temp = append(temp, this.stack[i])
	}

	for i := len(temp) - 1; i >= 0; i-- {
		newStack = append(newStack, temp[i])
	}

	this.stack = newStack

	return headNum
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
