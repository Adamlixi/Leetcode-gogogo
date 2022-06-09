package leetcode

/*
 * @lc app=leetcode id=696 lang=golang
 *
 * [696] Count Binary Substrings
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		diffCount := 0
		sameCount := 1
		for j := i + 1; j < len(s); j++ {
			if s[j] == c {
				if diffCount > 0 {
					break
				}
				sameCount++
			} else {
				if j == i+1 {
					count++
					break
				}
				diffCount++
			}
			if sameCount == diffCount {
				count++
				break
			}
		}
	}
	return count
}

// @lc code=end
