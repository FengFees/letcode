package com.树;

public class Solution106 {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    public static TreeNode buildTree(int[] inorder, int[] postorder) {
        int postL=0,postR=postorder.length-1,inL=0,inR=inorder.length-1;
        return createTree(inorder,inL,inR,postorder,postL,postR);
    }

    private static TreeNode createTree(int[] inorder, int inL, int inR, int[] postorder, int postL, int postR) {
        if (postL>postR || inL>inR ) return null;
        TreeNode result = new TreeNode(postorder[postR]);

        int i;
        for (i=inL ; i<inR ; i++ ){
            if(inorder[i]==postorder[postR])
                break;
        }

        result.left=createTree(inorder,inL,i-1,postorder,postL,postL+i-inL-1);
        result.right=createTree(inorder,i+1,inR,postorder,postL+i-inL,postR-1);

        return result;
    }

    public static void main(String[] args) {
        int[] post={9,15,7,20,3};
        int[] inorder={9,3,15,20,7};

        TreeNode treeNode = buildTree(inorder,post);

        System.out.println(treeNode);
    }

}
