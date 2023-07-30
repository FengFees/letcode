package 栈

/*
剑指 Offer 59 - II. 队列的最大值
中等
498
相关企业
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5
*/

type MaxQueue struct {
	stack    []int
	maxStack []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.maxStack[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.stack = append(this.stack, value)
	for len(this.maxStack) > 0 && this.maxStack[len(this.maxStack)-1] < value {
		this.maxStack = this.maxStack[:len(this.maxStack)-1]
	}
	this.maxStack = append(this.maxStack, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.stack) == 0 {
		return -1
	}
	head := this.stack[0]

	this.stack = this.stack[1:len(this.stack)]

	if head == this.maxStack[0] {
		this.maxStack = this.maxStack[1:len(this.maxStack)]
	}

	return head
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
