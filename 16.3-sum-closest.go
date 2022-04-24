package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	var answer int = 0
	if len(nums) <= 3 {
		for i := 0; i < len(nums); i++ {
			answer += nums[i]
		}
		return answer
	}
	sort.Ints(nums)
	answer = nums[0] + nums[1] + nums[2]
	n := len(nums)
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if Abs(sum-target) < Abs(answer-target) {
				answer = sum
				if answer == target {
					return answer
				}
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return answer
}

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

// @lc code=end
