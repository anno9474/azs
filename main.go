package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	rgs, err := FetchResourceGroups()
	if err != nil {
		fmt.Println("Error fetching resource groups:", err)
		os.Exit(1)
	}

	m := NewModel(rgs)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
