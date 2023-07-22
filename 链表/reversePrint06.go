package main

import "fmt"

/*
剑指 Offer 06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]


限制：

0 <= 链表长度 <= 10000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

var result []int

func reverseListNode(head *ListNode, temp *ListNode) {
	if head == nil {
		return
	}
	reverseListNode(head.Next, temp)
	temp = head
	//fmt.Println(temp.Val)
	result = append(result, temp.Val)
	temp = temp.Next
}

func reversePrint(head *ListNode) []int {
	var temp *ListNode

	reverseListNode(head, temp)

	return result
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	fmt.Println(reversePrint(head))
}
