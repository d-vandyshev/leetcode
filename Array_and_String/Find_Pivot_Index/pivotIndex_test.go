package main

import "testing"

type testpair struct {
	nums        []int
	pivot_index int
}

var tests = []testpair{
	{[]int{1,7,3,6,5,6}, 3},
	{[]int{1,2,3}, -1},
	{[]int{2,1,-1}, 0},
	{[]int{-1,-1,0,0,-1,-1}, 2},
}

func TestFindNumberswithEvenNumberofDigits(t *testing.T) {
	for _, pair := range tests {
		v := pivotIndex(pair.nums)
		if v != pair.pivot_index {
			t.Error("For", pair.nums, "expected", pair.pivot_index, "got", v)
		}
	}
}
