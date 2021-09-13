package main

import (
	"testing"
)

type testcase struct {
	arr       []int
	traversal []int
}

func arr2tree(arr []int, index int) *Node {
	if index >= len(arr) {
		return &Node{val: -1000}
	}
	node := Node{val: arr[index]}

	index_left := index * 2 + 1
	if index_left < len(arr) && arr[index_left] != -1000 {
		node.left = arr2tree(arr, index_left)
	}

	index_right := index * 2 + 2
	if index_right < len(arr) && arr[index_right] != -1000 {
		node.right = arr2tree(arr, index_right)
	}
	return &node
}

var tests = []testcase{
	{[]int{10,20,30,40,50,-1000,70,-1000,-1000,100,110,120}, []int{10,20,40,50,100,110,30,70,120}},
	{[]int{1, -1000, 2, 3}, []int{1, 2, 3}},
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
	{[]int{1, -1, 2}, []int{1, 2}},
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

func TestPreorderTraversal(t *testing.T) {
	for i, test := range tests {
		if i > 0 {
			break
		}
		tree := arr2tree(test.arr, 0)
		result_arr := preorderTraversal(*tree)
		if !Equal(result_arr, test.traversal) {
      t.Error("For", test.arr, "expected", test.traversal, "got", result_arr)
		}
	}
}
