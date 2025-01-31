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

//	anotherStyle = lipgloss.NewStyle().
//			BorderStyle(lipgloss.RoundedBorder()).
//			Border(myCuteBorder, true, true, true, true).
//			Align(lipgloss.Center).Margin(4, 4)

// game_window = lipgloss.NewStyle().
//
//	Padding(1, 2).
//	Border(lipgloss.BlockBorder(), true)
)
