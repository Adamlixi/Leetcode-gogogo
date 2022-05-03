package leetcode

/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	for {
		if n == 0 {
			return false
		}
		if n == 1 {
			return true
		}
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
}

// @lc code=end
