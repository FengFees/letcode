package main

/*
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var mapNodeCache map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := mapNodeCache[node]; ok {
		return n
	}

	temp := &Node{Val: node.Val}
	mapNodeCache[node] = temp
	temp.Next = deepCopy(node.Next)
	temp.Random = deepCopy(node.Random)
	return temp
}

func copyRandomList(head *Node) *Node {
	mapNodeCache = map[*Node]*Node{}
	return deepCopy(head)
}

func main() {

}
