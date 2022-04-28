package leetcode

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func partition(head *ListNode, x int) *ListNode {
	ll := make([]*ListNode, 0)
	la := make([]*ListNode, 0)
	if head == nil {
		return nil
	}
	for head != nil {
		if head.Val < x {
			ll = append(ll, head)
		} else {
			la = append(la, head)
		}
		head = head.Next
	}
	ll = append(ll, la...)
	for i := 0; i < len(ll)-1; i++ {
		ll[i].Next = ll[i+1]
	}
	ll[len(ll)-1].Next = nil
	return ll[0]
}

// @lc code=end
