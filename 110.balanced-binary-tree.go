package leetcode

import "math"

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	baL := getDeepth(root.Left)
	baR := getDeepth(root.Right)
	if math.Abs(float64(baL-baR)) > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

func getDeepth(root *TreeNode) int {
	if root == nil {
		return 1
	}
	countL := getDeepth(root.Left) + 1
	countR := getDeepth(root.Right) + 1
	if countL > countR {
		return countL
	} else {
		return countR
	}
}

// @lc code=end
