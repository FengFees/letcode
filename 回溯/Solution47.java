package com.回溯;

import java.lang.reflect.Array;
import java.util.*;

/**
 *给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 *   [1,1,2],
 *   [1,2,1],
 *   [2,1,1]
 * ]
 *
 *
 */


public class Solution47 {

    public static List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        // 标记该位置是否已经排序过
        int[] visited = new int[nums.length];
        Arrays.sort(nums);
        backtrack(result,nums,new ArrayList<Integer>(),visited);
        return result;
    }

    // 进行剪枝
    private static void backtrack(List<List<Integer>> result, int[] nums, ArrayList<Integer> tmp, int[] visited) {
        if (tmp.size() == nums.length ){
            tmp = new ArrayList<>(tmp);
            // 在全排列1（46）的基础上，进行剪枝（效果较差）该处是结果剪枝，效果不太行。
//            if(result.contains(tmp)) result.remove(tmp);
            result.add(new ArrayList<>(tmp));
            return;
        }

        for (int i=0 ; i<nums.length ; i++ ){
            if ( visited[i] == 1 ) continue;
            if ( i>0 && nums[i-1] == nums[i] && visited[i-1] != 1 ) continue;
            visited[i] = 1;
            tmp.add(nums[i]);
            backtrack(result,nums,tmp,visited);
            visited[i]=0;
            tmp.remove(tmp.size()-1);
        }

    }


    public static void main(String[] args) {
        int[] nums = {3,3,0,3};
        System.out.print(permuteUnique(nums));
    }
}
