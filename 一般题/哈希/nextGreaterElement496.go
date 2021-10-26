package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int
	for _, num := range nums1 {
		i := 0
		for t, n := range nums2 {
			if n == num {
				i = t
			}
		}

		j := 0
		for j = i; j < len(nums2); j++ {
			if nums2[j] > num {
				result = append(result, nums2[j])
				break
			}
		}

		if j == len(nums2) {
			result = append(result, -1)
		}
	}

	return result
}

func main() {

}
