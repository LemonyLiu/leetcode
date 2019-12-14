package leetcode

/**
题目描述参见：https://leetcode-cn.com/problems/valid-sudoku/
[["8","3",".",".","7",".",".",".","."],
 ["6",".",".","1","9","5",".",".","."],
 [".","9","8",".",".",".",".","6","."],
 ["8",".",".",".","6",".",".",".","3"],
 ["4",".",".","8",".","3",".",".","1"],
 ["7",".",".",".","2",".",".",".","6"],
 [".","6",".",".",".",".","2","8","."],
 [".",".",".","4","1","9",".",".","5"],
 [".",".",".",".","8",".",".","7","9"]]
 */
func IsValidSudoku(board [][]byte) bool {
	cols := make(map[int]map[byte]int)
	rows := make(map[int]map[byte]int)
	chunk := make(map[int]map[byte]int)
	for i:=0;i<9;i++{
		cols[i] = make(map[byte]int)
		rows[i] = make(map[byte]int)
		chunk[i] = make(map[byte]int)
	}

	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j] != '.' {
				if _,ok := cols[i][board[i][j]];ok{
					return false
				} else {
					cols[i][board[i][j]] = 1
				}
				if _,ok := rows[j][board[i][j]];ok{
					return false
				} else {
					rows[j][board[i][j]] = 1
				}
				chunkIndex := (i/3)*3 + j/3
				if _,ok := chunk[chunkIndex][board[i][j]];ok{
					return false
				} else {
					chunk[chunkIndex][board[i][j]] = 1
				}
			}

		}
	}
	return true
}
