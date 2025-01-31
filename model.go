package main

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type cell struct {
	value    int
	editable bool
}

type Subgrid struct {
	cells [3][3]cell
}

type model struct {
	grid       [3][3]Subgrid // 3x3 grid of 3x3 subgrids
	cursorX    int           // Current X position (0-8)
	cursorY    int           // Current Y position (0-8)
	message    string
	emptycells int
}

func initialModel() model {
	m := model{}
	sampleGrid := [9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 0, 7, 9},
	}

	// Convert 9x9 grid to 3x3 subgrids
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			sgY := y / 3
			sgX := x / 3
			cellY := y % 3
			cellX := x % 3

			if sampleGrid[y][x] == 0 {
				m.emptycells++
			}
			m.grid[sgY][sgX].cells[cellY][cellX] = cell{
				value:    sampleGrid[y][x],
				editable: sampleGrid[y][x] == 0,
			}
		}
	}
	return m
}

func isValid(m model, x, y, num int) bool {
	// Check row
	for checkX := 0; checkX < 9; checkX++ {
		sgY := y / 3
		sgX := checkX / 3
		cellY := y % 3
		cellX := checkX % 3
		if checkX != x && m.grid[sgY][sgX].cells[cellY][cellX].value == num {
			return false
		}
	}

	// Check column
	for checkY := 0; checkY < 9; checkY++ {
		sgY := checkY / 3
		sgX := x / 3
		cellY := checkY % 3
		cellX := x % 3
		if checkY != y && m.grid[sgY][sgX].cells[cellY][cellX].value == num {
			return false
		}
	}

	// Check subgrid
	sgY := y / 3
	sgX := x / 3
	for cy := 0; cy < 3; cy++ {
		for cx := 0; cx < 3; cx++ {
			if (sgY*3+cy == y) && (sgX*3+cx == x) {
				continue
			}
			if m.grid[sgY][sgX].cells[cy][cx].value == num {
				return false
			}
		}
	}
	return true
}

func (m model) Init() tea.Cmd {
	return nil
}

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
					if m.grid[sgY][sgX].cells[cellY][cellX].value == 0 {
						m.emptycells--
					}
					m.grid[sgY][sgX].cells[cellY][cellX].value = num
					m.message = ""

					if m.emptycells == 0 {
						m.message = "You Won!"
					}
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
				if m.grid[sgY][sgX].cells[cellY][cellX].value != 0 {
					m.emptycells++
				}
				m.grid[sgY][sgX].cells[cellY][cellX].value = 0
				m.message = ""
			}
		}
	}
	return m, nil
}
