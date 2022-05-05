package leetcode

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	c := make(chan bool, 1)
	for _, v := range matrix {
		go func(t []int) {
			for i := 0; i < len(t); i++ {
				if t[i] == target {
					c <- true
					return
				}
				if t[i] > target {
					c <- false
					return
				}
			}
			c <- false
		}(v)
	}
	count := 0
	for {
		select {
		case b := <-c:
			if b {
				return true
			} else {
				count++
				if count == len(matrix) {
					return false
				}
			}
		}
	}
}

// @lc code=end
