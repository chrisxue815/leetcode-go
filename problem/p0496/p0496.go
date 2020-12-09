package problem

// O(n) time. O(n) space. Stack, hash table.
func nextGreaterElement(findNums []int, nums []int) []int {
	nextGreater := make(map[int]int, len(nums))
	stack := make([]int, 0, len(nums))

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			nextGreater[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for i, num := range findNums {
		next, ok := nextGreater[num]
		if !ok {
			next = -1
		}
		findNums[i] = next
	}

	return findNums
}
