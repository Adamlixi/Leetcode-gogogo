package leetcode

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i += 1 {
		j := 0
		if i+k < len(nums) {
			j = i + k
		} else {
			j = len(nums) - 1
		}
		for ; j > i; j-- {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	} else {
		return a - b
	}
}

// @lc code=end
