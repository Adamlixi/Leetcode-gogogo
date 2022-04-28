package leetcode

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	count := 0
	tt := 0
	for _, v := range s {
		if v != ' ' {
			count++
		} else {
			if count != 0 {
				tt = count
			}
			count = 0
		}
	}
	if count > 0 {
		return count
	} else {
		return tt
	}
}

// @lc code=end
