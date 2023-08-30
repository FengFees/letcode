package main

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]
示例2：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]
示例3：

输入：l1 = [0], l2 = [0]
输出：[0]


提示：

链表的长度范围为 [1, 100]
0 <= node.val <= 9
输入数据保证链表代表的数字无前导 0


进阶：如果输入链表不能修改该如何处理？换句话说，不能对列表中的节点进行翻转。

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 迭代循环
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var d1, d2 int
	a, b := l1, l2
	for a != nil {
		d1++
		a = a.Next
	}
	for b != nil {
		d2++
		b = b.Next
	}
	res, c := addTwoNumbers2(l1, l2, d1, d2)
	if c > 0 {
		return &ListNode{
			Val:  c,
			Next: res,
		}
	}

	return res
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode, d1 int, d2 int) (res *ListNode, c int) {
	if l1 != nil && l2 != nil {
		if l1.Next == nil && l2.Next == nil {
			n := l1.Val + l2.Val
			c = n / 10
			res = &ListNode{
				Val:  n % 10,
				Next: nil,
			}
		}
	}

	var a *ListNode
	var b, n int
	if d1 > d2 {
		a, b = addTwoNumbers2(l1.Next, l2, d1-1, d2)
		n = l1.Val + b
	} else if d1 < d2 {
		a, b = addTwoNumbers2(l1, l2.Next, d1, d2-1)
		n = l2.Val + b
	} else {
		a, b = addTwoNumbers2(l1.Next, l2.Next, d1-1, d2-1)
		n = l1.Val + l2.Val + b
	}
	res = &ListNode{Val: n % 10, Next: a}
	c = n / 10
	return
}

// 堆栈
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2, carry := pushList(l1), pushList(l2), 0
	var newHead *ListNode
	for len(*s1) > 0 || len(*s2) > 0 {
		carry += popStack(s1) + popStack(s2)
		newHead, carry = &ListNode{Val: carry % 10, Next: newHead}, carry/10
	}
	if carry != 0 {
		newHead = &ListNode{Val: carry, Next: newHead}
	}
	return newHead
}

func pushList(head *ListNode) *[]int {
	stack := new([]int)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		*stack = append(*stack, ptr.Val)
	}
	return stack
}

func popStack(stack *[]int) (pop int) {
	if len(*stack) > 0 {
		pop = (*stack)[len(*stack)-1]
		*stack = (*stack)[0 : len(*stack)-1]
	}
	return
}

func main() {

}
