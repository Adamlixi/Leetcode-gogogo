package leetcode

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		numSet[nums[i]] = true
	}
	maxCount := 0
	for x, _ := range numSet {
		count := 0
		if _, ok := numSet[x-1]; !ok {
			for {
				if _, ok2 := numSet[x]; ok2 {
					x = x + 1
					count++
					if count > maxCount {
						maxCount = count
					}
				} else {
					break
				}
			}
		}
	}
	return maxCount
}

// @lc code=end
