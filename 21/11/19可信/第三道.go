package main

import "fmt"

// 最后满足条件的所有组合
var solutions = make([][]int, 0)

func removeBlocks(blocks [][]int, limit int) int {
	if limit == 0 {
		return 0
	}

	if len(blocks) == 0 {
		return 0
	}

	// 思路：回溯查找
	maybeBlocks := make(map[int][]int, limit+1)
	for i := 1; i <= limit; i++ {
		maybeBlocks[i] = make([]int, len(blocks))
	}

	// 遍历时间
	for l := limit; l >= 1; l-- {
		// 是否被拿掉 需要有个数组进行记录
		isMove := make([]bool, len(blocks))

		// 这里需要用到回溯
		for step := len(blocks) - 1; step >= 0; step-- {
			backend(step, 0, blocks, l, isMove, [][]int{{}})
		}
	}

	return len(solutions) - 1
}

// 回溯
func backend(step int, limitTime int, blocks [][]int, limit int, isMove []bool, solution [][]int) {
	// 满足条件时，对结果进行比较
	if (step < 0 || limit < limitTime+blocks[step][2]) && len(solution) > len(solutions)-1 {
		solutions = solution
		if step > 0 {
			backend(step-1, limitTime, blocks, limit, isMove, solution)
		}

		return
	}

	// 从后往前找
	for num := step; num >= 0; num-- {
		// 判断是否存在重叠
		if limitTime+blocks[num][2] <= limit {
			begin := blocks[num][0]
			end := blocks[num][1]
			// 重叠参数
			isTrue := false
			// 从当前位置往后找，是否存在重叠且未被拿掉的积木
			for i := num + 1; i < len(blocks); i++ {
				if checkBlocks(begin, end, blocks[i][0], blocks[i][1]) {
					if !isMove[i] {
						// 重叠了
						isTrue = true
					}
				}
			}
			// 未重叠，才进行下一步操作
			if !isTrue {
				// 判断结束后
				isMove[num] = true
				solution = append(solution, blocks[num])
				limitTime += blocks[num][2]
				backend(step-1, limitTime, blocks, limit, isMove, solution)
				// 回溯
				isMove[num] = false
				solution = solution[:len(solution)-1]
				limitTime -= blocks[num][2]
			}
		}
	}
}

// 判断是否重合
func checkBlocks(begin1 int, end1 int, begin2 int, end2 int) bool {
	if begin1 == begin2 && end1 == end2 {
		return true
	}
	if end2 <= begin1 || begin2 >= end1 {
		// 未重叠
		return false
	}
	// 重叠
	return true
}

func main() {
	blocks := [][]int{{0, 0, 1}, {0, 0, 1}, {0, 0, 1}, {1, 1, 1}, {1, 1, 3}, {1, 1, 3}}
	fmt.Println(removeBlocks(blocks, 4))
}
