package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	groups []ResourceGroup
	cursor int
}

func NewModel(groups []ResourceGroup) Model {
	return Model{groups: groups}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.groups)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := "Resource Groups:\n\n"
	for i, g := range m.groups {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor points here
		}
		s += fmt.Sprintf("%s [%d] %s (%s)\n", cursor, i, g.Name, g.Location)
	}
	s += "\nPress q to quit, j/k to navigate."
	return s
}
