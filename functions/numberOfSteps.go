package main

// NumberOfSteps returns the number of steps required to reduce an integer to zero
func NumberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		steps++
	}
	return steps
}
