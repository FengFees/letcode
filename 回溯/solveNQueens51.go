package main

// n 皇后
/*
n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。

每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]


提示：

1 <= n <= 9
皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

*/

var result [][]string

func solveNQueens(n int) [][]string {
	result = [][]string{}
	if n <= 0 {
		return result
	}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	ds1, ds2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, ds1, ds2)

	return result
}

func backtrack(queens []int, n int, row int, columns map[int]bool, ds1 map[int]bool, ds2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		result = append(result, board)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		d1 := row - i
		if ds1[d1] {
			continue
		}
		d2 := row + i
		if ds2[d2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		ds1[d1], ds2[d2] = true, true
		backtrack(queens, n, row+1, columns, ds1, ds2)
		// 回溯
		queens[row] = -1
		delete(columns, i)
		delete(ds1, d1)
		delete(ds2, d2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}

	return board
}

func main() {

}
