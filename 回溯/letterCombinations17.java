package com.回溯;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static com.sun.org.apache.xalan.internal.lib.ExsltStrings.split;

/**
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */


// 全排列的回溯思想
public class letterCombinations17 {

    static class Solution {
        public static List<String> letterCombinations(String digits) {
            int len = digits.length();
            char[] nums = digits.toCharArray();
            Map< Character,String > origin = new HashMap<>();
//            origin.put(1,"!@#");
            origin.put('2',"abc");
            origin.put('3',"def");
            origin.put('4',"ghi");
            origin.put('5',"jkl");
            origin.put('6',"mno");
            origin.put('7',"pqrs");
            origin.put('8',"tuv");
            origin.put('9',"wxyz");

            List<String> result = new ArrayList<>();

            if ( "".equals(digits) ) return result;

            backTrack(result,0, digits,origin,new StringBuffer() );

            return result;
        }

        private static void backTrack(List<String> result, int i, String digits, Map<Character, String> origin, StringBuffer temp) {
            if ( temp.length() == digits.length() ){
                result.add(temp.toString());
                return;
            }

            char nums = digits.charAt(i);
            String key = origin.get(nums);

            for (int j=0 ; j<key.length() ; j++ ){
                temp.append(key.charAt(j));
                backTrack(result,i+1,digits,origin,temp);
                temp.deleteCharAt(temp.length()-1);
            }

        }

        public static void main(String[] args) {
            String phone = "233";
            System.out.println(letterCombinations(phone));
        }

    }

}
