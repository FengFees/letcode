package main

import (
	"fmt"
)

/*
46. 全排列
中等
2.6K
相关企业
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

//func backTrack(result *[][]int, nums []int, temp []int, visited []bool) {
//	if len(temp) == len(nums) {
//		*result = append(*result, append([]int{}, temp...))
//		return
//	}
//
//	for i := 0; i < len(nums); i++ {
//		if visited[i] == true {
//			continue
//		}
//		visited[i] = true
//		temp = append(temp, nums[i])
//		backTrack(result, nums, temp, visited)
//		visited[i] = false
//		temp = temp[:len(temp)-1]
//	}
//
//}
//
//func permute(nums []int) [][]int {
//	l := len(nums)
//	if l == 0 {
//		return nil
//	}
//	visited := make([]bool, l)
//	result := make([][]int, 0)
//	backTrack(&result, nums, make([]int, 0, l), visited)
//	return result
//}

var res [][]int
var used []bool

func permute(nums []int) [][]int {
	res = [][]int{}
	used = make([]bool, len(nums))
	dfs(0, nums, []int{})
	return res
}

func dfs(x int, nums []int, tmp []int) {
	if x == len(nums) {
		t := make([]int, len(nums))
		copy(t, tmp)
		res = append(res, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			tmp = append(tmp, nums[i])
			dfs(x+1, nums, tmp)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3)
	fmt.Print(nums)
	result := permute(nums)
	fmt.Print(result)
}
