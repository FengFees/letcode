package main

/*
229. 求众数 II
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

示例 1：

输入：[3,2,3]
输出：[3]
示例 2：

输入：nums = [1]
输出：[1]
示例 3：

输入：[1,1,1,3,3,2,2,2]
输出：[1,2]


提示：

1 <= nums.length <= 5 * 104
-109 <= nums[i] <= 109

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
*/

func majorityElement(nums []int) []int {
	var result []int
	n := len(nums) / 3
	cache := make(map[int]int)
	for _, v := range nums {
		cache[v]++
	}
	// 使用map来存次数即可
	for k, v := range cache {
		if v > n {
			result = append(result, k)
		}
	}

	return result
}

func main() {

}
