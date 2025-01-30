package main

import "github.com/charmbracelet/lipgloss"

var (
	// Subgrid border style (darker)
	subgridBorder = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")) // Dark gray

	// Cell border style (lighter)
	cellBorder = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")) // Light gray

	// Cursor style for the selected cell
	cursorStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")). // Purple
			Foreground(lipgloss.Color("230")) // Light yellow

	// Fixed cell style (non-editable)
	fixedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")) // Gray

	// Editable cell style
	editableStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")) // White

	// Error message style
	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")) // Red
)
