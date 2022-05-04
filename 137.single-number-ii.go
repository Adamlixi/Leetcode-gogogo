package leetcode

/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	one := 0
	two := 0
	three := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		two ^= (one & v)
		one ^= v
		three = ^(one & two)
		one &= three
		two &= three
	}
	return one
}

// @lc code=end
