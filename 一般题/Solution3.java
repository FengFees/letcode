package com.一般题;

import java.util.Scanner;

/**
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 *      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 */
public class Solution3 {
    public static int lengthOfLongestSubstring(String s) {
        if (s.length()==0)
            return 0;
        if (s.length()==1)
            return 1;
        int sum=1;
        int temp=1;
        int left=0,right=1;
        for ( ; right<=s.length() ; ){
            if (temp > sum)
                sum = temp;
            temp=1;
            for (int j=left ; j<right && right<s.length(); j++ ){
                if (s.charAt(right) == s.charAt(j) ){
                    left=j+1;
                    break;
                }
                temp++;
            }
            right++;
        }
        return sum;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String s = scanner.next();
        int sum = lengthOfLongestSubstring(s);
        System.out.println(sum);
    }
}
