package problem

// O(n) time. O(1) space. Iteration.
func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	count := 0

	for _, num := range nums {
		if num == 0 {
			if result < count {
				result = count
			}
			count = 0
		} else {
			count++
		}
	}

	if result < count {
		return count
	} else {
		return result
	}
}
