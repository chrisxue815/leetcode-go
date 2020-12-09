package problem

import "math/bits"

// O(1) time. O(1) space. Bit manipulation.
func findComplement(num int) int {
	mask := (1 << uint(bits.Len(uint(num)))) - 1
	return num ^ mask
}
