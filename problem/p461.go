package problem

import "math/bits"

// O(1) time. O(1) space. Xor and bit count.
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
