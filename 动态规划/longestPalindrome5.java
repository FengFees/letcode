package com.动态规划;

public class longestPalindrome5 {
    public String longestPalindrome(String s) {
        if (s == null || s.length()<1 )  return "";
        int len = s.length();
        int start=0,end=0;

        for (int i=0 ; i<len ; i++){
            int len1 = extendCenter(s,i,i);
            int len2 = extendCenter(s,i,i+1);
            int length = Math.max(len1,len2);
            if(length>end-start){
                start = i - (length-1)/2;
                end = i + length/2;
            }
        }

        return s.substring(start,end+1);
    }

    private int extendCenter(String s, int left, int right) {
        int R=right,L=left;
        while(L>=0 && R<s.length() && s.charAt(L) == s.charAt(R) ){
            L--;
            R++;
        }
        return R-L-1;
    }

    public static void main(String[] args) {

    }
}
