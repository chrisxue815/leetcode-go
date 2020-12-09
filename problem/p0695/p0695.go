package problem

func getArea(grid [][]int, r int, c int) int {
	if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r]) && grid[r][c] == 1 {
		grid[r][c] = 0
		return 1 + getArea(grid, r-1, c) + getArea(grid, r+1, c) + getArea(grid, r, c-1) + getArea(grid, r, c+1)
	}
	return 0
}

// O(n) time. O(max(w,h)) space. DFS.
func maxAreaOfIsland(grid [][]int) int {
	result := 0

	for r, row := range grid {
		for c, num := range row {
			if num == 1 {
				area := getArea(grid, r, c)
				if result < area {
					result = area
				}
			}
		}
	}

	return result
}
