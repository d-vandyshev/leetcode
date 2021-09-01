package main

import "testing"

type testpair struct {
	values     []int
	even_count int
}

var tests = []testpair{
	{[]int{234, 1, 0}, 0},
	{[]int{76, 5, 6785}, 2},
	{[]int{12, 1, 0, 1345, 100, 157890}, 3},
}

func TestFindNumberswithEvenNumberofDigits(t *testing.T) {
	for _, pair := range tests {
		v := findNumberswithEvenNumberofDigits(pair.values)
		if v != pair.even_count {
			t.Error("For", pair.values, "expected", pair.even_count, "got", v)
		}
	}
}
