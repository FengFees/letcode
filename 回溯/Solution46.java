package com.回溯;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Scanner;

/**
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 *   [1,2,3],
 *   [1,3,2],
 *   [2,1,3],
 *   [2,3,1],
 *   [3,1,2],
 *   [3,2,1]
 * ]
 *
 */

public class Solution46 {

    static class Solution {
        public static List<List<Integer>> permute(int[] nums) {
            List<List<Integer>> result = new LinkedList<>();
            int[] visited = new int[nums.length];
            backtrack(result,nums,new LinkedList<Integer>(),visited);
            return result;
        }
        // 剪枝
        private static void backtrack(List<List<Integer>> result , int[] nums , LinkedList<Integer> temp , int[] visited ){
            if (temp.size() == nums.length ){
                // 这里要new的原因是回溯结果会重置temp 需要重新new链表来存放temp
                result.add(new LinkedList<>(temp));
                return ;
            }

            for (int i=0 ; i<nums.length ; i++ ){
                if (visited[i] == 1 ) continue;
                visited[i] = 1;
                temp.add(nums[i]);
                backtrack(result,nums,temp,visited);
                // 回溯
                visited[i] = 0;
                temp.remove(temp.size()-1);
            }

        }

        public static void main(String[] args) {
//        Scanner scanner = new Scanner(System.in);

            int[] nums = {1,2,3};
            System.out.print(permute(nums));
        }

    }



}
