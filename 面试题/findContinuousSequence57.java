package com.面试题;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;

/**
 *II. 和为s的连续正数序列
 * 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
 *
 * 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
 *
 *  
 *
 * 示例 1：
 *
 * 输入：target = 9
 * 输出：[[2,3,4],[4,5]]
 * 示例 2：
 *
 * 输入：target = 15
 * 输出：[[1,2,3,4,5],[4,5,6],[7,8]]
 *  
 *
 * 限制：
 *
 * 1 <= target <= 10^5
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 *
 */

//滑动窗口
public class findContinuousSequence57 {
    static class Solution {
        public static int[][] findContinuousSequence(int target) {
            List<int[]> result = new ArrayList<>();
            int left=1,right=2;
//            int cnt=0;
            while(left<right && right<=target/2+1){
                if ( (left+right)*(right-left+1)/2 == target ){
                    int[] temp = new int[right-left+1];
                    for (int i=0 ; i<right-left+1 ; i++ )
                        temp[i]=i+left;
                    result.add(temp);
                    left++;
                }else if ( (left+right)*(right-left+1)/2 < target ){
                    right++;
                }else {
                    left++;
                }
            }

            return result.toArray(new int[result.size()][]);
        }

        public static void main(String[] args) {
            System.out.println(findContinuousSequence(9));
        }
    }


}
