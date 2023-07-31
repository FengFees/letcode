package 双指针

import (
	"strings"
)

func reverse(s string) string {
	b := []byte(s)
	start, end := 0, len(b)-1
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
	return string(b)
}

func reverseWords(s string) string {
	str := strings.TrimSpace(s)

	strs := strings.Split(str, " ")

	var temp []string

	for i := 0; i < len(strs); i++ {
		temp = append(temp, reverse(strs[i]))
	}

	tempS := ""

	for _, ss := range temp {
		if ss == "" {
			continue
		}
		tempS += ss + " "
	}

	result := strings.TrimSpace(tempS)
	return reverse(result)
}
