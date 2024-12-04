package utils

func Transpose[T any](arr [][]T) [][]T {
  rowLength := len(arr)
  colLength := len(arr[0])

  transposed := make([][]T, colLength)

  for i := 0; i < colLength; i++ {
    transposed[i] = make([]T, rowLength)
  }

  for i := 0; i < rowLength; i++ {
    for j := 0; j < colLength; j++ {
      transposed[j][i] = arr[i][j]
    }
  }

  return transposed
}

func Diagonals[T any](grid [][]T) [][]T {
	rowLength := len(grid)
	colLength := len(grid[0])

  diagonals := make([][]T, 0)

	// Top-left to bottom-right diagonals
	for rowStart := 0; rowStart < rowLength; rowStart++ {
		diagonal := make([]T, 0)
		row, col := rowStart, 0
		for row < rowLength && col < colLength {
			diagonal = append(diagonal, grid[row][col])
			row++
			col++
		}
		diagonals = append(diagonals, diagonal)
	}

	for colStart := 1; colStart < colLength; colStart++ {
		diagonal := make([]T, 0)
		row, col := 0, colStart
		for row < rowLength && col < colLength {
			diagonal = append(diagonal, grid[row][col])
			row++
			col++
		}
		diagonals = append(diagonals, diagonal)
	}

	// Top-right to bottom-left diagonals
	for rowStart := 0; rowStart < rowLength; rowStart++ {
		diagonal := make([]T, 0)
		row, col := rowStart, colLength-1
		for row < rowLength && col >= 0 {
			diagonal = append(diagonal, grid[row][col])
			row++
			col--
		}
		diagonals = append(diagonals, diagonal)
	}

	for colStart := colLength - 2; colStart >= 0; colStart-- {
		diagonal := make([]T, 0)
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
