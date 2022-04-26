package leetcode

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	var totalNumber []int
	numberMap := make(map[int][]int, 0)
	for _, v := range intervals {
		if val, ok := numberMap[v[0]]; ok {
			if val[1] < v[1] {
				numberMap[v[0]] = v
			} else {
				continue
			}
		} else {
			totalNumber = append(totalNumber, v[0], v[1])
			numberMap[v[0]] = v
		}
	}
	sort.Ints(totalNumber)
	left := numberMap[totalNumber[0]][0]
	right := numberMap[totalNumber[0]][1]
	var ans [][]int
	for _, v := range totalNumber {
		if v >= left && v <= right {
			if r, ok := numberMap[v]; ok {
				if r[1] > right {
					right = r[1]
				}
			}
		} else {
			ans = append(ans, []int{left, right})
			left = numberMap[v][0]
			right = numberMap[v][1]
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}

// @lc code=end
