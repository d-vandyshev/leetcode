package main

import (
	"testing"
)

type testcase struct {
	nums1  []int
	m      int
	nums2  []int
	n      int
	result []int
}

var tests = []testcase{
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	{[]int{1}, 1, []int{}, 0, []int{1}},
	{[]int{0}, 0, []int{1}, 1, []int{1}},
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

func TestSortedSquares(t *testing.T) {
	for _, tc := range tests {
		v := merge(tc.nums1, tc.m, tc.nums2, tc.n)
		if !Equal(v, tc.result) {
			t.Error("For", tc.nums1, tc.m, tc.nums2, tc.n, "expected", tc.result, "got", v)
		}
	}
}
