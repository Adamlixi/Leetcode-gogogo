package leetcode

/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}
	if s != goal && len(s) == 1 {
		return false
	}
	n := len(s)
	for i := 0; i < n; i++ {
		a := s[0]
		s = s[1:]
		s += string(a)
		if s == goal {
			return true
		}
	}
	return false
}

// @lc code=end
