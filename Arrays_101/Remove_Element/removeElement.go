package main

func removeElement(nums []int, val int) (int, []int) {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
		if len(nums) > i && nums[i] != val {
			i++
		}
	}
	return len(nums), nums
}
