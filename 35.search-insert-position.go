package leetcode

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		n := (low + high) / 2
		if nums[n] == target {
			return n
		} else if nums[n] > target {
			high = n - 1
		} else {
			low = n + 1
		}
	}
	return low
}

// @lc code=end
