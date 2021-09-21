package main

import "testing"

type testcase struct {
	nums       []int
	k          int
	resultNums []int
}

var testcases = []testcase{
	{[]int{1, 1, 2}, 2, []int{1, 2}},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	{[]int{1, 1}, 1, []int{1}},
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

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range testcases {
		resK, resNums := removeDuplicates(test.nums)
		if resK != test.k || !Equal(resNums, test.resultNums) {
			t.Error("For nums", test.nums, "expected k", test.k, "nums", test.resultNums, "got", resK, resNums)
		}
	}
}
