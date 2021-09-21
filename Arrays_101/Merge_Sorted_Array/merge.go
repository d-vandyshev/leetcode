package main

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for j, n2 := range nums2 {
		nums2_item_used := false
		for i := 0; i < m+j; i++ {
			if nums1[i] > n2 {
				for z := m + n - 1; z > i; z-- {
					nums1[z] = nums1[z-1]
				}
				nums1[i] = n2
				nums2_item_used = true
				break
			}
		}
		if !nums2_item_used {
			nums1[m+j] = n2
		}
	}
	return nums1
}
