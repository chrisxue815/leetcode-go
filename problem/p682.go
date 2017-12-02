package problem

import "strconv"

// O(n) time. O(n) space. Stack.
func calPoints(ops []string) int {
	sum := 0
	stack := make([]int, 0, len(ops))

	for _, op := range ops {
		switch op {
		case "+":
			curr := stack[len(stack)-1] + stack[len(stack)-2]
			sum += curr
			stack = append(stack, curr)
		case "D":
			curr := stack[len(stack)-1] * 2
			sum += curr
			stack = append(stack, curr)
		case "C":
			sum -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			i, _ := strconv.ParseInt(op, 10, 32)
			curr := int(i)
			sum += curr
			stack = append(stack, curr)
		}
	}

	return sum
}
