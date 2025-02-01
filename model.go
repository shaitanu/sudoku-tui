package main

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

// Cell represents a single cell in the Sudoku grid
type cell struct {
	value    int
	editable bool
}

// Model represents the game state
type model struct {
	grid       [9][9]cell // 9x9 grid of cells
	cursorX    int        // Current X position of the cursor
	cursorY    int        // Current Y position of the cursor
	message    string     // Message to display (e.g., errors)
	emptycells int        // Number of empty cells
}

// initialModel initializes the game with a sample Sudoku puzzle
func initialModel() model {
	m := model{}
	// sampleGrid := [9][9]int{
	// 	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	// 	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	// 	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	// 	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	// 	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	// 	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	// 	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	// 	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	// 	{3, 4, 5, 2, 8, 6, 0, 7, 9}, // One empty cell
	// }
	var sampleGrid [9][9]int = generateSudoku()
	removeNumbers((&sampleGrid), 5)

	// Populate the grid with the sample puzzle
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sampleGrid[y][x] == 0 {
				m.emptycells++
			}
			m.grid[y][x] = cell{
				value:    sampleGrid[y][x],
				editable: sampleGrid[y][x] == 0,
			}
		}
	}
	return m
}

// isValid checks if a number can be placed at (x, y) without violating Sudoku rules
func isValid(m model, x, y, num int) bool {
	// Check row
	for checkX := 0; checkX < 9; checkX++ {
		if checkX != x && m.grid[y][checkX].value == num {
			return false
		}
	}

	// Check column
	for checkY := 0; checkY < 9; checkY++ {
		if checkY != y && m.grid[checkY][x].value == num {
			return false
		}
	}

	// Check subgrid
	startX := (x / 3) * 3
	startY := (y / 3) * 3
	for cy := 0; cy < 3; cy++ {
		for cx := 0; cx < 3; cx++ {
			if (startY+cy == y) && (startX+cx == x) {
				continue // Skip current cell
			}
			if m.grid[startY+cy][startX+cx].value == num {
				return false
			}
		}
	}
	return true
}

// Init initializes the model (required by Bubble Tea)
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles user input and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursorY > 1 {
				m.cursorY--
			}

		case "down", "j":
			if m.cursorY < 9 {
				m.cursorY++
			}

		case "left", "h":
			if m.cursorX > 0 {
				m.cursorX--
			}

		case "right", "l":
			if m.cursorX < 8 {
				m.cursorX++
			}

			//due to header decrease y by -1
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if m.grid[m.cursorY-1][m.cursorX].editable {
				num, _ := strconv.Atoi(msg.String())
				if isValid(m, m.cursorX, m.cursorY-1, num) {
					if m.grid[m.cursorY-1][m.cursorX].value == 0 {
						m.emptycells-- // Decrease empty cell count
					}
					m.grid[m.cursorY-1][m.cursorX].value = num
					m.message = "" // Clear error message on valid move

					if m.emptycells == 0 {
						m.message = "You Won!"
					}
				} else {
					m.message = "Invalid move!"
				}
			}

		case "0", " ":
			if m.grid[m.cursorY-1][m.cursorX].editable {
				if m.grid[m.cursorY-1][m.cursorX].value != 0 {
					m.emptycells++ // Increase empty cell count
				}
				m.grid[m.cursorY-1][m.cursorX].value = 0
				m.message = ""
			}
		}
	}
	return m, nil
}
