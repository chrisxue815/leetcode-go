package problem

import "math"

// O(sqrt(n)) time. O(1) space. Math.
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))

	for ; w > 1; w-- {
		if area%w == 0 {
			break
		}
	}

	return []int{area / w, w}
}
