package leetcode

import "sort"

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
	ans = make([][]int, 0)
	sort.Ints(candidates)
	var tt []int
	calculateNANA(candidates, target, 0, tt)
	return ans
}

func calculateNANA(candidates []int, remain int, start int, tt []int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		ss := make([]int, 0)
		for _, s := range tt {
			ss = append(ss, s)
		}
		ans = append(ans, ss)
		return
	} else {
		for i := start; i < len(candidates); i++ {
			tt = append(tt, candidates[i])
			calculateNANA(candidates, remain-candidates[i], i, tt)
			tt = tt[:len(tt)-1]
		}
	}
}

// @lc code=end
