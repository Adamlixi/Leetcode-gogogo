package leetcode

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	numMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && numMap[s[i]] < numMap[s[i+1]] {
			ans += numMap[s[i+1]] - numMap[s[i]]
			i++
		} else {
			ans += numMap[s[i]]
		}
	}
	return ans
}

// @lc code=end
