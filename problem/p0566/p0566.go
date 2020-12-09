package problem

// O(n) time. O(n) space. Iteration.
func matrixReshape(nums [][]int, rows int, cols int) [][]int {
	oldRows, oldCols := len(nums), len(nums[0])
	if oldRows*oldCols != rows*cols {
		return nums
	}

	result := make([][]int, rows)
	i := 0

	for r := 0; r < rows; r++ {
		row := make([]int, cols)
		result[r] = row

		for c := 0; c < cols; c++ {
			row[c] = nums[i/oldCols][i%oldCols]
			i++
		}
	}

	return result
}
