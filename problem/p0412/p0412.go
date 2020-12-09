package problem

import "strconv"

// O(n) time. O(n) space. Iteration.
func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		var s string
		if i%15 == 0 {
			s = "FizzBuzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else {
			s = strconv.FormatInt(int64(i), 10)
		}
		result[i-1] = s
	}

	return result
}
