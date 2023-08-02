package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var inL int = 0
	var inR int = len(inorder) - 1
	var postL int = 0
	var postR int = len(postorder) - 1
	return createTree(inorder, inL, inR, postorder, postL, postR)
}

// 中序+后序 建树
func createTree(inorder []int, inL int, inR int, postorder []int, postL int, postR int) *TreeNode {
	if inL > inR || postL > postR {
		return nil
	}
	result := new(TreeNode)
	result.Val = postorder[postR]

	i := 0

	// 在中序中找到左节点的根节点
	for j, a := range inorder {
		if a == postorder[postR] {
			i = j
		}
	}

	result.Left = createTree(inorder, inL, i-1, postorder, postL, postL+i-inL-1)
	result.Right = createTree(inorder, i+1, inR, postorder, postL+i-inL, postR-1)

	return result
}

func main() {

	fmt.Println(111)
}
