package main

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	half_sum := 0
	for i := 0; i < len(nums); i++ {
		if sum-half_sum-nums[i] == half_sum {
			return i
		}
		half_sum += nums[i]
	}
	return -1
}
