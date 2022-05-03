package leetcode

/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	mapZero := make(map[int][]int, 0)
	for i := 0; i < len(matrix); i++ {
		mapZero[i] = make([]int, 0)
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				mapZero[i] = append(mapZero[i], j)
			}
		}
	}
	for key, v := range mapZero {
		for i := 0; i < len(matrix); i++ {
			for _, t := range v {
				matrix[i][t] = 0
			}
		}
		if len(v) == 0 {
			continue
		}
		matrix[key] = make([]int, len(matrix[key]))
	}
}

// @lc code=end
