package main

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var got int
	got = findMaxConsecutiveOnes([]int{0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0})
	if got != 4 {
		t.Errorf("findMaxConsecutiveOnes([]int{0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0}) = %d; want 4", got)
	}
	got = findMaxConsecutiveOnes([]int{1, 1, 1, 1, 1, 1})
	if got != 6 {
		t.Errorf("findMaxConsecutiveOnes([]int{1, 1, 1, 1, 1, 1}) = %d; want 6", got)
	}
	got = findMaxConsecutiveOnes([]int{0, 0, 0, 0, 0})
	if got != 0 {
		t.Errorf("findMaxConsecutiveOnes([]int{0, 0, 0, 0, 0}) = %d; want 0", got)
	}
	got = findMaxConsecutiveOnes([]int{1, 1, 0, 0, 0, 0})
	if got != 2 {
		t.Errorf("findMaxConsecutiveOnes([]int{1, 1, 0, 0, 0, 0}) = %d; want 2", got)
	}
	got = findMaxConsecutiveOnes([]int{0, 0, 0, 0, 1, 1, 1})
	if got != 3 {
		t.Errorf("findMaxConsecutiveOnes([]int{0, 0, 0, 0, 1, 1, 1}) = %d; want 3", got)
	}
}
