package leetcode

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
var ans [][]int

func combine(n int, k int) [][]int {
	ans = [][]int{}
	traceBack(n, k, 1, []int{})
	return ans
}

func traceBack(n, k, start int, temp []int) {
	if k == len(temp) {
		var s []int
		for _, v := range temp {
			s = append(s, v)
		}
		ans = append(ans, s)
		return
	}
	for i := start; i <= n; i++ {
		a := i
		tt := append(temp, a)
		a++
		traceBack(n, k, a, tt)
	}
}

// @lc code=end
