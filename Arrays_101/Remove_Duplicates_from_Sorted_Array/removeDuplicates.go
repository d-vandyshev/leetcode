package main

func removeDuplicates(nums []int) (int, []int) {
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i = i - 1
		}
	}
	return len(nums), nums
}
