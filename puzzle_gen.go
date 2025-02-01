package main

import (
	//	"fmt"
	"math/rand"
)

const gridSize = 9

// Generates a complete valid Sudoku board
func generateSudoku() [gridSize][gridSize]int {
	var board [gridSize][gridSize]int
	fillDiagonalBoxes(&board) // Fill 3x3 diagonal subgrids
	solveSudoku(&board)       // Fill the rest of the board
	return board
}

// Removes numbers to create a puzzle while ensuring a unique solution
func removeNumbers(board *[gridSize][gridSize]int, holes int) {
	positions := rand.Perm(gridSize * gridSize) // Shuffle positions
	for i := 0; i < holes; i++ {
		row, col := positions[i]/gridSize, positions[i]%gridSize
		board[row][col] = 0
	}
}

// Fills 3x3 diagonal boxes with random numbers (ensures valid starting points)
func fillDiagonalBoxes(board *[gridSize][gridSize]int) {
	for i := 0; i < gridSize; i += 3 {
		fillBox(board, i, i)
	}
}

// Fills a 3x3 subgrid with unique random numbers
func fillBox(board *[gridSize][gridSize]int, row, col int) {
	nums := rand.Perm(gridSize) // Shuffle numbers 1-9
	for i := 0; i < 9; i++ {
		board[row+i/3][col+i%3] = nums[i] + 1
	}
}

// Solves the Sudoku using backtracking
func solveSudoku(board *[gridSize][gridSize]int) bool {
	row, col, empty := findEmpty(board)
	if !empty {
		return true // No empty cell left, puzzle solved
	}
	for num := 1; num <= gridSize; num++ {
		if isValidi(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0 // Backtrack
		}
	}
	return false
}

// Finds an empty cell in the board
func findEmpty(board *[gridSize][gridSize]int) (int, int, bool) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// Checks if a number is valid in a given position
func isValidi(board *[gridSize][gridSize]int, row, col, num int) bool {
	for i := 0; i < gridSize; i++ {
		if board[row][i] == num || board[i][col] == num || board[row-row%3+i/3][col-col%3+i%3] == num {
			return false
		}
	}
	return true
}

// Prints the Sudoku board
// func printBoard(board [gridSize][gridSize]int) {
// 	for i, row := range board {
// 		if i%3 == 0 && i != 0 {
// 			fmt.Println("------+-------+------")
// 		}
// 		for j, val := range row {
// 			if j%3 == 0 && j != 0 {
// 				fmt.Print("| ")
// 			}
// 			if val == 0 {
// 				fmt.Print(". ")
// 			} else {
// 				fmt.Print(val, " ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func sample() {
//
// 	// Generate a full Sudoku board
// 	solvedBoard := generateSudoku()
//
// 	// Create a puzzle by removing numbers
// 	removeNumbers(&solvedBoard, 45) // Adjust difficulty (20-55 holes)
//
// 	// Print the puzzle
// 	// fmt.Println("Generated Sudoku Puzzle:")
// 	// printBoard(solvedBoard)
// }
