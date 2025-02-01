package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
	grid.WriteString("\n")

	// // Top border
	// grid.WriteString("\n+-------+-------+-------+\n")
	// for y := 0; y < 9; y++ {
	// 	// Add vertical separators between subgrids
	// 	grid.WriteString("| ")
	// 	for x := 0; x < 9; x++ {
	// 		cell := m.grid[y][x]
	// 		value := " "
	// 		if cell.value > 0 {
	// 			value = strconv.Itoa(cell.value)
	// 		}
	//
	// 		// Highlight the cursor
	// 		if x == m.cursorX && y == m.cursorY {
	// 			value = lipgloss.NewStyle().
	// 				Background(lipgloss.Color("62")).  // Purple
	// 				Foreground(lipgloss.Color("230")). // Light yellow
	// 				Render(value)
	// 		}
	//
	// 		grid.WriteString(value + " ")
	//
	// 		// Vertical separator every 3 columns
	// 		if (x+1)%3 == 0 {
	// 			grid.WriteString("| ")
	// 		}
	// 	}
	// 	grid.WriteString("\n")
	//
	// 	// Horizontal separator every 3 rows
	// 	if (y+1)%3 == 0 {
	// 		grid.WriteString("+-------+-------+-------+\n")
	// 	}
	//}
	// var board [][]string = [][]string{
	// 	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	// 	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	// 	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	// 	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	// 	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	// 	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	// 	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	// 	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	// 	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	// }
	//
	var board [][]string = make([][]string, 9)

	for i := 0; i < 9; i++ {
		board[i] = make([]string, 9) // Inner slice
		for j := 0; j < 9; j++ {
			if m.grid[i][j].value == 0 {
				board[i][j] = " " // Represent empty cells as "."
			} else {
				board[i][j] = strconv.Itoa(m.grid[i][j].value) // Convert int to string
			}
		}
	}
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		Rows(board...).StyleFunc(func(row, col int) lipgloss.Style {
		if row == m.cursorY && col == m.cursorX {
			return selected_cell
		}

		return default_cell
	})

	grid.WriteString(t.Render())
	// Add error message
	grid.WriteString("\n")
	grid.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.message))
	grid.WriteString("X:" + strconv.Itoa(m.cursorX) + "\n")
	grid.WriteString("Y:" + strconv.Itoa(m.cursorY))

	return game_window.Render(grid.String())
}
