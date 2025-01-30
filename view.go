package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the game UI
func (m model) View() string {
	gridData := m.get2DStringGrid() // Helper function to get [][]string
	var grid strings.Builder

	// Top border
	grid.WriteString("+-------+-------+-------+\n")

	for i, row := range gridData {
		// Add vertical separators between subgrids
		grid.WriteString("| ")
		for j, cell := range row {
			// Highlight the cursor
			if i == m.cursorY && j == m.cursorX {
				cell = lipgloss.NewStyle().
					Background(lipgloss.Color("62")).  // Purple
					Foreground(lipgloss.Color("230")). // Light yellow
					Render(cell)
			}

			grid.WriteString(cell + " ")

			// Vertical separator every 3 columns
			if (j+1)%3 == 0 {
				grid.WriteString("| ")
			}
		}
		grid.WriteString("\n")

		// Horizontal separator every 3 rows
		if (i+1)%3 == 0 {
			grid.WriteString("+-------+-------+-------+\n")
		}
	}

	// Add error message
	if m.message != "" {
		grid.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.message))
	}

	return lipgloss.NewStyle().Padding(1, 2).Render(grid.String())
}

// Helper function that returns a 2D string slice
func (m model) get2DStringGrid() [][]string {
	grid := make([][]string, 9)

	for y := 0; y < 9; y++ {
		grid[y] = make([]string, 9)
		for x := 0; x < 9; x++ {
			sgY := y / 3
			sgX := x / 3
			cellY := y % 3
			cellX := x % 3

			cellValue := m.grid[sgY][sgX].cells[cellY][cellX].value

			// Use space for empty cells, else use the number
			if cellValue == 0 {
				grid[y][x] = " "
			} else {
				grid[y][x] = strconv.Itoa(cellValue)
			}
		}
	}
	return grid
}
