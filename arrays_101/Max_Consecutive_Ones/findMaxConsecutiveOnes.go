// https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3238/

package main

func findMaxConsecutiveOnes(nums []int) int {
	var max_ones int
	var cur_ones int
	for _, v := range nums {
		if v == 1 {
			cur_ones++
			if cur_ones > max_ones {
				max_ones = cur_ones
			}
		} else {
			cur_ones = 0
		}
	}
	return max_ones
}
