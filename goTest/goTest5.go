package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ans []byte

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	ans = make([]byte, n)

	dfs(0, n)
}

func dfs(index int, n int) {
	var i, j, k int
	if index == n {
		fmt.Println(string(ans))
		return
	}

	for i = 1; i < 4; i++ {
		ans[index] = byte('0' + i)
		for j = 1; j <= (index+1)/2; j++ {
			for k = 0; k < j; k++ {
				if ans[index-k] != ans[index-j-k] {
					break
				}
			}
			if j == k {
				break
			}
		}
		if j <= (index+1)/2 {
			continue
		}
		dfs(index+1, n)
	}
}
