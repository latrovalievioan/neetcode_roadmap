// https://leetcode.com/problems/valid-sudoku/
package main

import (
	"fmt"
)

func main() {
	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }

	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	// fmt.Println(isValidSudoku(board))
	fmt.Println(isValidSudoku(board2))
}

func isValidSudoku(board [][]byte) bool {
	return checkRows(board) && checkCols(board) && checkBoxes(board)
}

func checkRows(board [][]byte) bool {
	for _, row := range board {

		lookup := make(map[byte]bool)
		for _, e := range row {
			if e == '.' {
				continue
			}
			_, ok := lookup[e]
			if ok {
				return false
			}
			lookup[e] = true
		}
	}
	return true
}

func checkCols(board [][]byte) bool {
	for i := range board[0] {
		lookup := make(map[byte]bool)
		for j := range board {
			current := board[j][i]
			if current == '.' {
				continue
			}
			_, ok := lookup[current]
			if ok {
				return false
			}
			lookup[current] = true
		}
	}
	return true
}

func checkBoxes(board [][]byte) bool {
	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			lookup := make(map[byte]bool)
			for r := i; r < i+3; r++ {
				for c := j; c < j+3; c++ {
					current := board[r][c]
					if current == '.' {
						continue
					}
					_, ok := lookup[current]
					if ok {
						return false
					}
					lookup[current] = true
				}
			}
		}
	}
	return true
}
