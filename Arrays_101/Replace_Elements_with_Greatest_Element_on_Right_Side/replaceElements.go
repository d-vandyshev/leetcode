package main

func replaceElements(arr []int) []int {
	greatest := arr[len(arr)-1]
	for i := len(arr)-1; i >=0; i-- {
		if i == len(arr)-1 {
			arr[i] = -1
			continue
		}
		current := arr[i]
		arr[i] = greatest
		if greatest < current { 
			greatest = current
		}
	}
	return arr
}
