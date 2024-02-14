package main

import (
	// "fmt"
)

func checkIfExist(arr []int) bool {
	// fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		// fmt.Println("I elem:", arr[i])
		for j := 0; j < i; j++ {
			// fmt.Println("\tJ elem:", arr[j])
			if arr[i] == 2 * arr[j] || 2 * arr[i] == arr[j]{
				return true
			}
		}
	}
	return false
}
