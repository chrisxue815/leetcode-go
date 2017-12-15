package problem

// O(n) time. O(n) space. Iteration.
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := nums[i] - 1

		for nums[index] != index+1 {
			nums[index], index = index+1, nums[index]-1
		}
	}

	result := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}

	return result
}
