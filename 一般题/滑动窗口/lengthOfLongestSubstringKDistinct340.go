package main

/*
340. 至多包含 K 个不同字符的最长子串
给定一个字符串 s ，找出 至多 包含 k 个不同字符的最长子串 T。

示例 1:

输入: s = "eceba", k = 2
输出: 3
解释: 则 T 为 "ece"，所以长度为 3。
示例 2:

输入: s = "aa", k = 1
输出: 2
解释: 则 T 为 "aa"，所以长度为 2。


提示：

1 <= s.length <= 5 * 104
0 <= k <= 50
*/
func main() {

}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}
	var ans, maxcnt int
	str := []byte(s)
	mp := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		if mp[str[right]-'a'] == 0 {
			maxcnt++
		}
		mp[str[right]-'a']++
		//maxcnt = maxInt(maxcnt,mp[str[right]-'a'])
		right++

		if maxcnt > k {
			mp[str[left]-'a']--
			// 判断左边移动的那位是否只有一个，若是，则maxcnt需要减一
			if mp[str[left]-'a'] == 0 {
				maxcnt--
			}
			left++
		}

		ans = maxInt(ans, right-left)
	}

	return ans
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
