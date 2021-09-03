package main

import (
	// "fmt"
	// "fmt"
	"testing"
)

type testpair struct {
  value []int
  result []int
}

var tests = []testpair{
  {[]int{1,0,2,3,0,4,5,0}, []int{1,0,0,2,3,0,0,4}},
  {[]int{1,2,3}, []int{1,2,3}},  
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
    param := make([]int, len(pair.value))
    copy(param, pair.value)
    v := duplicateZeros(param)
    if !Equal(v, pair.result) {
      t.Error("For", pair.value, "expected", pair.result, "got", v)
    }
  }
}
