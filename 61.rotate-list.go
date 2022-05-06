package leetcode

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func rotateRight(head *ListNode, k int) *ListNode {
	h := head
	if h == nil {
		return h
	}
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	if k >= count {
		k = k % count
	}
	if k == 0 {
		return h
	}
	p := h
	for k != 0 {
		p = p.Next
		k--
	}
	q := h
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	head = q.Next
	q.Next = nil
	p.Next = h
	return head
}

// @lc code=end
