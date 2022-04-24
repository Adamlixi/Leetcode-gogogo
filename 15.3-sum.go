package leetcode

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				tt := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, tt)
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}
		for i < n-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

// @lc code=end
