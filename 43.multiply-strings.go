package leetcode

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	n := len(num1)
	m := len(num2)
	ansNum := make([]int, m+n)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			ansNum[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
			ansNum[i+j] += ansNum[i+j+1] / 10
			ansNum[i+j+1] = ansNum[i+j+1] % 10
		}
	}
	var ans []byte
	i := 0
	for i < len(ansNum)-1 && ansNum[i] == 0 {
		i++
	}
	for ; i < len(ansNum); i++ {
		ans = append(ans, byte(ansNum[i]+'0'))
	}
	return string(ans)
}

// @lc code=end
