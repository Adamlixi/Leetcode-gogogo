package leetcode

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ans := 0
	origin := x
	for x > 0 {
		ans *= 10
		t := x % 10
		ans += t
		x /= 10
	}
	return ans == origin
}

// @lc code=end
