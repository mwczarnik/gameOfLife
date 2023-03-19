package main

import (
	"fmt"
	"os"

	"gameOfLife/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		// Enable cursor blinking if crashed
		fmt.Print("\033[?25h")
		os.Exit(1)
	}

	// Enable cursor blinking after program execution
	fmt.Print("\033[?25h")
}
