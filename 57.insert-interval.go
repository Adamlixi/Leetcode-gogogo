package leetcode

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	var left [][]int
	var right [][]int
	l := newInterval[0]
	r := newInterval[1]
	for _, v := range intervals {
		if l > v[1] {
			left = append(left, v)
		} else if r < v[0] {
			right = append(right, v)
		} else {
			l = minNum(v[0], l)
			r = maxNum(v[1], r)
		}
	}
	ll := []int{l, r}
	left = append(left, ll)
	left = append(left, right...)
	return left
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minNum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
