package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
剑指 Offer 20. 表示数值的字符串
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

数值（按顺序）可以分成以下几个部分：

若干空格
一个 小数 或者 整数
（可选）一个 'e' 或 'E' ，后面跟着一个 整数
若干空格
小数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
下述格式之一：
至少一位数字，后面跟着一个点 '.'
至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
一个点 '.' ，后面跟着至少一位数字
整数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
至少一位数字
部分数值列举如下：

["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：

["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]


示例 1：

输入：s = "0"
输出：true
示例 2：

输入：s = "e"
输出：false
示例 3：

输入：s = "."
输出：false
示例 4：

输入：s = "    .1  "
输出：true


提示：

1 <= s.length <= 20
s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，空格 ' ' 或者点 '.' 。

*/

func isNumber(s string) bool {
	// 去掉两边的空格
	s = strings.TrimSpace(s)

	// 先判断小数点
	splitNum := strings.Split(s, ".")

	// 1.1.1  -1.1... ....1
	if len(splitNum) > 2 {
		return false
	}

	if len(splitNum) == 1 {
		matched, err := regexp.MatchString("(^[+-]?\\d+(?:[eE][+-]?\\d+)?$)|(^[+-]\\d$)", splitNum[0])
		if !matched || err != nil {
			return false
		}
	}

	if len(splitNum) == 2 {
		// .123 .123e+10
		if splitNum[0] == "" {
			matched, err := regexp.MatchString("^\\d+(?:[eE][+-]?\\d+)?$", splitNum[1])
			if !matched || err != nil {
				return false
			}
		} else {

			if splitNum[1] == "" {
				// 123e-1.  0
				matched, err := regexp.MatchString("(^[+-]?\\d+$)", splitNum[0])
				if !matched || err != nil {
					return false
				}
			} else {
				// +123.123  -123.123
				matched, err := regexp.MatchString("^[+-]?[0-9]*?$", splitNum[0])
				if !matched || err != nil {
					return false
				}
				matched, err = regexp.MatchString("(^(\\d+)?[eE]?[+-]?[0-9][0-9]*$)|(^[0-9]*$)", splitNum[1])
				if !matched || err != nil {
					return false
				}
			}
		}
	}

	return true
}

func main() {

	fmt.Println(isNumber("2e0"))

}
