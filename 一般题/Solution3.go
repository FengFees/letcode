package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	judge := make([]bool,128)
	var curLen,maxLen int
	curLen,maxLen = 0,0
	for left,right := 0,0; right <len(s) ; right++  {
		//若right位置的字符没有出现过，则往后搜索
		if !judge[bytes[right]] {
			judge[bytes[right]] = true
			curLen++
		} else {
				//记录right的位置
			temp := bytes[right]
			if curLen > maxLen {
				maxLen = curLen
			}
			//修改left的位置
			for {
				if bytes[left] != temp {
					judge[bytes[left]] = false
					curLen--
					left++
				} else {
					left++
					break
				}
			}

		}
	}
	if curLen > maxLen {
		maxLen=curLen
	}

	return maxLen
}

func main()  {
	fmt.Print(lengthOfLongestSubstring("blbk"))
}