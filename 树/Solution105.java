package com.树;

/**
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 * 返回如下的二叉树：
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 */


public class Solution105 {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    // 先序+中序创建树
    public static TreeNode buildTree(int[] preorder, int[] inorder) {
        int preL=0,preR=preorder.length-1,inL=0,inR=inorder.length-1;

        return createTree(preorder,preL,preR,inorder,inL,inR);
    }

    private static TreeNode createTree(int[] preorder, int preL, int preR, int[] inorder, int inL, int inR) {
        if (preL>preR || inL>inR ) return null;

        TreeNode treeNode = new TreeNode(preorder[preL]);

        int i;
        // 寻找中序遍历的根节点
        for ( i=inL ; i<=inR ; i++ ){
            if(inorder[i] == preorder[preL])
                break;
        }

        int n= i-inL;
        // i左边的为左子树
        treeNode.left = createTree(preorder,preL+1,preR+n,inorder,inL,i-1);
        // i右边的为右子树 先序需要往前走n个数字
        treeNode.right = createTree(preorder,preL+n+1,preR,inorder,i+1,inR);

        return treeNode;
    }

    public static void main(String[] args) {
        int[] pre={3,9,20,15,7};
        int[] inorder={9,3,15,20,7};

        TreeNode treeNode = buildTree(pre,inorder);
    }

}
