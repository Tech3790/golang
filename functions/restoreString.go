package main

// RestoreString returns a shuffled string using the given array indices
func RestoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for i, v := range indices {
		result[v] = s[i]
	}
	return string(result)
}
