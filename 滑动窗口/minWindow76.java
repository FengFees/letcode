package com.一般题;

import java.util.Arrays;

/**
 * 76. 最小覆盖子串
 * 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
 *
 * 示例：
 *
 * 输入: S = "ADOBECODEBANC", T = "ABC"
 * 输出: "BANC"
 * 说明：
 *
 * 如果 S 中不存这样的子串，则返回空字符串 ""。
 * 如果 S 中存在这样的子串，我们保证它是唯一的答案。
 */
public class minWindow76 {
    public static String minWindow(String s, String t) {
        int left=0,right=t.length()-1;
        String result = s;
        if (t.length()==0 || s.length()==0) return result;

        while (left<=right && right<s.length()) {
            int maxSite=0;
            int flag=0;
            boolean[] boo = new boolean[t.length()];
            Arrays.fill(boo,false);
            String sub = s.substring(left,right);
            for (int j = 0; j < t.length() ; j++) {
        if (sub.contains(String.valueOf(t.charAt(j))) ){
            maxSite = Math.max(sub.indexOf(t.charAt(j)), maxSite);
        }else {
            right++;
            flag=1;
        }

        if (flag == 0){

        }
    }
}

        return result;
    }

    public static void main(String[] args) {

    }
}
