package leetcode

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	numMap := make(map[int]int, 0)
	for p != nil {
		if _, ok := numMap[p.Val]; ok {
			numMap[p.Val] += 1
		} else {
			numMap[p.Val] = 1
		}
		p = p.Next
	}
	p = new(ListNode)
	s := p
	for head != nil {
		if numMap[head.Val] > 1 {
			head = head.Next
			continue
		}
		s.Next = head
		head = head.Next
		s = s.Next
	}
	if s != nil {
		s.Next = nil
	}
	return p.Next
}

// @lc code=end
