package main

/*
424. 替换后的最长重复字符
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，
总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意：字符串长度 和 k 不会超过 104。

示例 1：

输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。
示例 2：

输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。
*/

func main() {

}

/*
方法：双指针（滑动窗口）
右边界先移动找到一个满足题意的可以替换 k 个字符以后，所有字符都变成一样的当前看来最长的子串，直到右边界纳入一个字符以后，不能满足的时候停下；
然后考虑左边界向右移动，左边界只须要向右移动一格以后，右边界就又可以开始向右移动了，继续尝试找到更长的目标子串；
替换后的最长重复子串就产生在右边界、左边界交替向右移动的过程中。

复杂度分析：

时间复杂度：O(N)O(N)，这里 NN 是输入字符串 S 的长度；
空间复杂度：O(A)O(A)，这里 AA 是输入字符串 S 出现的字符 ASCII 值的范围（感谢用户 @hyponarch 指正）。
以下是我们在编码的过程中思考的一些问题。我们建议大家先思考，通过调试，理解代码结果正确的原因。
欢迎大家参与讨论。

1. 证明：如果长度为 L 的子串不符合题目的要求，那么左边界固定，长度更长的子串也不符合题目的要求。
答：记 count(X)count(X) 表示长度为 L 的子串中，字符 X 出现的次数。

不失一般性，假设长度为 L 的子串，出现最多的字符为 A，记 count(A) = xcount(A)=x。
其余字符均为 B，记 count(B) = ycount(B)=y。由字符 A 出现次数最多，可知 x \ge yx≥y。
又由于长度为 L 的子串不符合题目的要求，可知 y > ky>k。起点固定的情况下，考虑更长的子串：

如果接下来看到的字符都是 A（频数最多的字符越来越多），依然须要考虑把之前看到的 B 全部替换成为 A，
由于 count(B) = y > kcount(B)=y>k，这是不能做到的；
如果接下来看到的字符不是 A（频数较少的字符超过原来频数最多的字符），
那么须要考虑把之前看到的 A 全部替换成为新的频数最多的字符，
由于 count(A) = x \ge y > kcount(A)=x≥y>k，这也是不能做到的。

说明：这里我们只讨论了滑动窗口扫过的子区间只含有 2 种字符的情况，
如果滑动窗口扫过的子区间只含有 3 种以及 3 种以上字符，讨论是类似的。

2. maxCount 在内层循环「左边界向右移动一个位置」的过程中，没有维护它的定义，结论是否正确？
答：结论依然正确。「左边界向右移动一个位置」的时候，maxCount 或者不变，或者值减 11。

maxCount 的值虽然不维护，但数组 freq 的值是被正确维护的；
当「左边界向右移动」之前：
如果有两种字符长度相等，左边界向右移动不改变 maxCount 的值。例如 s = [AAABBB]、k = 2，
左边界 A 移除以后，窗口内字符出现次数不变，依然为 33；
如果左边界移除以后，使得此时 maxCount 的值变小，又由于 我们要找的只是最长替换 k 次以后重复子串的长度。
接下来我们继续让右边界向右移动一格，有两种情况：① 右边界如果读到了刚才移出左边界的字符，
恰好 maxCount 的值被正确维护；② 右边界如果读到了不是刚才移出左边界的字符，
新的子串要想在符合题意的条件下变得更长，maxCount 一定要比之前的值还要更多，因此不会错过更优的解。


3. 内层循环里的 if 能不能改成 while?
答：可以但没有必要。理由依然是：我们只关心最长替换 k 次以后重复子串的长度。

正是因为多读了一个字符，使得 right - left > maxCount + k 成立；
在 left++ 以后，由于可以不维护 maxCount 的定义，right - left > maxCount + k 不成立。因此 if 里面的代码块只会被执行一次。
4. 可以不用一直用 res 记录滑动窗口的最大长度，最后返回 right - left 即可。
答：依然是 我们只关心最长替换 k 次以后重复子串的长度，并且 maxCount 只会增加不会减少。
在退出内层 if 语句的时候，区间 [left, right) 不一定是符合要求的子串，
但是子串的长度一定等于题目要求的替换 k 次以后字符全都相等的最长子串（maxCount 的值不会变小，
所以它会一直撑着滑动窗口的长度直到 right 遍历到字符串的末尾）。
这一点如果很难理解的话，我们建议大家使用小测试数据、跟踪代码进行理解。

*/
func characterReplacement(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}
	var ans, maxcnt int
	mp := make(map[byte]int) // 计该长度下的字母有多少
	str := []byte(s)
	left, right := 0, 0 // 滑动窗口
	for right < len(str) {
		mp[str[right]-'A']++                        // 计数
		maxcnt = maxint(maxcnt, mp[str[right]-'A']) // 判断最大字母数
		right++                                     // 右边界移动

		if right-left > maxcnt+k {
			// 当right-left 比 最大字母数和k之和还要大时，说明k已经不够了，需要一定左边界
			mp[str[left]-'A']--
			left++
		}
		ans = maxint(ans, right-left)

	}

	return ans
}

func maxint(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
