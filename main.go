package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
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
	var s string

	if m.Submitted {
		s = installationView(m)
	} else {
		s = choicesView(m)
	}

	return indent.String("\n"+s+"\n\n", 2)
}

func choicesView(m model) string {
	s := keyword("What do you want to install?\n\n")

	for i, choice := range choices {
		name := choice.Name
		cursorChar := " "
		checkedChar := " "
		isCursorOnChoice := m.Cursor == i
		_, choiceSelected := m.Selected[i]

		if isCursorOnChoice {
			cursorChar = "×"
		}

		if choiceSelected {
			checkedChar = "×"
			name = colorFg(name, "#0096FF")
		}

		bracketed := "[ ] "

		if cursorChar == "×" && choiceSelected {
			bracketed = colorFg("[×]", "#FF8000") + " "
		} else if cursorChar == "×" || checkedChar == "×" {
			bracketed = colorFg("[×]", "#0096FF") + " "
		}

		s += fmt.Sprintf("%s%s\n\n", bracketed, name)
	}

	s += subtle("s: submit") + dot + subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")

	return s
}

func installationView(m model) string {
	msg := "Setting up your system..."
	label := "Downloading:"

	return msg + "\n\n" + label
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
