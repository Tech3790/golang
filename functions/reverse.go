package main

// Reverse is a function to reverse a string
func Reverse(word string) string {
	foo := []rune(word)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		foo[i], foo[j] = foo[j], foo[i]
	}
	return string(foo)
}
