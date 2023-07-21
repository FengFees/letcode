package main

import (
	"fmt"
	"math"
	"strings"
)

func strToInt(str string) int {
	result := 0

	// 去掉两边的空格
	s := strings.TrimSpace(str)

	b := []byte(s)

	isNegative := false
	needChange := true

	for index, value := range b {
		if value == '-' && index == 0 {
			isNegative = true
			continue
		}
		if value == '+' && index == 0 {
			continue
		}
		if value <= '9' && value >= '0' {
			num := int(value - '0')
			if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && num > 7) {
				if isNegative == true {
					result = math.MinInt32
				} else {
					result = math.MaxInt32
				}

				needChange = false

				break
			}
			result *= 10
			result += num
		} else {
			break
		}
	}

	if isNegative && needChange {
		result *= -1
	}

	return result
}

func main() {
	fmt.Println(strToInt("9223372036854775808"))
}
