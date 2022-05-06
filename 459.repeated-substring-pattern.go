package leetcode

/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	if len(s) == 2 {
		if s[0] == s[1] {
			return true
		} else {
			return false
		}
	}
	for i := 2; i <= len(s); i++ {
		if len(s)%i == 0 {
			n := len(s) / i
			ss := s[:n]
			ans := ""
			for j := 0; j < i; j++ {
				ans += ss
			}
			if ans == s {
				return true
			}
		}
	}
	return false
}

// @lc code=end
