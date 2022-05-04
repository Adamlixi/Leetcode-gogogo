package leetcode

/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	tank := 0
	for i := 0; i < l; i++ {
		tank += gas[i] - cost[i]
	}
	if tank < 0 {
		return -1
	}
	tank = 0
	start := 0
	for i := 0; i < l; i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	return start
}

// @lc code=end
