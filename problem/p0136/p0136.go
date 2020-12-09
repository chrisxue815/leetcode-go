package test_0136

// O(n) time. O(1) space. Bit manipulation.
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
