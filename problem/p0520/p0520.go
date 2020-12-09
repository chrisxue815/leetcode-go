package problem

// O(n) time. O(1) space. Iteration.
func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	isCapital := word[1] <= 'Z'
	for i := 2; i < len(word); i++ {
		if word[i] <= 'Z' != isCapital {
			return false
		}
	}

	return !isCapital || word[0] <= 'Z'
}
