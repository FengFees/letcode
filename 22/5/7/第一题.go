/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2019-2021. All rights reserved.
 * Description: 上机编程认证
 * Note: 缺省代码仅供参考，可自行决定使用、修改或删除
 * 只能import Go标准库
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 待实现函数，在此函数中填入答题代码
func getMaxQuantity(productions []int, plans []int, window int) int {
	length := len(productions)
	allPro := 0
	dp := make([]int, length)
	left := 0
	right := window

	// cul all productions in plans
	for i, pro := range productions {
		if plans[i] == 1 {
			allPro += pro
		}
	}

	if window == 1 {
		noPlan := true
		maxProduction := 0
		for i, production := range productions {
			if plans[i] == 1 {
				noPlan = false
				break
			}

			maxProduction = maxPro(maxProduction, production)
		}

		if noPlan {
			return maxProduction
		}

		return allPro
	}

	if window == 2 {
		maxProduction := 0
		for i, _ := range productions {
			if plans[i] == 1 {
				if i > 0 && i < len(productions)-1 {
					maxProduction = maxPro(allPro+productions[i-1], allPro+productions[i+1])
				} else if i == 0 {
					maxProduction = allPro + productions[i+1]
				} else if i == len(productions)-1 {
					maxProduction = allPro + productions[i-1]
				}
			}
		}

		return maxProduction
	}

	dp[0] = allPro

	for i := right - 1; i >= left; i-- {
		dp[i] = allPro
	}

	// 滑动窗口
	for right <= length {
		tempPro := allPro
		for i := right - 1; i >= left; i-- {
			if plans[i] == 0 {
				tempPro += productions[i]
			}
		}
		dp[right-1] = maxPro(tempPro, dp[right-2])
		right++
		left++
	}

	return dp[length-1]
}

func maxPro(production1, production2 int) int {
	if production1 > production2 {
		return production1
	}
	return production2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	totalDays := readInputInt(reader)
	productions := readInputIntArray(reader, totalDays)
	plans := readInputIntArray(reader, totalDays)
	window := readInputInt(reader)
	result := getMaxQuantity(productions, plans, window)
	fmt.Println(result)
}

func readInputIntArray(reader *bufio.Reader, totalDays int) []int {
	if totalDays == 0 {
		return []int{}
	}
	lineBuf, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
		return nil
	}
	lineBuf = strings.TrimRight(lineBuf, "\r\n")
	lineBuf = strings.TrimSpace(lineBuf)
	intNums := map2IntArray(lineBuf, " ")
	if len(intNums) != totalDays {
		fmt.Println("int string len is error")
		return nil
	}
	return intNums
}

func map2IntArray(str string, dem string) []int {
	tempArray := strings.Split(str, dem)
	result := make([]int, len(tempArray))
	for index, value := range tempArray {
		value = strings.TrimSpace(value)
		intVal, err := strconv.Atoi(value)
		if err == nil {
			result[index] = intVal
		}
	}
	return result
}

func readInputInt(reader *bufio.Reader) int {
	var window int
	if _, err := fmt.Fscanf(reader, "%d\n", &window); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return window
}
