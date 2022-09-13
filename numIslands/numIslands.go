package numislands

func travelGrid(x, y int, grid [][]byte) {

	row := len(grid)
	col := len(grid[0])
	if x < 0 || y < 0 || x == row || y == col {
		return
	}
	if grid[x][y] == '0' || grid[x][y] == '2' {
		return
	}
	grid[x][y] = '2'
	travelGrid(x-1, y, grid)
	travelGrid(x+1, y, grid)
	travelGrid(x, y-1, grid)
	travelGrid(x, y+1, grid)
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	numIslands := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				numIslands++
				travelGrid(i, j, grid)
			}
		}
	}
	return numIslands
}

func Test() {
	numIslands(nil)
}
