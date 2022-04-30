package leetcode

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
var ans [][]int

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans = make([][]int, 0)
	var t []int
	hasNum := make([]bool, len(nums))
	permutationsYY(nums, t, hasNum)
	return ans
}

func permutationsYY(nums []int, temp []int, hasNum []bool) {
	if len(temp) == len(nums) {
		var aa []int
		for _, v := range temp {
			aa = append(aa, v)
		}
		ans = append(ans, aa)
		return
	}
	if len(temp) > len(nums) {
		return
	}
	for i := 0; i < len(nums); i++ {
		if hasNum[i] || (i > 0 && !hasNum[i-1] && nums[i] == nums[i-1]) {
			continue
		}
		temp = append(temp, nums[i])
		hasNum[i] = true
		permutationsYY(nums, temp, hasNum)
		hasNum[i] = false
		temp = temp[:len(temp)-1]
	}
}

// @lc code=end
