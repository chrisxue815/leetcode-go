package problem

// O(1) time. O(1) space. Mathematical induction.
func addDigits(num int) int {
	if num == 0 {
		return 0
	} else {
		return (num-1)%9 + 1
	}
}
