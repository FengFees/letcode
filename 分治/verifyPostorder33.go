package 分治

/*
剑指 Offer 33. 二叉搜索树的后序遍历序列
中等
744
相关企业
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000
*/

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}

	// 找到根节点
	root := len(postorder) - 1
	value := postorder[root]

	for i, v := range postorder {
		// 找到比根节点大的节点，这些节点为root的右子树
		if root == len(postorder)-1 && v > value {
			root = i
		}
		// 若右子树中存在比根节点还大的节点，则该树不为搜索树
		if root != len(postorder)-1 && value > v {
			return false
		}
	}
	// 分治判断
	return verifyPostorder(postorder[:root]) && verifyPostorder(postorder[root:len(postorder)-1])
}
