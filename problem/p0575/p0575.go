package problem

type empty struct {
}

// O(n) time. O(n) space. Hash table.
func distributeCandies(candies []int) int {
	kinds := make(map[int]empty)

	for _, candy := range candies {
		kinds[candy] = empty{}
	}

	if len(kinds) < len(candies)/2 {
		return len(kinds)
	} else {
		return len(candies) / 2
	}
}
