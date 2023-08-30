package 单调栈

/*
42. 接雨水
困难
4.6K
相关企业
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



示例 1：



输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9


提示：

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

// 单调栈，单调递减，若存在高低高的情况，则可以接到雨水
func trap(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1] // 低
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]                     // 高
			curWidth := i - left - 1                        // 计算宽度
			curHeight := min(height[left], h) - height[top] //计算高度
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
