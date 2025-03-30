package main

import "fmt"

func main() {
	fmt.Println(insertionSort([]int{5, 3, 8, 1}))
}

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
