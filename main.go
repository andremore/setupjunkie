package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return handleQuit(m)
		case "up", "k":
			return moveCursorUp(m), nil
		case "down", "j":
			return moveCursorDown(m), nil
		case "enter", " ":
			return toggleChoice(m), nil
		case "s":
			if !m.Submitted {
				return handleSubmit(m)
			}
		}
	case string:
		fmt.Println(msg)
		return m, tea.Quit
	case error:
		fmt.Println("Error:", msg)
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	if m.Submitted {
		return "Setting up your system..." 
	}

	s := "What do you want to set up?\n\n"
	for i, choice := range choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}
	s += "\nPress q to quit. Press s to submit.\n"

	return s
}

func main() {
	if err := ensureGitAndCurlInstalled(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
