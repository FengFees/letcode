package main

/*
剑指 Offer 51. 数组中的逆序对
困难
1K
相关企业
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。



示例 1:

输入: [7,5,6,4]
输出: 5


限制：

0 <= 数组长度 <= 50000
*/

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left == right {
		return 0
	}

	mid := left + (right-left)/2
	leftCnt := mergeSort(nums, left, mid)
	rightCnt := mergeSort(nums, mid+1, right)

	cnt := 0
	var tmp []int
	i, j := left, mid+1
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - mid - 1
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	for i <= mid {
		tmp = append(tmp, nums[i])
		cnt += right - mid
		i++
	}
	for j <= right {
		tmp = append(tmp, nums[j])
		j++
	}

	for t := left; t <= right; t++ {
		nums[t] = tmp[t-left]
	}

	return leftCnt + rightCnt + cnt
}
