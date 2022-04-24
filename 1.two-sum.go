package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numSet := make(map[int]int, len(nums))
	for i, v := range nums {
		numSet[v] = i
	}
	var ans []int
	for i, v := range nums {
		if j, ok := numSet[target-v]; ok && i != j {
			ans = append(ans, i, j)
			return ans
		}
	}
	return nil
}

// @lc code=end
