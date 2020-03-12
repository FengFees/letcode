package com.树;

public class diameterOfBinaryTree543 {

    static class Solution {

        public static class TreeNode {
            int val;
            TreeNode left;
            TreeNode right;
            TreeNode(int x) { val = x; }
        }

        public static int ans;
        public static int diameterOfBinaryTree(TreeNode root) {
            ans=1;
            if ( root == null ) return 0;
//            int max=maxHeight(root.left)+maxHeight(root.right);
            maxHeight(root);
//            return max>diameterOfBinaryTree(root.left) ? (max>diameterOfBinaryTree(root.right) ? max : diameterOfBinaryTree(root.right)) :diameterOfBinaryTree(root.left);
            // 超时哈哈哈
            return ans-1;
        }

        public static int maxHeight(TreeNode root){
            if(root != null ){
                int maxL=maxHeight(root.left);
                int maxR=maxHeight(root.right);
                ans = Math.max(ans,maxL+maxR+1);
                return maxL>maxR?maxL+1:maxR+1;
            }
            return 0;
        }

        public static void main(String[] args) {
            TreeNode treeNode=new TreeNode(1);
            treeNode.left=new TreeNode(2);
            treeNode.right=new TreeNode(3);
            treeNode.left.left=new TreeNode(4);
            treeNode.left.right=new TreeNode(5);
            System.out.println(diameterOfBinaryTree(treeNode));
        }

    }
}
