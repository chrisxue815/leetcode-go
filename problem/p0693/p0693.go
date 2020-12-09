package problem

// O(1) time. O(1) space. Bit manipulation.
func hasAlternatingBits(n int) bool {
	n ^= n >> 1
	return n&(n+1) == 0
}
