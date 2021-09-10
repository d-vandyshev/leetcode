package main

import "testing"

type testpair struct {
  values []int
  result []int
}

var tests = []testpair{
  {[]int{-4,-1,0,3,10}, []int{0,1,9,16,100}},
  {[]int{-7,-3,2,3,11}, []int{4,9,9,49,121}},  
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
  for _, pair := range tests {
    v := sortedSquares(pair.values)
    if !Equal(v, pair.result) {
      t.Error("For", pair.values, "expected", pair.result, "got", v)
    }
  }
}
