package com.回溯;

import java.util.ArrayList;
import java.util.List;

/**
 * 842. 将数组拆分成斐波那契序列
 * 给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。
 *
 * 形式上，斐波那契式序列是一个非负整数列表 F，且满足：
 *
 * 0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
 * F.length >= 3；
 * 对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。
 * 另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。
 *
 * 返回从 S 拆分出来的所有斐波那契式的序列块，如果不能拆分则返回 []。
 *
 * 示例 1：
 *
 * 输入："123456579"
 * 输出：[123,456,579]
 * 示例 2：
 *
 * 输入: "11235813"
 * 输出: [1,1,2,3,5,8,13]
 * 示例 3：
 *
 * 输入: "112358130"
 * 输出: []
 * 解释: 这项任务无法完成。
 * 示例 4：
 *
 * 输入："0123"
 * 输出：[]
 * 解释：每个块的数字不能以零开头，因此 "01"，"2"，"3" 不是有效答案。
 * 示例 5：
 *
 * 输入: "1101111"
 * 输出: [110, 1, 111]
 * 解释: 输出 [11,0,11,11] 也同样被接受。
 * 提示：
 *
 * 1 <= S.length <= 200
 * 字符串 S 中只含有数字。
 */
//回溯加剪枝
public class splitIntoFibonacci842 {
    static class Solution {
        public static List<Integer> splitIntoFibonacci(String S) {
            List<Integer> result = new ArrayList<>();
            if (S.charAt(0) == '0') return result; //判断开头是否为0
            if (backTrack(result,0,S)) return result;
            return result;
        }

        public static boolean backTrack(List<Integer> result , int index , String S ){
            if (index == S.length() && result.size()>2) return true;

            for (int i=index ; i<S.length() ; i++ ){
                String segment = S.substring(index,i+1);
                //剪枝部分
                //如果超过数值范围
                if (Long.parseLong(segment) > Integer.MAX_VALUE) break;
                //开头为0 可以不判断单独是否为0的状况，往下递归依旧可以剪枝
                if (!"0".equals(segment) && segment.startsWith("0")) break;
                //判断是否为fibonacci数列
                if (judge(Integer.valueOf(segment),result)){
                    result.add(Integer.valueOf(segment));
                    //递归
                    if (backTrack(result,i+1,S)) return true;
                    //回溯
                    result.remove(result.size()-1);
                }

            }
            return false;
        }

        private static boolean judge(Integer valueOf, List<Integer> result) {
            if (result.size()<2 ) return true;

            return valueOf == result.get(result.size() - 1) + result.get(result.size() - 2);
        }


        public static void main(String[] args) {
            String S="1011";
            System.out.println(splitIntoFibonacci(S));
        }
    }
}
