package main

import (
	"log"
	"strings"
)

func day4() {
	data := getRawInputData(2021, 4)
	parsedData := strings.Split(data, "\n\n")
	randomNums := StringSliceToIntegerSlice(strings.Split(parsedData[0], ","))
	parsedData = parsedData[1:]

	boards := getBingoBoards(parsedData)

	day4part1(randomNums, boards)
	day4part2(randomNums, boards)

}

func day4part1(randomNums []int, boards [][][]int) {

	markedIndices := getEmptyMarkedIndices(len(boards))

	for _, number := range randomNums {

		markNumbers(number, &markedIndices, boards)
		boardNum := boardWhichWon(markedIndices)

		if boardNum != -1 {
			sum := getSumOfUnmarkedNumbers(markedIndices[boardNum], boards[boardNum])
			log.Println("Answer for Day 4, Part 1 is", sum*number)
			return
		}
	}
}

func day4part2(randomNums []int, boards [][][]int) {

	markedIndices := getEmptyMarkedIndices(len(boards))

	for _, number := range randomNums {
		markNumbers(number, &markedIndices, boards)
	}

	for _, number := range reverseRandomNumbers(randomNums) {
		currWinningBoards := getWinningBoards(markedIndices, boards)
		unmarkNumbers(number, &markedIndices, boards)
		boardsWinning := getWinningBoards(markedIndices, boards)
		if len(boardsWinning) != len(currWinningBoards) {
			boardNum := whichBoardDiffers(boardsWinning, currWinningBoards)
			sum := getSumOfUnmarkedNumbers(markedIndices[boardNum], boards[boardNum]) - number
			log.Println("Answer for Day 4, Part 2 is", number*sum)
			return
		}

	}

}

func getBingoBoards(parsedData []string) [][][]int {

	var boards [][][]int

	for _, block := range parsedData {
		lines := strings.Split(block, "\n")
		var board [][]int
		for _, line := range lines {
			splitLine := removeEmptyElements(strings.Split(line, " "))
			x := StringSliceToIntegerSlice(splitLine)
			board = append(board, x)
		}
		boards = append(boards, board)
	}
	// HACK: WEIRD EXTRA ROW IN THE LAST BOARD
	boards[len(boards)-1] = boards[len(boards)-1][:len(boards[len(boards)-1])-1]

	return boards

}

func reverseRandomNumbers(randNums []int) []int {
	if len(randNums) == 0 {
		return randNums
	}
	return append(reverseRandomNumbers(randNums[1:]), randNums[0])
}

func whichBoardDiffers(boardsWinning, currWinningBoards []int) int {
	for _, b1index := range currWinningBoards {
		if !isElementinSlice(b1index, boardsWinning) {
			return b1index
		}
	}
	return -1

}

func isElementinSlice(element int, s []int) bool {
	for _, val := range s {
		if val == element {
			return true
		}
	}
	return false
}

func getWinningBoards(markedIndices [][][]bool, boards [][][]int) []int {
	var boardsWinning []int

	for boardNum, board := range boards {
		if isBoardWinning(markedIndices[boardNum], board) {
			boardsWinning = append(boardsWinning, boardNum)
		}
	}

	return boardsWinning

}

func isBoardWinning(markedIndices [][]bool, board [][]int) bool {

	isColumnMarked := checkColumnMarked(markedIndices)
	isRowMarked := checkRowMarked(markedIndices)

	if isRowMarked || isColumnMarked {
		return true
	}

	return false
}

func getSumOfUnmarkedNumbers(markedIndices [][]bool, board [][]int) int {

	var sum int

	for rowNum, row := range markedIndices {
		for colNum, isMarked := range row {
			if !isMarked {
				sum += board[rowNum][colNum]
			}
		}
	}

	return sum

}

func boardWhichWon(markedIndices [][][]bool) int {

	for boardNum, board := range markedIndices {
		isColumnMarked := checkColumnMarked(board)
		isRowMarked := checkRowMarked(board)

		if isRowMarked || isColumnMarked {
			return boardNum
		}

	}

	return -1
}

func checkAllTrue(boolSlice []bool) bool {
	for _, boolVal := range boolSlice {
		if !boolVal {
			return false
		}
	}
	return true
}

func checkRowMarked(board [][]bool) bool {
	for _, row := range board {
		isAllTrue := checkAllTrue(row)
		if isAllTrue {
			return true
		}
	}
	return false
}

func checkColumnMarked(board [][]bool) bool {
	for i := 0; i < len(board); i++ {
		var column []bool
		for _, row := range board {
			column = append(column, row[i])
		}
		isAllTrue := checkAllTrue(column)
		if isAllTrue {
			return true
		}
	}
	return false
}

func unmarkNumbers(number int, markedIndices *[][][]bool, boards [][][]int) {

	for boardNum, board := range boards {
		x, y := findNumberInBoard(number, board)
		if x == -1 && y == -1 {
			continue
		}

		(*markedIndices)[boardNum][x][y] = false

	}

}

func markNumbers(number int, markedIndices *[][][]bool, boards [][][]int) {

	for boardNum, board := range boards {
		x, y := findNumberInBoard(number, board)
		if x == -1 && y == -1 {
			continue
		}

		(*markedIndices)[boardNum][x][y] = true

	}

}

func findNumberInBoard(randNumber int, board [][]int) (int, int) {
	for i, line := range board {
		for j, number := range line {

			if number == randNumber {
				return i, j
			}
		}

	}
	return -1, -1
}

func removeEmptyElements(s []string) []string {
	var tmp []string
	for _, val := range s {
		if val != "" {
			tmp = append(tmp, val)
		}
	}
	return tmp
}

func getEmptyMarkedIndices(boardCount int) [][][]bool {
	var markedIndices [][][]bool

	for i := 0; i < boardCount; i++ {

		var emptyRows [][]bool

		for j := 0; j < 5; j++ {

			var emptyColumns []bool

			for k := 0; k < 5; k++ {
				emptyColumns = append(emptyColumns, false)
			}
			emptyRows = append(emptyRows, emptyColumns)
		}
		markedIndices = append(markedIndices, emptyRows)
	}

	return markedIndices

}
