package 查找

func findRepeatNumber(nums []int) int {
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num]++
	}

	for num, cnt := range numsMap {
		if cnt > 1 {
			return num
		}
	}

	return 0
}
