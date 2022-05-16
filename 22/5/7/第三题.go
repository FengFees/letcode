///*
// * Copyright (c) Huawei Technologies Co., Ltd. 2019-2021. All rights reserved.
// * Description: 上机编程认证
// * Note: 缺省代码仅供参考，可自行决定使用、修改或删除
// * 只能import Go标准库
// */
package main

//
//import (
//	"bufio"
//	"errors"
//	"fmt"
//	"io"
//	"os"
//	"strconv"
//	"strings"
//)
//
//var (
//	xNum        = [4]int{-1, 0, 1, 0}
//	yNum        = [4]int{0, 1, 0, -1}
//	maxBlockNum = 0
//	row         = 0
//	col         = 0
//)
//
//// 待实现函数，在此函数中填入答题代码
//func getOptimalSolution(sideLen int, blocks [][]int) int {
//	row, col = sideLen, sideLen
//	dp := make([][]int, row)
//	for i, _ := range dp {
//		dp[i] = make([]int, col)
//		for j, _ := range dp[i] {
//			dp[i][j] = 0
//		}
//	}
//
//	dfs(2,2, dp, blocks, 0, false)
//
//	return maxBlockNum
//}
//
//func dfs(x, y int, dp [][]int, blocks [][]int, maxLen int, endDfs bool) {
//	if endDfs {
//		if maxLen > maxBlockNum {
//			maxBlockNum = maxLen
//		}
//		return
//	}
//
//	dp[x][y] = 1
//	for i := 0; i < 4; i++ {
//		r := x + xNum[i]
//		c := y + yNum[i]
//		if dp[r][c] == 0 &&  {
//
//		}
//	}
//
//}
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	sideLen, blockNum, err := readInputInt(reader)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	blocks, err := readInputIntArrayFromNlines(reader, blockNum, 2)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	result := getOptimalSolution(sideLen, blocks)
//	fmt.Printf("%d", result)
//}
//
//func readInputInt(reader *bufio.Reader) (int, int, error) {
//	var num1, num2 int
//	if _, err := fmt.Fscanf(reader, "%d %d\n", &num1, &num2); err != nil {
//		return 0, 0, err
//	}
//	return num1, num2, nil
//}
//
//func readInputIntArrayFromNlines(reader *bufio.Reader, blockNum int, col int) ([][]int, error) {
//	if blockNum <= 0 {
//		return [][]int{}, nil
//	}
//
//	result := make([][]int, 0, blockNum)
//	for i := 0; i < blockNum; i++ {
//		lineBuf, err := reader.ReadString('\n')
//		if err != nil && err != io.EOF {
//			return nil, err
//		}
//		lineBuf = strings.TrimRight(lineBuf, "\r\n")
//		lineBuf = strings.TrimSpace(lineBuf)
//		ints := map2IntArray(lineBuf, " ")
//		if len(ints) != col {
//			return [][]int{}, errors.New("col len is not " + strconv.Itoa(col) + ".")
//		}
//		result = append(result, ints)
//	}
//	return result, nil
//}
//
//func map2IntArray(str string, dem string) []int {
//	tempArray := strings.Split(str, dem)
//	result := make([]int, len(tempArray))
//	for index, value := range tempArray {
//		value = strings.TrimSpace(value)
//		intVal, err := strconv.Atoi(value)
//		if err == nil {
//			result[index] = intVal
//		}
//	}
//	return result
//}
