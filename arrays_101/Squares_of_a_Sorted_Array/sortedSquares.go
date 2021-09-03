package main

import "sort"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}
