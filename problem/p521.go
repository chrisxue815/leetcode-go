package problem

// O(n) time. O(1) space.
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}
