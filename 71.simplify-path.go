package leetcode

import "strings"

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start

type stack struct {
	Size int
	Val  []string
}

func (s *stack) Push(a string) {
	s.Size++
	s.Val = append(s.Val, a)
}

func (s *stack) Pop() string {
	if s.Size == 0 {
		return "nil"
	}
	ans := s.Val[s.Size-1]
	s.Val = s.Val[0 : s.Size-1]
	s.Size--
	return ans
}

func simplifyPath(path string) string {
	sta := new(stack)
	list := strings.Split(path, "/")
	for _, v := range list {
		if v == "" {
			continue
		}
		if v == "." {
			continue
		}
		if v == ".." {
			sta.Pop()
		} else {
			sta.Push(v)
		}
	}
	ans := ""
	for _, v := range sta.Val {
		ans += "/"
		ans += v
	}
	if ans == "" {
		ans = "/"
	}
	return ans
}

// @lc code=end
