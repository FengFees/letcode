package main

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

 */

func longestPalindrome(s string) string {
	if len(s) < 1 || s == "" {
		return ""
	}
	l := len(s)
	start , end := 0,0
	for i := 0 ; i<l  ; i++ {
		len1 := extendCenter(s, i,i)
		len2 := extendCenter(s,i,i+1)
		len := max(len1,len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start:end+1]
}

func max(num1 int,num2 int)  int {
	if num1>num2 {
		return num1
	}else {
		return num2
	}
}

func extendCenter(s string , left int , right int) int {
	L,R := left,right
	for L>=0 && R<len(s) && s[L] == s[R] {
		L--
		R++
	}

	return R-L-1
}

func main() {

}
