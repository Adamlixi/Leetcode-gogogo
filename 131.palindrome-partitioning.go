package leetcode

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start

var ans [][]string

func partition(s string) [][]string {

}

func combineTT(ss []string) {
	if len(ss) == 1 {
		ans = append(ans, ss)
		return
	}
	for i := 1; i < len(ss)-1; i++ {
		if ss[i-1] == ss[i+1] {
			t := ss[i-1] + ss[i] + ss[i+1]
			var aa []string
			if i == 1 {
				if len(ss) == 3 {
					combineTT([]string{t})
				}
				aa = aa[i+2:]
				combineTT(aa)
			} else if i == len(ss)-2 {

			}
		}
	}
}

// @lc code=end
