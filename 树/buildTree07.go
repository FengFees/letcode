package main

/*
剑指 Offer 07. 重建二叉树
中等
1.1K
相关企业
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



示例 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
示例 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]


限制：

0 <= 节点个数 <= 5000



注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 先序+中序构造数
func buildTree2(preorder []int, inorder []int) *TreeNode {

	return createTree2(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func createTree2(preorder []int, prel, pree int, inorder []int, inl, ine int) *TreeNode {
	if prel > pree || inl > ine {
		return nil
	}

	tree := &TreeNode{Val: preorder[prel]}

	var i int // index
	for i = inl; i < ine; i++ {
		// 如果在中序数组中找到了根节点
		if preorder[prel] == inorder[i] {
			break
		}
	}

	n := i - inl
	// 中序左边n个节点都为prel节点的左子树
	tree.Left = createTree2(preorder, prel+1, prel+n, inorder, inl, i-1)
	// 其余的都是prel节点的右子树
	tree.Right = createTree2(preorder, prel+n+1, pree, inorder, i+1, ine)
	return tree
}
