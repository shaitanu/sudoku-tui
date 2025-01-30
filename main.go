package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize the Bubble Tea program
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running game: %v", err)
		os.Exit(1)
	}
}
