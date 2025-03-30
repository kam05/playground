package main

func main() {

	// Example usage
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	println(result) // Output: 49

}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		if height[l] < height[r] {
			max = maxInt(max, (r-l)*height[l])
			l++
		} else {
			max = maxInt(max, (r-l)*height[r])
			r--
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
