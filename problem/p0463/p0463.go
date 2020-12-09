package problem

// O(n) time. O(1) space. Iteration.
func islandPerimeter(grid [][]int) int {
	islands := 0
	neighbors := 0

	for r, row := range grid {
		for c, num := range row {
			if num == 1 {
				islands++

				if r >= 1 && grid[r-1][c] == 1 {
					neighbors++
				}
				if c >= 1 && row[c-1] == 1 {
					neighbors++
				}
			}
		}
	}

	return islands*4 - neighbors*2
}
