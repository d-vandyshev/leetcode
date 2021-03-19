package main

import "testing"

type testpair struct {
	values        []int
	max_cons_ones int
}

var tests = []testpair{
	{[]int{1, 1, 1, 1, 1, 1}, 6},
	{[]int{0, 0, 0, 0, 0}, 0},
	{[]int{1, 1, 0, 0, 0, 0}, 2},
	{[]int{0, 0, 0, 0, 1, 1, 1}, 3},
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	for _, pair := range tests {
		v := findMaxConsecutiveOnes(pair.values)
		if v != pair.max_cons_ones {
			t.Error("For", pair.values, "expected", pair.max_cons_ones, "got", v)
		}
	}
}
