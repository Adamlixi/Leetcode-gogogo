package leetcode

/*
 * @lc app=leetcode id=989 lang=golang
 *
 * [989] Add to Array-Form of Integer
 */

// @lc code=start
func addToArrayForm(num []int, k int) []int {
	nn := 0
	var ans []int
	for i := len(num) - 1; i >= 0; i-- {
		tt := k % 10
		zz := num[i] + nn + tt
		zt := zz % 10
		nn = zz / 10
		k = k / 10
		ans = append(ans, zt)
	}
	for k > 0 {
		zz := k % 10
		tt := zz + nn
		nn = tt / 10
		k /= 10
		ans = append(ans, tt%10)
	}
	if nn > 0 {
		ans = append(ans, nn)
	}
	var an []int
	for i := len(ans) - 1; i >= 0; i-- {
		an = append(an, ans[i])
	}
	return an
}

// @lc code=end
