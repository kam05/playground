package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{5, 3, 8, 1}))
}

func mergeSort(arr []int) []int {
	mid := len(arr) / 2
	if mid == 0 {
		return arr
	}

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right))
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else if left[i] > right[j] {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result

}
