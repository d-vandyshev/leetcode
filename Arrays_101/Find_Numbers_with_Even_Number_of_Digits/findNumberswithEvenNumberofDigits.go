package main

func findNumberswithEvenNumberofDigits(nums []int) int {	
	numbers_count := 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if isEvenDigit(num) {
			numbers_count++
		}
	}
	return numbers_count;
}

func isEvenDigit(num int) bool {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count % 2 == 0
}
