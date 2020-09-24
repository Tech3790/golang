package main

import (
	"strconv"
)

// SubtractProductAndSum returns the difference between the product of a given integer's digits and the sum of them
func SubtractProductAndSum(n int) int {
	product, sum := 1, 0
	stringifiedNumber := strconv.Itoa(n)

	for _, value := range stringifiedNumber {
		number, _ := strconv.Atoi(string(value))
		product *= int(number)
		sum += int(number)
	}
	return product - sum
}
