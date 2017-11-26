package problem

import "sort"

// O(nlog(n)) time. O(log(n)) space. Sort.
func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}
