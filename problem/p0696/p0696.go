package problem

// O(n) time. O(1) space. Iteration.
func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	prev := 0
	lo := 0

	for hi := 1; hi < len(s); hi++ {
		if s[hi-1] != s[hi] {
			curr := hi - lo
			if prev < curr {
				result += prev
			} else {
				result += curr
			}
			prev = curr
			lo = hi
		}
	}

	curr := len(s) - lo
	if prev < curr {
		result += prev
	} else {
		result += curr
	}

	return result
}
