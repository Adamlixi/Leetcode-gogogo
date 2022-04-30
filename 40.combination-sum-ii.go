package leetcode

import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
var ans [][]int

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans = make([][]int, 0)
	hasUsed := make([]bool, len(candidates))
	var ss []int
	findAns(candidates, target, hasUsed, ss, 0)
	return ans
}

func findAns(candidates []int, target int, hasUsed []bool, temp []int, start int) {
	if target == 0 {
		var sss []int
		for i := 0; i < len(temp); i++ {
			sss = append(sss, temp[i])
		}
		ans = append(ans, sss)
		return
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		if hasUsed[i] || (i > 0 && !hasUsed[i-1] && candidates[i] == candidates[i-1]) {
			continue
		}
		hasUsed[i] = true
		tt := append(temp, candidates[i])
		findAns(candidates, target-candidates[i], hasUsed, tt, i+1)
		hasUsed[i] = false
	}
}

// @lc code=end
