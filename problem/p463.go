package problem

// O(n) time. O(1) space. Iteration.
func islandPerimeter(grid [][]int) int {
	peri := 0

	for r, row := range grid {
		for c, num := range row {
			if num == 1 {
				if r == 0 || grid[r-1][c] == 0 {
					peri++
				}
				if r == len(grid)-1 || grid[r+1][c] == 0 {
					peri++
				}
				if c == 0 || row[c-1] == 0 {
					peri++
				}
				if c == len(row)-1 || row[c+1] == 0 {
					peri++
				}
			}
		}
	}

	return peri
}
