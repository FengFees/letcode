package main

/*
139. 单词拆分
给你一个字符串 s 和一个字符串列表 wordDict 作为字典，判定 s 是否可以由空格拆分为一个或多个在字典中出现的单词。

说明：拆分时可以重复使用字典中的单词。



示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false


提示：

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅有小写英文字母组成
wordDict 中的所有字符串 互不相同

*/

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	word := make(map[string]bool)
	dp[0] = true

	var maxLen int
	for _, d := range wordDict {
		word[d] = true
		if len(d) > maxLen {
			maxLen = len(d)
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := i; j >= 0 && j >= i-maxLen; j-- {
			if dp[j] && word[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {

}
