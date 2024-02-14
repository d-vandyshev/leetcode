package main

import (
	"testing"
)

type testpair struct {
	value  []int
	result bool
}

var tests = []testpair{
	{[]int{10,2,5,3}, true},
	{[]int{3,1,7,11}, false},
}

func TestCheckIfExist(t *testing.T) {
	for _, tp := range tests {
		param := make([]int, len(tp.value))
		copy(param, tp.value)
		v := checkIfExist(param)
		if v != tp.result {
			t.Error("For", tp.value, "expected", tp.result, "got", v)
		}
	}
}
