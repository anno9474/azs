package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestModelNavigation(t *testing.T) {
	// Sample resource groups
	groups := []ResourceGroup{
		{"rg1", "eastus"},
		{"rg2", "westus"},
	}

	// Create a fresh model
	m := NewModel(groups)

	// Initial cursor should be 0
	if m.cursor != 0 {
		t.Errorf("expected cursor 0 initially, got %d", m.cursor)
	}

	// Simulate down key
	model, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = model.(Model)
	if m.cursor != 1 {
		t.Errorf("expected cursor 1 after down, got %d", m.cursor)
	}

	// Simulate down key again (should not go past last element)
	model, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = model.(Model)
	if m.cursor != 1 {
		t.Errorf("expected cursor 1 after down at bottom, got %d", m.cursor)
	}

	// Simulate up key
	model, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m = model.(Model)
	if m.cursor != 0 {
		t.Errorf("expected cursor 0 after up, got %d", m.cursor)
	}

	// Simulate up key again (should not go below 0)
	model, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m = model.(Model)
	if m.cursor != 0 {
		t.Errorf("expected cursor 0 after up at top, got %d", m.cursor)
	}
}
