package main

func getFreq(pair []int) []int {
	res := []int{}
	for i := 0; i < pair[1]; i++ {
		res = append(res, pair[0])
	}
	return res
}

// DecompressRLElist returns a decompressed list
func DecompressRLElist(nums []int) []int {

	res := []int{}
	for i := 0; i < len(nums); i += 2 {
		temp := getFreq([]int{nums[i+1], nums[i]})
		res = append(res, temp...)
	}
	return res
}
