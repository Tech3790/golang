package main

// KidsWithCandies is a function to see which among many kids
// can have highest number of candies if he/she get the extra candies
func KidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, current := range candies {
		if current > max {
			max = current
		}
	}

	result := make([]bool, len(candies))
	for i, value := range candies {
		if value+extraCandies >= max {
			result[i] = true
		}
	}
	return result
}
