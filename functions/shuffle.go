package main

// Shuffle is a function to shuffle an array
func Shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := range nums[:n] {
		res[i*2], res[i*2+1] = nums[i], nums[i+n]
	}
	return res
}
