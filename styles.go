package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Add a purple, rectangular border
	//style = lipgloss.NewStyle().
	//	BorderStyle(lipgloss.NormalBorder()).
	//	BorderForeground(lipgloss.Color("63"))
	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
	myCuteBorder = lipgloss.Border{
		Top:         "._.:*:",
		Bottom:      "._.:*:",
		Left:        "|*",
		Right:       "|*",
		TopLeft:     "*",
		TopRight:    "*",
		BottomLeft:  "*",
		BottomRight: "*",
	}

	//logo border for now
	anotherStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			Border(myCuteBorder, true, true, true, true).
			Align(lipgloss.Center).Margin(8, 4)
	//gane border
	game_window = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.BlockBorder(), true)

	//selected cell
	selected_cell = lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230"))

	default_cell = lipgloss.NewStyle().
			Padding(0, 1)
)
