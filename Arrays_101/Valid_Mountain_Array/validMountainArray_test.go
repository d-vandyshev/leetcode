package main

import (
	"testing"
)

type testpair struct {
	value  []int
	result bool
}

var tests = []testpair{
	{[]int{2,1}, false},
	{[]int{3,5,5}, false},
	{[]int{0,3,2,1}, true},
	{[]int{0,1,2,3,4,5,6,7,8,9}, false},
	{[]int{2,0,2}, false},
	{[]int{0,1,2,4,2,1}, true},
}

func TestValidMountainArray(t *testing.T) {
	for _, tp := range tests {
		param := make([]int, len(tp.value))
		copy(param, tp.value)
		v := validMountainArray(param)
		if v != tp.result {
			t.Error("For", tp.value, "expected", tp.result, "got", v)
		}
	}
}
