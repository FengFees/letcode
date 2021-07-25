package main

/*
1004. 最大连续1的个数 III
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。

示例 1：

输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。


提示：

1 <= A.length <= 20000
0 <= K <= A.length
A[i] 为 0 或 1
*/

func main() {

}

/*
是的，这个写法维护的是一个只能单调变长的窗口。这种窗口经常出现在寻求”最大窗口“的问题中：因为要求的是”最大“，所以我们没有必要缩短窗口，于是代码就少了缩短窗口的部分；从另一个角度讲，本题里的K是消耗品，一旦透支，窗口就不能再增长了（也意味着如果K == 0还是有可能增长的）。所以K所代表的”资源“，通常是滑窗维护逻辑的核心，能这么写有两个先决条件：

固定一个左端点，K随窗口增大是单调变化的。据此我们可以推知长度为n的窗口如若已经”透支“（K < 0）了，那么长度大于n的也一定不符合条件；
K的变化与数组元素有简单的算术关系。向窗口纳入（A[r++]）或移除（A[l++]）一个数组元素，可以在O(1)内更新K。
虽说有条件，但仔细排查会发现许多滑窗问题都可以满足
*/
func longestOnes(nums []int, k int) int {
	var ans int
	// 维护一个最长窗口即可
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			k--
		}
		if k < 0 {
			if nums[l] == 0 {
				k++
			}
			k++
		}
		r++
	}
	ans = r - l
	return ans
}
