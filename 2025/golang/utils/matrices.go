package utils

func Transpose[T any](arr [][]T) [][]T {
	rowLength := len(arr)
	colLength := len(arr[0])

	transposed := make([][]T, colLength)

	for i := range colLength {
		transposed[i] = make([]T, rowLength)
	}

	for i := range rowLength {
		for j := range colLength {
			transposed[j][i] = arr[i][j]
		}
	}

	return transposed
}

// Top-left to bottom-right diagonals
func TLBRDiagonals[T any](grid [][]T) [][]T {
	rowLength := len(grid)
	colLength := len(grid[0])

	var diagonals [][]T

	for rowStart := range rowLength {
		var diagonal []T
		row, col := rowStart, 0
		for row < rowLength && col < colLength {
			diagonal = append(diagonal, grid[row][col])
			row++
			col++
		}
		diagonals = append(diagonals, diagonal)
	}

	for colStart := 1; colStart < colLength; colStart++ {
		var diagonal []T
		row, col := 0, colStart
		for row < rowLength && col < colLength {
			diagonal = append(diagonal, grid[row][col])
			row++
			col++
		}
		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

// Top-right to bottom-left diagonals
func TRBLDiagonals[T any](grid [][]T) [][]T {
	rowLength := len(grid)
	colLength := len(grid[0])

	var diagonals [][]T

	for rowStart := range rowLength {
		var diagonal []T
		row, col := rowStart, colLength-1
		for row < rowLength && col >= 0 {
			diagonal = append(diagonal, grid[row][col])
			row++
			col--
		}
		diagonals = append(diagonals, diagonal)
	}

	for colStart := colLength - 2; colStart >= 0; colStart-- {
		var diagonal []T
		row, col := 0, colStart
		for row < rowLength && col >= 0 {
			diagonal = append(diagonal, grid[row][col])
			row++
			col--
		}
		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func Diagonals[T any](grid [][]T) [][]T {
	var diagonals [][]T

	diagonals = append(diagonals, TLBRDiagonals(grid)...)
	diagonals = append(diagonals, TRBLDiagonals(grid)...)

	return diagonals
}
