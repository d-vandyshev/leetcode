// https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3238/

package main

func findMaxConsecutiveOnes(nums []int) int {
	max_ones, cur_ones := 0, 0
	for _, v := range nums {
		if v == 1 {
			cur_ones++
		} else {
			cur_ones = 0
		}
		if cur_ones > max_ones {
			max_ones = cur_ones
		}
	}
	return max_ones
}
