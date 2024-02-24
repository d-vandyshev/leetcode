package main

import (
	"testing"
	"fmt"
)

type testcases struct {
	input []int
	output []int
}

var tests = []testcases{
	{[]int{17,18,5,4,6,1}, []int{18,6,6,6,1,-1}},
	{[]int{400}, []int{-1}},
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

func TestReplaceElements(t *testing.T) {
	fmt.Println("tests:", tests)

	for _, tc := range tests {
		fmt.Println("tc.input:", tc.input)
		v := replaceElements(tc.input)
		if !Equal(v, tc.output) {
			t.Error("Expected", tc.output, "got", v)
		}
	}
}
