package leetcode

import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start

var ans [][]int

func fourSum(nums []int, target int) [][]int {
	ans = make([][]int, 0)
	var a []int
	sort.Ints(nums)
	TraceBack(a, nums, target, 0)
	return ans
}

func TraceBack(temp []int, nums []int, target int, start int) {
	if target == 0 && len(temp) == 4 {
		var tt []int
		for _, v := range temp {
			tt = append(tt, v)
		}
		ans = append(ans, tt)
		return
	}
	if len(temp) >= 4 {
		return
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		ss := append(temp, nums[i])
		TraceBack(ss, nums, target-nums[i], i+1)
	}
}

// @lc code=end
