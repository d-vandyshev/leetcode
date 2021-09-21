package main

func duplicateZeros(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
			i++
		}
	}
	return arr
}
