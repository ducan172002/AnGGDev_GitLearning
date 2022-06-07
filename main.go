package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
	}

	fmt.Println(shortestPathBinaryMatrix(grid))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[row-1][col-1] == 1 {
		return -1
	}
	x_move := []int{1, 0, -1, 0, 1, -1, -1, 1}
	y_move := []int{0, 1, 0, -1, 1, -1, 1, -1}
	move := [][3]int{{0, 0, 1}}
	grid[0][0] = 1
	for len(move) > 0 {
		get := move[0]
		move = move[1:]
		x, y, path := get[0], get[1], get[2]

		if x == row-1 && y == col-1 {
			return path
		}
		for i := 0; i < len(x_move); i++ {
			x_next, y_next := x+x_move[i], y+y_move[i]

			if isValid(grid, x_next, y_next) {
				move = append(move, [3]int{x_next, y_next, path + 1})
				grid[x_next][y_next] = path + 1
			}
		}
	}
	return -1
}

func isValid(grid [][]int, row int, col int) bool {
	if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == 0 {
		return true
	}
	return false
}
