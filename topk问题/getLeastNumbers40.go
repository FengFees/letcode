package topk

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = arr[i]
	}
	return result
}
