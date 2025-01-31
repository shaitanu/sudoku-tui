package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	anotherStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	game_window  = lipgloss.NewStyle().Padding(1, 2)
	sudoku       = `███████ ██    ██ ██████   ██████  ██   ██ ██    ██
██      ██    ██ ██   ██ ██    ██ ██  ██  ██    ██
███████ ██    ██ ██   ██ ██    ██ █████   ██    ██
     ██ ██    ██ ██   ██ ██    ██ ██  ██  ██    ██
███████  ██████  ██████   ██████  ██   ██  ██████`
)

func (m model) View() string {
	var grid strings.Builder

	// Logo
	grid.WriteString(anotherStyle.Render(sudoku))
	grid.WriteString("\n+-------+-------+-------+\n")

	for y := 0; y < 9; y++ {
		sgY := y / 3
		cellY := y % 3
		grid.WriteString("| ")

		for x := 0; x < 9; x++ {
			sgX := x / 3
			cellX := x % 3
			c := m.grid[sgY][sgX].cells[cellY][cellX]

			value := " "
			if c.value > 0 {
				value = strconv.Itoa(c.value)
			}

			if x == m.cursorX && y == m.cursorY {
				value = lipgloss.NewStyle().
					Background(lipgloss.Color("62")).
					Foreground(lipgloss.Color("230")).
					Render(value)
			}

			grid.WriteString(value + " ")

			if (x+1)%3 == 0 {
				grid.WriteString("| ")
			}
		}
		grid.WriteString("\n")

		if (y+1)%3 == 0 {
			grid.WriteString("+-------+-------+-------+\n")
		}
	}

	if m.message != "" {
		grid.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.message))
	}
	return game_window.Render(grid.String())
}
