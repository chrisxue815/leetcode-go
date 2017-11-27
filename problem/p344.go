package problem

// O(n) time. O(n) space. Iteration.
func reverseString(s string) string {
	buf := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		buf[i] = s[len(s)-1-i]
	}
	return string(buf)
}
