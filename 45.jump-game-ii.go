package leetcode

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	stepNums := make([]int, len(nums))
	n := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= n {
			stepNums[i] = 1
		} else {
			t := 1
			minSep := 99999
			for t+i < len(nums) && t <= nums[i] {
				if stepNums[t+i] == 0 {
					continue
				}
				if minSep > stepNums[t+i] {
					minSep = stepNums[t+i]
				}
				t++
			}
			stepNums[i] = minSep + 1
		}
		n++
	}
	return stepNums[0]
}

// @lc code=end
