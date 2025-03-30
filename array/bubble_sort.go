package main

import "fmt"

func main() {
	// Example usage
	sortedArr := bubbleSort([]int{5, 3, 8, 1})
	fmt.Println(sortedArr) // Output: [1 3 5 8]
}

func bubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
