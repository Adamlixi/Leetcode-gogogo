package leetcode

/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyStack struct {
	List []int
	size int
}

func (s *MyStack) Push(x int) {
	s.size++
	s.List = append(s.List, x)
}

func (s *MyStack) Pop() int {
	ans := s.List[len(s.List)-1]
	s.List = s.List[:len(s.List)-1]
	s.size--
	return ans
}

func (s *MyStack) IsEmpty() bool {
	return len(s.List) == 0
}

type MyQueue struct {
	size int
	s1   MyStack
	s2   MyStack
}

func Constructor() MyQueue {
	m := new(MyQueue)
	m.s1 = MyStack{}
	m.s2 = MyStack{}
	return *m
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
	this.size++
	for !this.s2.IsEmpty() {
		this.s1.Push(this.s2.Pop())
	}
	for !this.s1.IsEmpty() {
		this.s2.Push(this.s1.Pop())
	}
}

func (this *MyQueue) Pop() int {
	this.size--
	for !this.s2.IsEmpty() {
		this.s1.Push(this.s2.Pop())
	}
	ans := this.s1.Pop()
	for !this.s1.IsEmpty() {
		this.s2.Push(this.s1.Pop())
	}
	return ans
}

func (this *MyQueue) Peek() int {
	for !this.s2.IsEmpty() {
		this.s1.Push(this.s2.Pop())
	}
	ans := this.s1.Pop()
	this.s1.Push(ans)
	for !this.s1.IsEmpty() {
		this.s2.Push(this.s1.Pop())
	}
	return ans
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
