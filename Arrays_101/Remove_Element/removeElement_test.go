package main

import (
	"testing"
)

type testcase struct {
  nums []int
  val int
	k int
	res_nums []int
}

var tests = []testcase{
  {[]int{3,2,2,3}, 3, 2, []int{2,2}},
  {[]int{0,1,2,2,3,0,4,2}, 2, 5, []int{0,1,4,0,3}},
  {[]int{1}, 1, 0, []int{}},
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
    k, res_nums := removeElement(tc.nums, tc.val)
    if k != tc.k || !Equal(res_nums, tc.res_nums) {
      t.Error("For k", tc.nums, tc.val, "expected", tc.k, tc.res_nums, "got", k, res_nums)
    }
  }
}
