package leetcode

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	set3Map := make(map[int]map[byte]bool, 0)
	for i := 0; i < len(board); i++ {
		lineSet := make(map[byte]bool, 0)
		for j := 0; j < len(board[i]); j++ {
			key := (i/3)*3 + (j / 3)
			key2 := (j + 1) * 9
			_, ok := set3Map[key]
			if !ok {
				set3Map[key] = make(map[byte]bool, 0)
			}
			_, ok2 := set3Map[key2]
			if !ok2 {
				set3Map[key2] = make(map[byte]bool, 0)
			}
			c := board[i][j]
			if c == '.' {
				continue
			}
			if c > '9' || c < '0' {
				return false
			}
			if _, ok := set3Map[key][c]; ok {
				return false
			} else {
				set3Map[key][c] = true
			}
			if _, ok := set3Map[key2][c]; ok {
				return false
			} else {
				set3Map[key2][c] = true
			}
			if _, ok := lineSet[c]; ok {
				return false
			} else {
				lineSet[c] = true
			}
		}
	}
	return true
}

// @lc code=end
