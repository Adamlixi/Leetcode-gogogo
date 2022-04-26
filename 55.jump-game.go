package leetcode

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	arriveNum := make([]bool, len(nums))
	n := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= n {
			arriveNum[i] = true
			n = 1
		} else {
			arriveNum[i] = false
			n++
		}
	}
	return arriveNum[0]
}

// @lc code=end
