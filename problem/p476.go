package problem

import "math/bits"

// O(1) time. O(1) space. Bit manipulation.
func findComplement(num int) int {
	return ^num & ((1 << uint(bits.Len(uint(num)))) - 1)
}
