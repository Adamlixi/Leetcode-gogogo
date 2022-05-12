package leetcode

/*
 * @lc app=leetcode id=565 lang=golang
 *
 * [565] Array Nesting
 */

// @lc code=start
func arrayNesting(nums []int) int {
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		size := 0
		for k := i; nums[k] >= 0; size++ {
			ak := nums[k]
			nums[k] = -1
			k = ak
		}
		if size > maxCount {
			maxCount = size
		}
	}
	return maxCount
}

// @lc code=end
