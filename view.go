package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var sudoku = `███████ ██    ██ ██████   ██████  ██   ██ ██    ██
██      ██    ██ ██   ██ ██    ██ ██  ██  ██    ██
███████ ██    ██ ██   ██ ██    ██ █████   ██    ██
     ██ ██    ██ ██   ██ ██    ██ ██  ██  ██    ██
███████  ██████  ██████   ██████  ██   ██  ██████`

// View renders the game UI
func (m model) View() string {
	var grid strings.Builder

	//logo
	grid.WriteString(anotherStyle.Render(sudoku))

	// Top border
	grid.WriteString("\n+-------+-------+-------+\n")

	for y := 0; y < 9; y++ {
		// Add vertical separators between subgrids
		grid.WriteString("| ")
		for x := 0; x < 9; x++ {
			cell := m.grid[y][x]
			value := " "
			if cell.value > 0 {
				value = strconv.Itoa(cell.value)
			}

			// Highlight the cursor
			if x == m.cursorX && y == m.cursorY {
				value = lipgloss.NewStyle().
					Background(lipgloss.Color("62")).  // Purple
					Foreground(lipgloss.Color("230")). // Light yellow
					Render(value)
			}

			grid.WriteString(value + " ")

			// Vertical separator every 3 columns
			if (x+1)%3 == 0 {
				grid.WriteString("| ")
			}
		}
		grid.WriteString("\n")

		// Horizontal separator every 3 rows
		if (y+1)%3 == 0 {
			grid.WriteString("+-------+-------+-------+\n")
		}
	}

	// Add error message
	if m.message != "" {
		grid.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.message))
	}
	return game_window.Render(grid.String())
}
