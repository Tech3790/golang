package main

// NumberOfSteps returns the number of steps required to reduce an integer to zero
func NumberOfSteps(num int) int {
	steps := 0
	for i := num; i > 0; {
		if i%2 == 0 {
			i = i / 2
			steps++
		} else {
			i = i - 1
			steps++
		}
	}
	return steps
}
