package leetcode

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head
	q := head.Next
	for p != nil && q != nil {
		if p.Val == q.Val {
			q = q.Next
			p.Next = q
		} else {
			p.Next = q
			p = p.Next
			q = q.Next
		}
	}
	return head
}

// @lc code=end
