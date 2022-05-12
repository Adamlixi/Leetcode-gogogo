package leetcode

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	p := 0
	q := 1
	s := 0
	for q < len(nums) {
		if nums[p] != nums[q] {
			nums[s] = nums[p]
			s++
			p++
			q++
		} else {
			nums[s] = nums[q]
			s++
			for q < len(nums) && nums[p] == nums[q] {
				p++
				q++
			}
			if q >= len(nums) {
				nums[s] = nums[p]
				return s + 1
			}
		}
	}
	nums[s] = nums[len(nums)-1]
	return s + 1
}

// @lc code=end
