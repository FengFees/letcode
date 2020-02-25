package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	len := len(nums)
	if len==0 {
		return nil
	}

	sort.Ints(nums)
	result := make([][]int,0)
	visited := make([]bool,len)

	backTrack1(&result,nums,visited,make([]int,0,len))
	return result
}

func backTrack1( result *[][]int , nums []int , visited []bool , temp []int )  {
	if len(nums) == len(temp) {
		*result = append(*result,append([]int{},temp ...))
		return
	}

	for i:=0 ; i<len(nums) ; i++  {
		if visited[i] == true {
			continue
		}
		if i>0 && nums[i-1] == nums[i] && visited[i-1] != true {
			continue
		}
		visited[i]=true
		temp = append(temp, nums[i])
		backTrack1(result,nums,visited,temp)
		visited[i]=false
		temp = temp[:len(temp)-1]
	}
}


func main() {
	nums := make([]int,0)
	nums = append(nums, 1,1,2,2)
	result := permuteUnique(nums)
	fmt.Print(result)
}
