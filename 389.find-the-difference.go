package leetcode

/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	charMap := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; !ok {
			charMap[s[i]] = 1
		} else {
			charMap[s[i]]++
		}
	}
	charMapT := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		if _, ok := charMap[t[i]]; !ok {
			charMapT[t[i]] = 1
		} else {
			charMapT[t[i]]++
		}
	}
	for key, count := range charMapT {
		if cc, ok := charMap[key]; !ok || cc != count {
			return key
		}
	}
	return '0'
}

// @lc code=end
