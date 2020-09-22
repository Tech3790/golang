package main

func numIdenticalPairs(nums []int) int {
	var counter int
	dictionary := make(map[int]int)
	for _, v := range nums {
		n, ok := dictionary[v]
		if ok {
			counter += n
		}
		dictionary[v]++
	}
	return counter
}
