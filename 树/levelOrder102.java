package com.树;

import com.sun.jmx.remote.internal.ArrayQueue;
import sun.reflect.generics.tree.Tree;

import java.util.*;

/**
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 * 例如:
 * 给定二叉树: [3,9,20,null,null,15,7],
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回其层次遍历结果：
 *
 * [
 *   [3],
 *   [9,20],
 *   [15,7]
 * ]
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class levelOrder102 {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }
    //class Solution {
    //  public List<List<Integer>> levelOrder(TreeNode root) {
    //    List<List<Integer>> levels = new ArrayList<List<Integer>>();
    //    if (root == null) return levels;
    //
    //    Queue<TreeNode> queue = new LinkedList<TreeNode>();
    //    queue.add(root);
    //    int level = 0;
    //    while ( !queue.isEmpty() ) {
    //      // start the current level
    //      levels.add(new ArrayList<Integer>());
    //
    //      // number of elements in the current level
    //      int level_length = queue.size();
    //      for(int i = 0; i < level_length; ++i) {
    //        TreeNode node = queue.remove();
    //
    //        // fulfill the current level
    //        levels.get(level).add(node.val);
    //
    //        // add child nodes of the current level
    //        // in the queue for the next level
    //        if (node.left != null) queue.add(node.left);
    //        if (node.right != null) queue.add(node.right);
    //      }
    //      // go to next level
    //      level++;
    //    }
    //    return levels;
    //  }
    //}
    //

    // 方法一：迭代： 时间：O(N) 空间：O(N)
//    static class Solution {
//        public static List<List<Integer>> levelOrder(TreeNode root) {
//            List<List<Integer>> result = new ArrayList<>();
//            if (root == null ) return result;
//
//            Queue<TreeNode> queue = new LinkedList<>();
//            queue.add(root);
//            // 层数
//            int l=0;
//
//            while( !queue.isEmpty() ){
//                result.add(new ArrayList<>());
//                int len = queue.size();
//                for (int i=0 ; i<len ; i++ ){
//                    TreeNode p = queue.remove();
//                    result.get(l).add(p.val);
//
//                    if( p.left != null ){
//                        queue.add(p.left);
//                    }
//
//                    if( p.right != null ){
//                        queue.add(p.right);
//                    }
//
//                }
//                // 该层遍历完 层数加一
//                l++;
//            }
//
//            return result;
//        }



//        public static void main(String[] args) {
//            TreeNode treeNode = new TreeNode(3);
//            treeNode.left = new TreeNode(9);
//            treeNode.right = new TreeNode(20);
//            List<List<Integer>> result = levelOrder(treeNode);
//            System.out.println(result);
//        }
//    }
    /**
     * public List<List<Integer>> levelOrder(TreeNode root) {
     *         List<List<Integer>> res = new ArrayList<List<Integer>>();
     *         levelHelper(res, root, 0);
     *         return res;
     *     }
     *
     *     public void levelHelper(List<List<Integer>> res, TreeNode root, int height) {
     *         if (root == null) return;
     *         if (height >= res.size()) {
     *             res.add(new LinkedList<Integer>());
     *         }
     *         res.get(height).add(root.val);
     *         levelHelper(res, root.left, height+1);
     *         levelHelper(res, root.right, height+1);
     *     }
     */
    //方法二：递归(本质为dfs)  O(N) O(N)
    static class Solution1{

        public static void levelOrderHelper(List<List<Integer>> result,TreeNode root, Integer level){
            if (root == null ) return;

            if (level >= result.size()) result.add(new ArrayList<>());

            result.get(level).add(root.val);
            levelOrderHelper(result,root.left,level+1);
            levelOrderHelper(result,root.right,level+1);

        }

        public static List<List<Integer>> levelOrder1(TreeNode root){
            List<List<Integer>> result = new ArrayList<>();
            //递归
            levelOrderHelper(result,root,0);

            return result;
        }

        public static void main(String[] args) {
            TreeNode treeNode = new TreeNode(3);
            treeNode.left = new TreeNode(9);
            treeNode.right = new TreeNode(20);
            List<List<Integer>> result = levelOrder1(treeNode);
            System.out.println(result);
        }
    }

}
