package main

// NumJewelsInStones returns the number of equal characters in two arrays
func NumJewelsInStones(jewels string, stones string) int {
	// brute force solution (high time complexity)
	// mutual := 0
	// for _,v := range J {
	//     for _,val := range S {
	//         if v == val{
	//             mutual++
	//         }
	//     }
	// }

	mutual := 0
	// creating an empty map with length of jewels
	seen := make(map[string]int, len(jewels))
	for _, value := range jewels {
		seen[string(value)] = 1
	}
	// checking if each element in stones already exists in the map
	for _, val := range stones {
		if seen[string(val)] > 0 {
			mutual++
		}
	}
	return mutual
}
