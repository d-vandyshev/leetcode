package main

import (
	"testing"
)

type testcase struct {
	Nodes []int
	Pos int
	Result bool
}

var testcases = []testcase{
	{[]int{3,2,0,-4}, 1, true},
	{[]int{1,2}, 0, true},
	{[]int{1}, -1, false},
}

func createList(values []int) ListNode {
	head := Constructor()
	for _, v := range values {
		head.AddAtTail(v)
	}
	return head
}

func setCicle(list *ListNode, pos int) {
	cur := list
	var pointClosure *ListNode
	for j := 0; ; j++ {
		if pos == j {
			pointClosure = cur
		}
		if cur.Next == nil {
			cur.Next = pointClosure
			return
		}
		cur = cur.Next
	}
}

func TestHasCycle(t *testing.T) {
	for _, test := range testcases {
		list := createList(test.Nodes)
		if test.Pos != -1 {
			setCicle(&list, test.Pos)
		}
		r := hasCycle(&list)
		if r != test.Result {
			t.Error("For list", test.Nodes, "and pos", test.Pos, "expected", test.Result, "got", r)
		}
	}
}
