package main

import "fmt"

func backTrack(result *[][]int ,nums []int , temp []int , visited []bool) {
	if len(temp) == len(nums) {
		*result = append(*result,append([]int{},temp...))
		return
	}

	for i:= 0 ; i<len(nums) ; i++ {
		if visited[i] == true {
			continue
		}
		visited[i] = true
		temp = append(temp,nums[i])
		backTrack(result,nums,temp,visited)
		visited[i] = false
		temp = temp[:len(temp)-1]
	}

}

func permute(nums []int) [][]int {
	len:= len(nums)
	if len==0 {
		return nil
	}
	visited := make([]bool, len)
	result := make([][]int , 0)
	backTrack(&result,nums,make([]int,0,len),visited)
	return result
}

func main() {
	nums := make([]int,0)
	nums = append(nums,1,2,3)
	fmt.Print(nums)
	result := permute(nums)
	fmt.Print(result)
}
