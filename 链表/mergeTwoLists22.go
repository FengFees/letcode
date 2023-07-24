package main

/*
剑指 Offer 25. 合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	res := result

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			res.Next = l2
			l2 = l2.Next
		} else {
			res.Next = l1
			l1 = l1.Next
		}
		res = res.Next
	}

	for l1 != nil {
		res.Next = l1
		l1 = l1.Next
		res = res.Next
	}

	for l2 != nil {
		res.Next = l2
		l2 = l2.Next
		res = res.Next
	}

	return result.Next
}
