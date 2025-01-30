package main

import (
	"strconv"

	"github.com/charmbracelet/bubbletea"
)

// Cell represents a single cell in the Sudoku grid
type cell struct {
	value    int
	editable bool
}

// Subgrid represents a 3x3 subgrid in the Sudoku grid
type Subgrid struct {
	cells [3][3]cell
}

// Model represents the game state
type model struct {
	grid    [3][3]Subgrid // 3x3 grid of subgrids
	cursorX int           // Current X position of the cursor
	cursorY int           // Current Y position of the cursor
	message string        // Message to display (e.g., errors)
}

// initialModel initializes the game with a sample Sudoku puzzle
func initialModel() model {
	m := model{}
	sampleGrid := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3}, {4, 0, 0, 8, 0, 3, 0, 0, 1}, {7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0}, {0, 0, 0, 4, 1, 9, 0, 0, 5}, {0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	// Populate the grid with the sample puzzle
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			sgY := y / 3
			sgX := x / 3
			cellY := y % 3
			cellX := x % 3

			m.grid[sgX][sgY].cells[cellX][cellY] = cell{
				value:    sampleGrid[x][y],
				editable: sampleGrid[x][y] == 0,
			}
		}
	}
	return m
}

// getCell retrieves the value of a cell at global coordinates (x, y)
func (m model) getCell(x, y int) int {
	sgY := y / 3
	sgX := x / 3
	cellY := y % 3
	cellX := x % 3
	return m.grid[sgY][sgX].cells[cellY][cellX].value
}

// isValid checks if a number can be placed at (x, y) without violating Sudoku rules
func isValid(m model, x, y, num int) bool {
	// Check row
	for checkX := 0; checkX < 9; checkX++ {
		if checkX != x && m.getCell(checkX, y) == num {
			return false
		}
	}

	// Check column
	for checkY := 0; checkY < 9; checkY++ {
		if checkY != y && m.getCell(x, checkY) == num {
			return false
		}
	}

	// Check subgrid
	sgY := y / 3
	sgX := x / 3
	for cy := 0; cy < 3; cy++ {
		for cx := 0; cx < 3; cx++ {
			if (sgY*3+cy == y) && (sgX*3+cx == x) {
				continue // Skip current cell
			}
			if m.grid[sgY][sgX].cells[cy][cx].value == num {
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
			if m.cursorY > 0 {
				m.cursorY--
			}

		case "down", "j":
			if m.cursorY < 8 {
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

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			sgY := m.cursorY / 3
			sgX := m.cursorX / 3
			cellY := m.cursorY % 3
			cellX := m.cursorX % 3

			if m.grid[sgY][sgX].cells[cellY][cellX].editable {
				num, _ := strconv.Atoi(msg.String())
				if isValid(m, m.cursorX, m.cursorY, num) {
					m.grid[sgY][sgX].cells[cellY][cellX].value = num
					m.message = ""
				} else {
					m.message = "Invalid move!"
				}
			}

		case "0", " ":
			sgY := m.cursorY / 3
			sgX := m.cursorX / 3
			cellY := m.cursorY % 3
			cellX := m.cursorX % 3

			if m.grid[sgY][sgX].cells[cellY][cellX].editable {
				m.grid[sgY][sgX].cells[cellY][cellX].value = 0
				m.message = ""
			}
		}
	}
	return m, nil
}
