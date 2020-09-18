package main

// Reverse is a function to reverse a string
func Reverse(word string) string {
	foo := []rune(word)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		foo[i], foo[j] = foo[j], foo[i]
	}
	return string(foo)
}

// RunningSum is a function to sum all elements in array
func RunningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}

// Shuffle is a function to shuffle an array
func Shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := range nums[:n] {
		res[i*2], res[i*2+1] = nums[i], nums[i+n]
	}
	return res
}
