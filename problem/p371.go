package problem

// O(n) time. O(n) space. Bit manipulation.
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b
		b <<= 1
	}
	return a
}
