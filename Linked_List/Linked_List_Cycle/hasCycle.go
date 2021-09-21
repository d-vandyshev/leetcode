package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor() ListNode {
	return ListNode{Val: -1, Next: nil}
}

func (head *ListNode) AddAtTail(val int) {
	if head.Val == -1 {
		head.Val = val
		return
	}
	cur := head
	for j := 0; ; j++ {
		if cur.Next == nil {
			cur.Next = &ListNode{Val: val, Next: nil}
			return
		}
		cur = cur.Next
	}
}

func hasCycle(head *ListNode) bool {
	oneStep, twoStep := head, head
	for oneStep != nil && twoStep != nil && twoStep.Next != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
		if oneStep == twoStep {
			return true
		}
	}
	return false
}
