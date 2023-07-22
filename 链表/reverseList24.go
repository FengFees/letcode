package main

func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = temp
		temp = cur
		cur = next
	}

	return temp
}

func main() {

}
