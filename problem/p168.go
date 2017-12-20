package problem

// O(n) time. O(1) space. Math.
func convertToTitle(n int) string {
	const initialCapacity = 16
	s := make([]byte, 0, initialCapacity)

	for ; n > 0; n /= 26 {
		n--
		s = append(s, 'A'+byte(n%26))
	}

	lo, hi := 0, len(s)-1
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}

	return string(s)
}
