package com.面试题;

/**
 * 面试题 01.06. 字符串压缩
 * 字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
 *
 * 示例1:
 *
 *  输入："aabcccccaaa"
 *  输出："a2b1c5a3"
 * 示例2:
 *
 *  输入："abbccd"
 *  输出："abbccd"
 *  解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
 * 提示：
 *
 * 字符串长度在[0, 50000]范围内。
 */
public class compressString0106 {
    static class Solution {
        public static String compressString(String S) {
            if("".equals(S)) return S;
            int cnt=0;
            int temp=0;
            int right=0;
            String result="";
            char s = S.charAt(right);
            while( right<=S.length()){
                if ( right<S.length() && s == S.charAt(right) ){
                    right++;
                    temp++;
                }else {
                    cnt+=2;
                    if (cnt >= S.length())
                        return S;
                    result = result+s+temp;
                    temp=0;
                    if (right!=S.length())
                        s = S.charAt(right);
                    else break;
                }

            }

            return result;
        }

        public static void main(String[] args) {
            String s = "aabcccccaaa";
            System.out.println(compressString(s));
        }
    }
}
