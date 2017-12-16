package problem

// O(n) time. O(1) space. Iteration.
func moveZeroes(nums []int) {
	write := 0
	for _, num := range nums {
		if num != 0 {
			nums[write] = num
			write++
		}
	}
	for ; write < len(nums); write++ {
		nums[write] = 0
	}
}
