package leetcode

/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */

// @lc code=start
func find132pattern(nums []int) bool {
	maxC := nums[0]
	minC := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxC {
			maxC = nums[i]
			continue
		}
		if nums[i] < minC {
			maxC = -999999999
			minC = nums[i]
			continue
		}
		if nums[i] < maxC && nums[i] > minC {
			return true
		}
	}
	return false
}

// @lc code=end
