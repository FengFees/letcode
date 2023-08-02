package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
剑指 Offer 17. 打印从1到最大的n位数
简单
304
相关企业
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


说明：

用返回一个整数列表来代替打印
n 为正整数
*/

var result []int
var num []byte
var char = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func printNumbers(n int) []int {
	result = []int{}
	num = make([]byte, n)
	dfs(0, n)
	return result
}

func dfs(x, n int) {
	if x == n {
		s := string(num)
		// 将数字前面的0删除
		s = strings.TrimLeft(s, "0") // 对非0数字进行剪枝
		//fmt.Println(s)
		if s != "" { // 对非0数字进行剪枝
			number, _ := strconv.Atoi(s)
			result = append(result, number)
		}
		return
	}

	// 全排列,将所有的数字都排列出来
	for _, b := range char {
		num[x] = b
		dfs(x+1, n)
	}
}

func main() {
	fmt.Println(printNumbers(3))
}
