package main

import (
	"fmt"
	"sort"
)

//func permuteUnique(nums []int) [][]int {
//	l := len(nums)
//	if l == 0 {
//		return nil
//	}
//
//	sort.Ints(nums)
//	result := make([][]int, 0)
//	visited := make([]bool, l)
//
//	backTrack1(&result, nums, visited, make([]int, 0, l))
//	return result
//}
//
//func backTrack1(result *[][]int, nums []int, visited []bool, temp []int) {
//	if len(nums) == len(temp) {
//		*result = append(*result, append([]int{}, temp...))
//		return
//	}
//
//	for i := 0; i < len(nums); i++ {
//		if visited[i] == true {
//			continue
//		}
//		if i > 0 && nums[i-1] == nums[i] && visited[i-1] != true {
//			continue
//		}
//		visited[i] = true
//		temp = append(temp, nums[i])
//		backTrack1(result, nums, visited, temp)
//		visited[i] = false
//		temp = temp[:len(temp)-1]
//	}
//}

var res2 [][]int
var used2 []bool

func permuteUnique(nums []int) [][]int {
	res2 = [][]int{}
	used2 = make([]bool, len(nums))
	sort.Ints(nums)
	dfs2(0, nums, []int{})
	return res2
}

func dfs2(x int, nums []int, tmp []int) {
	if x == len(nums) {
		t := make([]int, len(nums))
		copy(t, tmp)
		res2 = append(res2, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] && !used2[i-1] {
			continue
		}

		if !used2[i] {
			used2[i] = true
			tmp = append(tmp, nums[i])
			dfs2(x+1, nums, tmp)
			tmp = tmp[:len(tmp)-1]
			used2[i] = false
		}
	}
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1, 1, 2, 2)
	result := permuteUnique(nums)
	fmt.Print(result)
}
