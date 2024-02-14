package main

import "testing"

type testcase struct {
	arr       []int
	traversal []int
}

func arr2tree(arr []int, index int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	if index >= len(arr) {
		return &TreeNode{Val: -1000}
	}
	node := TreeNode{Val: arr[index]}

	index_left := index*2 + 1
	if index_left < len(arr) && arr[index_left] != -1000 {
		node.Left = arr2tree(arr, index_left)
	}

	index_right := index*2 + 2
	if index_right < len(arr) && arr[index_right] != -1000 {
		node.Right = arr2tree(arr, index_right)
	}
	return &node
}

var tests = []testcase{
	{[]int{1,-1000,2,-1000,-1000,3,-1000}, []int{1,3,2}},
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{2, 1}},
	{[]int{1, -1000, 2}, []int{1, 2}},
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestInorderTraversal(t *testing.T) {
	for _, test := range tests {
		tree := arr2tree(test.arr, 0)
		result_arr := inorderTraversal(tree)
		if !Equal(result_arr, test.traversal) {
			t.Error("For", test.arr, "expected", test.traversal, "got", result_arr)
		}
	}
}


