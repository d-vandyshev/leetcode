package main

import (
	"fmt"
)

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	wasM := false
	uphill := false
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
		if i == 0 && arr[i] > arr[i+1] {
			return false
		}
		if uphill == true && arr[i] > arr[i+1] {
			wasM = true
		}
		if wasM == false && arr[i] > arr[i+1] {
			return false
		}
		if wasM == true && arr[i] < arr[i+1] {
			return false
		}
		if arr[i] < arr[i+1] {
			uphill = true
		} else {
			uphill =false
		}
	}
	if wasM == true {
		return true 
	} else {
		return false
	}
}
